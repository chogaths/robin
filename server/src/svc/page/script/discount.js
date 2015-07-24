function template(obj) {
	
	var devide = $("#data")
	devide.empty()
	
	if (obj != null) {
		devide.append("<tr> <th>编号</th> <th>商品类型</th> <th>商品的值</th> <th>销售数额</th> <th>品质等级</th> <th>刷新几率</th> <th>使用货币</th> <th>货币数量</th> <th>开关</th> <th>商品标题</th> </tr>")
		
		for (var index in obj) {
			var data = obj[index]
			devide.append(
				"<tr>" +
				"<td> <input name=id style=width:80px;text-align:center disabled value=" + data.Id + " /> </td>" +
				"<td> <input name=pstyle style=width:80px;text-align:center disabled value=" + data.Pstyle + " /> </td>" +
				"<td> <input name=pvalue style=width:80px;text-align:center disabled value=" + data.Pvalue + " /> </td>" +
				"<td> <input name=pcount style=width:80px;text-align:center disabled value=" + data.Pcount + " /> </td>" +
				"<td> <input name=prlevel style=width:80px;text-align:center disabled value=" + data.Prlevel + " /> </td>" +
				"<td> <input name=pjilv style=width:80px;text-align:center disabled value=" + data.Pjilv + " /> </td>" +
				"<td> <input name=pcurrency style=width:80px;text-align:center disabled value=" + data.Pcurrency + " /> </td>" +
				"<td> <input name=pprice style=width:80px;text-align:center disabled copy=" + data.Pprice + " value=" + data.Pprice + " /> </td>" +
				"<td> <input name=popen style=width:80px;text-align:center disabled value=" + data.Popen + " /> </td>" +
				"<td> <input name=ptitle style=text-align:center disabled value=" + data.Ptitle + " /> </td>" +
				"<td> <button name=change>修改</button> </td>" +
				"<td> <button name=commit hidden>确认</button> </td>" +
				"<td> <button name=cancel hidden>取消</button> </td>" +
				"</tr>"
			)
		}
	} else {
		devide.append("无数据")
	}
	
	register()
	
}

function register() {
	
	$("[name=change]").click(function(){
		$(this).hide()
		var root = $(this).parent().parent()
		root.find("[name=commit]").show()
		root.find("[name=cancel]").show()
		root.find("[name=pprice]").removeAttr("disabled")
		root.find("[name=pprice]").focus()
	})
	
	$("[name=cancel]").click(function(){
		$(this).hide()
		var root = $(this).parent().parent()
		root.find("[name=commit]").hide()
		root.find("[name=change]").show()
		root.find("[name=pprice]").val(root.find("[name=pprice]").attr("copy"))
		root.find("[name=pprice]").attr("disabled", "disabled")
	})
	
	$("[name=commit]").click(function(){
		$(this).hide()
		var root = $(this).parent().parent()
		
		var price = parseInt(root.find("[name=pprice]").val())
		if (isNaN(price)) {
			alert("Invalid price")
			root.find("[name=pprice]").focus()
			return
		}
		
		$.post("discount", {
			id: parseInt(root.find("[name=id]").val()),
			pprice: price
		}, function(data, status){
			if (data == "execute success") {
				root.find("[name=change]").show()
				root.find("[name=cancel]").hide()
				root.find("[name=pprice]").attr("disabled", "disabled")
				root.find("[name=pprice]").attr("copy", root.find("[name=pprice]").val())
			}
			alert(data)
		})
	})
	
}

$(document).ready(function(){
	
	$("#discount").click(function(){
		var discount = parseFloat($("#discountinput").val())
		if (isNaN(discount) || discount<0.01) {
			alert("error discount")
			return
		}
		$.post("alldiscount", {discount:discount}, function(data, status) {
			location.reload(true)
			alert(data)
		})
	})
	
	$.post("getdiscount", {}, function(data, result) {
		Opr.result(data, result, template)
	})
	
})