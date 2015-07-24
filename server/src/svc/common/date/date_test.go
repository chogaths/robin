package date

import (
	"testing"
	"time"
)

func TestCrossDay(t *testing.T) {

	// 采用当前时区
	location := time.Now().Location()

	// 从午夜零点到给定时间点的时间差
	pointOffset, _ := time.ParseDuration("5h")

	// 过去某个点
	UTCPast := time.Date(2015, 1, 21, 23, 14, 0, 0, location).UTC().Unix()

	// 未跨越给定点
	UTCT1 := time.Date(2015, 1, 22, 4, 14, 0, 0, location).UTC().Unix()

	// 跨越1次
	UTCT2 := time.Date(2015, 1, 22, 5, 14, 0, 0, location).UTC().Unix()

	// 跨越2次
	UTCT3 := time.Date(2015, 1, 23, 14, 20, 0, 0, location).UTC().Unix()

	if GetCrossDayCount(UTCPast, UTCT1, pointOffset, 365) != 0 {
		t.Fail()
	}

	if GetCrossDayCount(UTCPast, UTCT2, pointOffset, 365) != 1 {
		t.Fail()
	}

	if GetCrossDayCount(UTCPast, UTCT3, pointOffset, 365) != 2 {
		t.Fail()
	}
}
