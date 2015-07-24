package netfw

import (
	"testing"
)

func TestSvcID(test *testing.T) {

	svcid := MakeServiceID("acceptor", "game<-client", 2)
	if svcid != "acceptor|game<-client#2" {
		test.Fail()
	}

	t, name, svcindex := ParseServiceID(svcid)

	if t != "acceptor" || name != "game<-client" || svcindex != 2 {
		test.Fail()
	}
}
