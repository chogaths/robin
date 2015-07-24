package main

import (
	_ "core/binarynet"
	_ "core/component/rpc"
	_ "core/martinihttp"
	"core/netfw"
	_ "svc/component/opr/discount"
	_ "svc/component/opr/icloud"
	_ "svc/component/opr/login"
	_ "svc/component/opr/prevseasonrank"
	_ "svc/component/opr/userdata"
)

func main() {

	netfw.SkipCommandLine = true

	netfw.Start()

	netfw.WaitForExit()

}
