package main
 
import (
    "flag"
    "fmt"
    "net"
)
 
const version = "0.1.0"
 
var Input_protocol = flag.String("p", "tcp", "The protocol you want to check")
 
func tcp(url string) int {
    _, err := net.Dial("tcp", url)
    if err != nil {
        fmt.Println(err)
        return 端口关闭
    } else {
        return 端口开启
    }
}
 
func udp(url string) int {
    _, err := net.Dial("udp", url)
    if err != nil {
        fmt.Println(err)
        return 端口关闭
    } else {
        return 端口开启
    }
}
 
func main() {
    flag.Parse()
    if flag.NArg() < 1 {
        useage := "使用示例: check -p tcp 192.168.7.26:22 或者 check -p udp 192.168.7.23:123 "
        fmt.Println(useage)
        return
    }
    p := *Input_protocol
    switch {
    case p == "tcp":
        r := tcp(flag.Args()[0])
        fmt.Println(r)
    case p == "udp":
        r := udp(flag.Args()[0])
        fmt.Println(r)
    }
}
