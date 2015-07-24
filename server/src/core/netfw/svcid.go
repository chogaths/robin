package netfw

import (
	"fmt"
	"strconv"
	"strings"
)

func MakeServiceID(peertype string, name string, peerIndex int32) string {
	return fmt.Sprintf("%s|%s#%d", peertype, name, peerIndex)
}

func ParseServiceID(svcid string) (peertype string, name string, peerIndex int32) {
	slicepos := strings.Index(svcid, "|")

	peertype = svcid[:slicepos]

	sharppos := strings.Index(svcid, "#")

	name = svcid[slicepos+1 : sharppos]

	peerstr := svcid[sharppos+1:]

	index, _ := strconv.Atoi(peerstr)

	peerIndex = int32(index)

	return
}
