Date.prototype.Format = function(fmt) {

	var style = {   
		"M+" : this.getMonth()+1,                 //month
		"d+" : this.getDate(),                    //day
		"h+" : this.getHours(),                   //hour
		"m+" : this.getMinutes(),                 //minute
		"s+" : this.getSeconds(),                 //second
		"q+" : Math.floor((this.getMonth()+3)/3), //season
		"S"  : this.getMilliseconds(),            //millisecond 
	}
	
	if (/(y+)/.test(fmt)) {
		fmt = fmt.replace(RegExp.$1, (this.getFullYear()+"").substr(4-RegExp.$1.length))
		for (var key in style) {
			if (new RegExp("("+key+")").test(fmt)) {
				fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? 
					style[key] : ("00"+style[key]).substr((""+style[key]).length))
			}
		}
	}
	
	return fmt
	
}

Date.prototype.String = function() {
	return this.Format("yyyy-MM-dd hh:mm:ss")
}

Date.GetSeconds = function() {
	return Math.floor(new Date().getTime()/1000)
}

// format utc second to string like "yyyy-MM-dd hh:mm:ss"
Date.ConvertToString = function() {
	
	var second = 0
	
	for (var i=0; i<arguments.length; i++){
		second += parseInt(arguments[i])
	}
	
	if (second < 0) {
		return "--"
	}
	
	return new Date(second*1000).String()

}

// verify time format like "yyyy-MM-dd hh:mm:ss"
Date.VerifyDateString = function(string) {
	
	var reg = /^(\d{1,4})-(\d{1,2})-(\d{1,2}) (\d{1,2}):(\d{1,2}):(\d{1,2})$/
	
	var r = string.match(reg)
	if (r == null) {
		return false
	}
	
	var date = new Date(r[1], r[2], r[3], r[4], r[5], r[6])
	
	return date.getFullYear() == r[1]
		&& (date.getMonth()) == r[2]
		&& date.getDate() == r[3]
		&& date.getHours() == r[4]
		&& date.getMinutes() == r[5]
		&& date.getSeconds() == r[6]
}

// convert string like "yyyy-MM-dd hh:mm:ss" to utc second
Date.ConvertToSecond = function(string) {

	if (!Date.VerifyDateString(string)) {
		return NaN
	}

	var date = new Date(string)
	return Math.floor(date.getTime()/1000)

}