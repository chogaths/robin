package util

import (
	"net"
	"strconv"
)

// 端口偏移
func PortOffset(address string, offset int) string {
	host, portstr, err := net.SplitHostPort(address)

	if err != nil {
		return address
	}

	port, err := strconv.Atoi(portstr)

	if err != nil {
		return address
	}

	return net.JoinHostPort(host, strconv.Itoa(port+offset))
}

// 替换IP
func ReplaceIP(address string, ip string) string {
	_, portstr, err := net.SplitHostPort(address)

	if err != nil {
		return address
	}

	return net.JoinHostPort(ip, portstr)
}
