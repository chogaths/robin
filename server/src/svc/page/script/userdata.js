function assembelUserData(obj) {
	
	$("#info").attr("zone", obj.Zone)
	$("#info").attr("service", obj.Service)
	$("#info").text("当前选择："+$("#zone option[value='"+obj.Zone+"']").text()+" - "+$("#"+obj.Zone+" option[value='"+obj.Service+"']").text())
	
	var devide = $("#data")
	devide.empty()
	
	if (obj.Data != null) {
		devide.append("<tr> <th>T_SID</th> <th>T_IMEI</th> <th>T_NICKNAME</th> <th>T_VICTORY</th> <th>T_YUXI_VALUE</th>" + (obj.Zone == "db_td2ios" ? "<th>T_PASSWORD</th>" : "") + "</tr>")
		
		for (var index in obj.Data) {
			var data = obj.Data[index]
			devide.append(
				"<tr name=usertsid_" + data.T_sid + " >" +
				"<td> <input name=sid style=width:80px;text-align:center disabled value=" + data.T_sid + " /> </td>" +
				"<td> <input name=imei style=width:160px;text-align:center disabled copy=" + data.T_imei + " value=" + data.T_imei + " /> </td>" +
				"<td> <input name=nickname style=width:120px;text-align:center disabled copy=" + data.T_nickname + " value=" + data.T_nickname + " /> </td>" +
				"<td> <input name=victory style=width:120px;text-align:center disabled copy=" + data.T_victory + " value=" + data.T_victory + " /> </td>" +
				"<td> <input name=yuxival style=width:120px;text-align:center disabled copy=" + data.T_yuxi_value + " value=" + data.T_yuxi_value + " /> </td>" +
				(obj.Zone == "db_td2ios" ? "<td> <input name=password style=width:240px;text-align:center disabled copy=" + data.T_password + " value=" + data.T_password + " /> </td>" : "") +
				"<td> <button name=change>修改</button> </td>" +
				"<td> <button name=commit hidden>确定</button> </td>" +
				"<td> <button name=cancel hidden>取消</button> </td>" +
				"<td> <button name=delete>删除</button> </td>" +
				"</tr>"
			)
		}
	} else {
		devide.append("无数据")
	}
	
	$("#pageinfo").text("当前:"+obj.Page+"页 共:"+obj.MaxPage+"页    当页数量:" + obj.Data.length)
	
	$("#search").show()
	
	register()
	
}

function register() {
	
	$("button[name='change']").click(function(){
		$(this).hide()
		var parent = $(this).parent().parent()
		parent.find("[name='commit']").show()
		parent.find("[name='cancel']").show()
		//parent.find("[name='imei']").removeAttr("disabled")
		//parent.find("[name='nickname']").removeAttr("disabled")
		parent.find("[name='victory']").removeAttr("disabled")
		parent.find("[name='yuxival']").removeAttr("disabled")
		parent.find("[name='password']").removeAttr("disabled")
	})
	
	$("button[name='cancel']").click(function(){
		$(this).hide()
		var parent = $(this).parent().parent()
		parent.find("[name='commit']").hide()
		parent.find("[name='change']").show()
		var elements = Array(parent.find("[name='victory']"), parent.find("[name='yuxival']"), /*parent.find("[name='imei']"), parent.find("[name='nickname']"),*/ parent.find("[name='password']"))
		for (var index in elements) {
			elements[index].val(elements[index].attr("copy"))
			elements[index].attr("disabled", "disabled")
		}
	})
	
	$("button[name='commit']").click(function(){		
		var ctlc = $(this)
		var parent = $(this).parent().parent()
		var ctlv = parent.find("[name='victory']")
		var ctly = parent.find("[name='yuxival']")
		var ctli = parent.find("[name='imei']")
		var ctln = parent.find("[name='nickname']")
		var ctlp = parent.find("[name='password']")
		
		var victory = parseInt(ctlv.val())
		if (isNaN(victory)) {
			alert("error victory value")
			return
		}
		ctlv.val(victory)
		
		var yuxival = parseInt(ctly.val())
		if (isNaN(yuxival)) {
			alert("error yuxi value")
			return
		}
		ctly.val(yuxival)
		
		var password = ""
		if (ctlp) {
			password = ctlp.val();
			if (password == ctlp.attr("copy")) {
				password = "";
			}
			if (password != "" && !confirm("modify password "+password+" ?")) {
				return
			}
		}
		
		$.post("userdata", {
			t_sid: parseInt(parent.find("[name='sid']").val()),
			//t_imei: ctli.val(),
			//t_nickname: ctln.val(),
			t_victory: victory,
			t_yuxi_value: yuxival,
			t_password: password,
			zone: $("#info").attr("zone"),
			service: parseInt($("#info").attr("service"))
		}, function(data, status) {
			if (typeof(data) == "string") {
				if (data == "execute success") {
					ctlc.hide()
					parent.find("[name='change']").show()
					parent.find("[name='cancel']").hide()
					ctlv.attr("disabled", "disabled")
					ctly.attr("disabled", "disabled")
					ctli.attr("disabled", "disabled")
					ctln.attr("disabled", "disabled")
					ctlp.attr("disabled", "disabled")
				}
				alert(data)
			} else {
				var pass = $("#data").find("[name=usertsid_"+data.T_sid+"]").find("[name=password]")
				pass.val(data.T_password)
				pass.attr("copy", data.T_password)
				ctlc.hide()
				parent.find("[name='change']").show()
				parent.find("[name='cancel']").hide()
				ctlv.attr("disabled", "disabled")
				ctly.attr("disabled", "disabled")
				ctli.attr("disabled", "disabled")
				ctln.attr("disabled", "disabled")
				ctlp.attr("disabled", "disabled")
				alert("execute success")
			}
		})
	})
	
	$("button[name='delete']").click(function(){
		var parent = $(this).parent().parent()
		$.post("deleteuser", {
			t_sid: parseInt(parent.find("[name='sid']").val()),
			zone: $("#info").attr("zone"),
			service: parseInt($("#info").attr("service"))
		}, function(data, status) {
			if (data == "execute success") {
				parent.hide()
			}
			alert(data)
		})
	})
	
}

function updatezoneselect() {	
//	$("#db_td2cn").hide()
//	$("#db_td2ios").hide()
//	$("#db_td2wp").hide()
	$("#zonelist").children().each(function(){$(this).hide()})
	$("#"+$("#zone").val()).show()	
}

function queryuserdata(info) {
	$.post("queryuserdata", info, function(data, status){
		Opr.result(data, status, function(obj){
			assembelUserData(obj)
		})
	})
}

$(document).ready(function() {
	
	updatezoneselect()
	$("#zone").change(updatezoneselect)
	
	$("#commit").click(function(){
		queryuserdata({zone:$("#zone").val(), service: parseInt($("#"+$("#zone").val()).val())})
	})
	
	$("#searchid").click(function(){
		var id = parseInt($("#idinput").val())
		if (isNaN(id)) {
			alert("invalid id")
			return
		}
		queryuserdata({zone:$("#zone").val(), service: parseInt($("#"+$("#zone").val()).val()), id:id})
	})
	
	$("#searchimei").click(function(){
		var imei = $("#imeiinput").val()
		if (imei == "") {
			alert("empty imei")
			return
		}
		queryuserdata({zone:$("#zone").val(), service: parseInt($("#"+$("#zone").val()).val()), imei:imei})
	})
	
	$("#searchpage").click(function(){
		var page = parseInt($("#pageinput").val())
		if (isNaN(page)) {
			alert("invalid page")
			return
		}
		queryuserdata({zone:$("#zone").val(), service: parseInt($("#"+$("#zone").val()).val()), page:page})
	})
	
})