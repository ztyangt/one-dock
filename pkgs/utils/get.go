package utils

import (
	"net"
	"os"

	"github.com/spf13/cast"
)

// Get - 获取
var Get *GetClass

type GetClass struct{}

// Mac - 获取本机MAC
func (g *GetClass) Mac() (result string) {

	interfaces, err := net.Interfaces()

	if err != nil {
		return ""
	}

	for _, item := range interfaces {
		// 过滤掉非物理接口类型
		if item.Flags&net.FlagUp != 0 && item.Flags&net.FlagLoopback == 0 && item.Flags&net.FlagPointToPoint == 0 {
			if value, err := item.Addrs(); err == nil {
				for _, val := range value {
					if IPNet, ok := val.(*net.IPNet); ok && !IPNet.IP.IsLoopback() {
						if mac := item.HardwareAddr; len(mac) > 0 {
							return cast.ToString(mac)
						}
					}
				}
			}
		}
	}

	return ""
}

// Pid - 获取进程ID
func (g *GetClass) Pid() (result int) {
	process, err := os.FindProcess(os.Getpid())
	if err != nil {
		return 0
	}
	return process.Pid
}

// Pwd - 获取当前目录
func (g *GetClass) Pwd() (result string) {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}
