package util

import (
	"testing"
)

func TestStrHash(t *testing.T) {
	if v := StringCRCNoCase("gamedef.EnterGameREQ"); v != 0x47c9ce66 {
		t.Errorf("expect 0x47c9ce66, got %x", v)
	}

	if v := StringCRCNoCase("gamedef.EnterGameACK"); v != 0x2c933204 {
		t.Errorf("expect 0x2c933204, got %x", v)
	}
}
