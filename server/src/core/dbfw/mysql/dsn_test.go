package dbfw

import (
	"testing"
)

func TestDSNParse(t *testing.T) {

	usr, passwd, nettype, addr, dbname, _ := parseDSN("root:passwd@tcp(127.0.0.1:3307)/mydb")

	if usr != "root" ||
		passwd != "passwd" ||
		nettype != "tcp" ||
		addr != "127.0.0.1:3307" ||
		dbname != "mydb" {
		t.Fail()
	}
}
