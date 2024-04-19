package ip

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func GetIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网卡信息出错:", err)
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("获取IP地址出错", err)
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil && strings.HasPrefix(ipNet.IP.String(), "192.168") {
					return ipNet.IP.String(), nil
				}
			}
		}
	}
	return "", errors.New("获取IP地址出错")
}
