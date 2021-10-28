package main

import "strings"

func GetIpFromRemoteAddr(remoteAddr string) string {
	if strings.Contains(remoteAddr, ":") {
		return strings.Split(remoteAddr, ":")[0]
	} else {
		return remoteAddr
	}
}
