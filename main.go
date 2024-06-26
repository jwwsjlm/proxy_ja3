package main

import (
	"github.com/qtgolang/SunnyNet/SunnyNet"
	"github.com/qtgolang/SunnyNet/public"
	"github.com/spf13/pflag"
	"log"
)

var Sunny = SunnyNet.NewSunny()
var Port int

func main() {

	pflag.IntVarP(&Port, "port", "p", 18080, "代理端口")
	pflag.Parse()
	log.Println("代理端口:", Port)
	Sunny.SetGoCallback(HttpCallback, TcpCallback, WSCallback, UdpCallback)
	//绑定端口号并启动
	s := Sunny.SetPort(Port)
	s.SetRandomTLS(true)
	s.Start()
	log.Println("服务启动成功")
	select {}
}
func HttpCallback(Conn *SunnyNet.HttpConn) {
	log.Printf("HttpCallback:请求方法:%v-请求ip:%v-请求地址:%v\n", Conn.Request.Method, Conn.ClientIP, Conn.Request.URL)
}
func TcpCallback(Conn *SunnyNet.TcpConn) {
	//捕获到数据可以修改,修改空数据,取消发送/接收
	//Conn.SetAgent("") //这个是针对这个一个请求的,TCP,只能设置S5代理,
	//fmt.Println(Conn.Pid, Conn.LocalAddr, Conn.RemoteAddr, Conn.Type, Conn.GetBodyLen())
}
func WSCallback(Conn *SunnyNet.WsConn) {}
func UdpCallback(Conn *SunnyNet.UDPConn) {
	//在 Windows 捕获UDP需要加载驱动,并且设置进程名
	//其他情况需要设置Socket5代理,才能捕获到UDP
	//捕获到数据可以修改,修改空数据,取消发送/接收
	if public.SunnyNetUDPTypeReceive == Conn.Type {
		//fmt.Println("接收UDP", Conn.LocalAddress, Conn.RemoteAddress, len(Conn.Data))
	}
	if public.SunnyNetUDPTypeSend == Conn.Type {
		//fmt.Println("发送UDP", Conn.LocalAddress, Conn.RemoteAddress, len(Conn.Data))
	}
	if public.SunnyNetUDPTypeClosed == Conn.Type {
		//fmt.Println("关闭UDP", Conn.LocalAddress, Conn.RemoteAddress)
	}
}
