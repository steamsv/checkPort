package main

import (
	"fmt"
	"net"
	"os"
	"time"
)



// 获取IP和端口
func getIpPort() []string {
	// 根据接收参数个数，定义动态数组，
	ip_ports := make([]string, len(os.Args)-1)
	i := 0
	for index, value := range os.Args {
		//排除脚本名称
		if index == 0 {
			continue
		}
		//写入数组
		ip_ports[i] = value
		i++
	}
	return ip_ports
}

// 检测端口
func checkPorts(ip_ports []string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	for _, ip_port := range ip_ports {
		// 检测端口
		conn, err := net.DialTimeout("tcp", ip_port, 3*time.Second)
		if err != nil {
			fmt.Println("["+now+"]", ip_port, "端口未开启(fail)!")
		} else {
			if conn != nil {
				fmt.Println("["+now+"]", ip_port, "端口已开启(success)!")
				conn.Close()
			} else {
				fmt.Println("["+now+"]", ip_port, "端口未开启(fail)!")
			}
		}
	}
}

func main() {

	ret := getIpPort()
	checkPorts(ret)
}
