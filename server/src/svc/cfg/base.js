
// 抓取全局配置
function FetchConfig( ){
	
	var Config = {}	

	Config.Channel = new Array( )
	
	// 添加信道定义
	for( var name in Channel ){
		
		var c = Channel[name]
				
		c.Name = name
		
		Config.Channel.push(c)
	}
	
	
	Config.Peer = new Array( )
	// 添加端定义
	for( var name in Peer ){
		
		var p = Peer[name]
				
		p.Name = name
		
		var channel = Channel[name]

		if ( channel == null ){
			console.log("peer name not found in channel: " + name )
			continue
		}

		// 将Channel的地址直接拷贝到这里
		// 如果提前被设定了值, 使用设定的值
		if( p.Address == undefined ){
			p.Address = channel.Address				
		}
		
		if(p.NotifyAddress == undefined){
			p.NotifyAddress = channel.NotifyAddress
		}
		
		//console.log( name + "," + p.Address )
	
		Config.Peer.push(p)
	}
	
	Config.DB = DB
	Config.DebugRoute = DebugRoute
	Config.Log = Log
	Config.Prof = Prof

	if( DB.GroupID > 4096 ) {
		panic("zoneid is more than 4096")
	}
	
	return JSON.stringify(Config)
}

// 是否开启BI功能
function EnableBI(){
	if ( SvcName == "gamesvc" ){
		Peer["game<-client"].Component.push("bi_yd")
	}
	else if ( SvcName == "loginsvc" ){
		Peer["login<-game"].Component.push("bi_yd")
	}
	
}

// 取全局的一个配置
function GetConfig( name ){
	var v = this[name]
	if (v != null) {
		return JSON.stringify(v)
	}
		
	return null
}

function setProperty( obj, name, value ){
	if ( obj.hasOwnProperty(name)  ){	
		obj[name] = value
	}else{
		Object.defineProperty(obj, name, value)
	}
}

// 偏移一个Channel的端口
function OffsetPeerPort( name, offset ){
	
	var p = Peer[name]
	if ( p == null )
		return
		
	// connector偏移	
	if ( p.Type == "connector"	 )
	{
		return
	}
	
	// 默认地址取自Channel
	if ( p.Address == undefined )
	{
		p.Address = Channel[name].Address
	}
	
	p.Address = PortOffset( p.Address, offset )
	
	//console.log("OffsetPort", name, offset )	
} 

// 替换一个channel的ip
function ReplacePeerIP( name, ip ){
	
	var p = Peer[name]
	if ( p == null )
		return
	
	// 默认地址取自Channel
	if ( p.Address == undefined )
	{
		p.Address = Channel[name].Address
	}
	
	p.Address = ReplaceIP( p.Address, ip )	
	
//	console.log("ReplaceIP", name, p.Address   )	
}