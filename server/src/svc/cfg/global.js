
// 端口号段分配:
// 位数对应 ABCD
// A: 内外区分, 6管理 7对内, 8对外
// B: 调试多套服务器标识, 0~9, 同样影响对外端口
// C+D: 同类型多组服务器索引标示(根据实际部署需要调整)

//  规则
//  多组服务器 发起连接连单体
//  端口号以第1台服务器端口号开始定义, 底层会自动根据SvcIndex为端口号偏移
//  SvcIndex 为base1

Channel = {
	
	"opr<-web": {		
		Address: "127.0.0.1:8080",
	}
}


DB = {
	DSN: "root:123456@tcp(127.0.0.1:3306)/testlib"
	Enable: false
	GroupID: 1		// 数据库autoid 分组
	ShowOperate: true
	ConnCount: 1
}

Log = {
	Enable: false		// linux下不要打开日志, 使用重定向
	FileName: ""		// 文件名不给定时, 默认为svcname.log
}

DebugRoute = {
	Enable: false
	BlockMsgName: [ "gamedef.OnlineUserACK" ]
}

Prof = {
	Enable: false
}