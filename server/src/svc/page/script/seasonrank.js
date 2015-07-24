function assembledata(obj) {
	
	var devide = $("#rank")
	devide.empty()
	
	devide.append(
		"<tr> <th>ID</th> <th>T_SID</th> <th>T_IMEI</th> <th>AREA_ID</th> <th>LAST_RANK</th> <th>STATUS</th> </tr>"
	)
	
	for (var index in obj) {
		var data = obj[index]
		devide.append(
		"<tr>" + 
		"<td> <input name=id style=width:80px;text-align:center disabled value=" +data.ID + " /> </td>" +
		"<td> <input name=id style=width:120px;text-align:center disabled value=" +data.T_sid + " /> </td>" +
		"<td> <input name=id style=width:120px;text-align:center disabled value=" +data.T_imei + " /> </td>" +
		"<td> <input name=id style=width:80px;text-align:center disabled value=" +data.Area_id + " /> </td>" +
		"<td> <input name=id style=width:80px;text-align:center disabled value=" +data.Last_rank + " /> </td>" +
		"<td> <input name=id style=width:80px;text-align:center disabled value=" +data.Status + " /> </td>" +
		"</tr>"
		)
	}
	
	$("#select").removeAttr("disabled")
	$("#info").text("当前选择："+$("#select option:selected").text())
	
}

$(document).ready(function() {
	$("#commit").click(function(){
		var zone = $("#select").val()
		$("#select").attr("disabled", "disabled")
		$.post("rankdata", {zone:zone}, function(data, status){
			Opr.result(data, status, assembledata)
		})
	})
})