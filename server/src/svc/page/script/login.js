// http://www.cnblogs.com/mofish/archive/2012/02/25/2367858.html
function enc( )
{
	var original = document.getElementById("password").value
	document.getElementById("password").value = hex_sha1(original)
	
	return true
}
