package net

import (
	"context"
	"fmt"
	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"
	"net"
)


type NetServer struct {
	Ctx context.Context

	Network      string
	Address string
	multicore bool

	handler _interface.IGmtNet
}

func NewNetPollServer(ctx context.Context, config *base.NetConfig, handler _interface.IGmtNet) *NetServer {
	return &NetServer{
		Ctx:       ctx,
		Network:   config.GetProtocol(),
		Address:   fmt.Sprintf("%s:%d",config.GetIp(), config.GetPort()),
		multicore: false,
		handler:   handler,
	}

}
 
func (s *NetServer) Run() error {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	if err:= s.handler.OnStart(nil); err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		if err = s.handler.OnConnect(conn); err != nil {
			// handle error
			continue
		}

		go s.handleConnection(conn)
	}
	return nil
}

func (s *NetServer) handleConnection(conn net.Conn) {
	//函数调用完毕，自动关闭conn
	defer conn.Close()

	//4、获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "连接成功！！！")
	rbuf,wbuf := []byte{},[]byte{}

	//5、获取用户数据
	for {
		_, err := conn.Read(rbuf)
		if err != nil {
			fmt.Println("获取数据错误！！！")
			return
		}

		if wbuf,err = s.handler.OnReceive(conn,rbuf); err != nil {
			return
		}



		//6、把数据转换成大写，再给用户发送回去
		conn.Write(wbuf)
	}
}
