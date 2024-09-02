package main

import (
	"fmt"
	"net"
	"time"
)


func Check(domain string, port string, timeout_seconds int) string {
	addr := domain + ":" + port

	timeout := time.Duration(time.Duration(timeout_seconds) * time.Second)

	conn, err := net.DialTimeout("tcp", addr, timeout)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable,\nError: %v", domain, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable,\nFrom: %v\nTo: %v", domain, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}