package main

import (
	"log"

	"github.com/qtgolang/SunnyNet/SunnyNet"
	"github.com/qtgolang/SunnyNet/public"
	"github.com/spf13/pflag"
)

var (
	Sunny = SunnyNet.NewSunny()
	Port  int
)

func main() {
	// 解析命令行参数，设置代理端口
	pflag.IntVarP(&Port, "port", "p", 18080, "代理端口")
	pflag.Parse()

	log.Printf("代理端口: %d\n", Port)

	// 设置回调函数
	Sunny.SetGoCallback(HttpCallback, TcpCallback, WSCallback, UdpCallback)

	// 绑定端口号并启动服务
	s := Sunny.SetPort(Port)
	s.SetRandomTLS(true)

	if err := s.Start(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
	log.Println("服务启动成功")

	// 阻止程序退出
	select {}
}

// HttpCallback 处理 HTTP 请求的回调函数
func HttpCallback(Conn *SunnyNet.HttpConn) {
	log.Printf("HttpCallback: 请求方法: %v, 请求 IP: %v, 请求地址: %v\n", Conn.Request.Method, Conn.ClientIP, Conn.Request.URL)
}

// TcpCallback 处理 TCP 请求的回调函数
func TcpCallback(Conn *SunnyNet.TcpConn) {
	// 捕获到数据可以修改, 修改空数据, 取消发送/接收
	// Conn.SetAgent("") // 这个是针对这个一个请求的, TCP, 只能设置 S5 代理
	// log.Printf("TCP: PID: %v, LocalAddr: %v, RemoteAddr: %v, Type: %v, BodyLen: %v\n", Conn.Pid, Conn.LocalAddr, Conn.RemoteAddr, Conn.Type, Conn.GetBodyLen())
}

// WSCallback 处理 WebSocket 请求的回调函数
func WSCallback(Conn *SunnyNet.WsConn) {
	// WebSocket 回调逻辑
}

// UdpCallback 处理 UDP 请求的回调函数
func UdpCallback(Conn *SunnyNet.UDPConn) {
	// 在 Windows 捕获 UDP 需要加载驱动, 并且设置进程名
	// 其他情况需要设置 Socket5 代理, 才能捕获到 UDP
	// 捕获到数据可以修改, 修改空数据, 取消发送/接收
	switch Conn.Type {
	case public.SunnyNetUDPTypeReceive:
		// log.Printf("接收 UDP: LocalAddress: %v, RemoteAddress: %v, DataLen: %d\n", Conn.LocalAddress, Conn.RemoteAddress, len(Conn.Data))
	case public.SunnyNetUDPTypeSend:
		// log.Printf("发送 UDP: LocalAddress: %v, RemoteAddress: %v, DataLen: %d\n", Conn.LocalAddress, Conn.RemoteAddress, len(Conn.Data))
	case public.SunnyNetUDPTypeClosed:
		// log.Printf("关闭 UDP: LocalAddress: %v, RemoteAddress: %v\n", Conn.LocalAddress, Conn.RemoteAddress)
	}
}
