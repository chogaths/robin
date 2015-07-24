function assembleempty() {
	devide = $("#icloud");
	devide.empty();
	devide.append("无数据");
}

function assembleinfo(obj) {
	
	devide = $("#icloud");
	devide.empty();
	devide.append("<tr> <th>id</th> <th>imei</th> <th>email</th> <th>password</th> <th>restoretimes</th> <th>enable</th> </tr>");
	
	for (var index in obj) {	
		var data = obj[index]	
		devide.append(
			"<tr id="+data.Id+">" +
				"<td> <input name=id style=width:80px;text-align:center disabled value=" + data.Id + " /> </td>" +
				"<td> <input name=imei style=width:120px;text-align:center disabled value='" + data.Imei + "' /> </td>" +
				"<td> <input name=email style=width:120px;text-align:center disabled value=" + data.Email + " /> </td>" +
				"<td> <input name=password style=width:280px;text-align:center disabled copy=" + data.Password + " value=" + data.Password + " /> </td>" +
				"<td> <input name=restore style=width:280px;text-align:center disabled copy=" + data.Remainder_restore_times + " value=" + data.Remainder_restore_times + " /> </td>" +
				"<td> <input name=enable style=width:120px;text-align:center disabled value=" + data.Enable + " /> </td>" +
				"<td> <button name=change onclick=change("+data.Id+")>修改</button> </td>" +
				"<td> <button name=commit onclick=commit("+data.Id+") hidden>确认</button> </td>" +
				"<td> <button name=cancel onclick=cancel("+data.Id+") hidden>取消</button> </td>" +
				"<td> <button onclick=deleteicloud("+data.Id+")>删除</button> </td>" +
			"</tr>"
		);		
	}
	
}

function change(id) {
	$("#"+id).find("[name=change]").hide()
	$("#"+id).find("[name=commit]").show()
	$("#"+id).find("[name=cancel]").show()
	$("#"+id).find("[name=password]").removeAttr("disabled")
	$("#"+id).find("[name=restore]").removeAttr("disabled")
}

function cancel(id) {
	$("#"+id).find("[name=change]").show()
	$("#"+id).find("[name=commit]").hide()
	$("#"+id).find("[name=cancel]").hide()
	$("#"+id).find("[name=password]").attr("disabled", "disabled")
	$("#"+id).find("[name=password]").val($("#"+id).find("[name=password]").attr("copy"))
	$("#"+id).find("[name=restore]").attr("disabled", "disabled")
	$("#"+id).find("[name=restore]").val($("#"+id).find("[name=restore]").attr("copy"))
}

function commit(id) {
	var cpass = $("#"+id).find("[name=password]")
	var password = cpass.val()
	if (password == cpass.attr("copy")) {
		password = ""
	}
	var crestore = $("#"+id).find("[name=restore]")
	var restore = crestore.val()
	$.post("updateicloud", {id:id, password:password, restore:restore}, function(data, status) {
		alert(data.Result)
		if (data.Result == "execute success") {
			$("#"+id).find("[name=change]").show()
			$("#"+id).find("[name=commit]").hide()
			$("#"+id).find("[name=cancel]").hide()
			if (data.Password != "") {
				$("#"+id).find("[name=password]").val(data.Password)
			}
			$("#"+id).find("[name=password]").attr("disabled", "disabled")
			crestore.attr("disabled", "disabled")
		}
	})
}

function deleteicloud(id) {
	$.post("icloud", {id:id}, function(data, status) {
		alert(data)
		if (data == "execute success") {
			$("#"+id).hide()
		}
	})
}

function postquery(query) {
	
	$.post("queryicloud", query, function(data, status) {
		Opr.result(data, status, assembleinfo);
		if ("page" in query) {
			$("#pageinfo").attr("cur", query.page)
			$("#pageinfo").text("当前:"+(query.page==0 ? 1:query.page)+"页 共:"+$("#pageinfo").attr("page")+"页")
		} else {
			$("#pageinfo").attr("cur", 1)
		}
	});

}

$(document).ready(function(){
	
	assembleempty();
	
	if (parseInt($("#icloud").attr("page")) == 0) {
		assembleempty();
	} else {
		postquery({page:0})
	}
	
	$("#searchid").click(function(){
		var id = parseInt($("#idinput").val())
		if (isNaN(id)) {
			alert("invalid id")
			return
		}
		postquery({id:id})
	})
	
	$("#searchimei").click(function(){
		var imei = $("#imeiinput").val()
		if (imei == "") {
			alert("empty imei")
			return
		}
		postquery({imei:imei})
	})
	
	$("#searchemail").click(function(){
		var email = $("#emailinput").val()
		if (email == "") {
			alert("empty email")
			return
		}
		postquery({email:email})
	})
	
	$("#searchpage").click(function(){
		var page = parseInt($("#pageinput").val())
		if (isNaN(page)) {
			alert("invalid page")
			return
		}
		postquery({page:page-1})
	})
	
})