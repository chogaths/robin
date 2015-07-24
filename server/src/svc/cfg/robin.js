require("base")
require("global")

Peer = { 

	"opr<-web": {
		Type: "acceptor"
		Implementor: "martinihttp"
		
		TemplateDir: "../page/templates"
		StaticFileDir: "../page/script"
		
		Component: [
			"login",
		]
	},
	
}

OperateConfig = {
	
	Account: [
		{
			Account: "liuzhuorui",
			Password: "111589"
		},
		{
			Account: "root",
			Password: "123"
		},
	]
	
}

DB.Enable = false

require("local")