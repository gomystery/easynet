package net

import (
	"context"
	"fmt"
	"github.com/baickl/logger"
	"github.com/gomystery/easynet/interface"
	"net"
)

type NetServer struct {
	Ctx context.Context

	Network   string
	Address   string
	multicore bool

	handler _interface.IEasyNet
}

func NewNetServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *NetServer {
	return &NetServer{
		Ctx:       ctx,
		Network:   config.GetProtocol(),
		Address:   fmt.Sprintf("%s:%d", config.GetIp(), config.GetPort()),
		multicore: false,
		handler:   handler,
	}

}

func (s *NetServer) Run() error {
	ln, err := net.Listen(s.Network, s.Address)
	if err != nil {
		return err
	}
	if err := s.handler.OnStart(nil); err != nil {
		logger.Errorf("net OnStart err %v", err)
		return err
	}
	logger.Infoln("net OnStart")

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
		logger.Infoln("net OnConnect")

		go s.handleConnection(conn)
	}
	return nil
}

func (s *NetServer) handleConnection(conn net.Conn) {
	//函数调用完毕，自动关闭conn
	defer conn.Close()

	//4、获取客户端的网络地址信息
	rbuf, wbuf := []byte{}, []byte{}

	//5、获取用户数据
	for {
		_, err := conn.Read(rbuf)
		if err != nil {
			logger.Errorf("net read message err %v", err)
			return
		}

		if wbuf, err = s.handler.OnReceive(conn, rbuf); err != nil {
			logger.Errorf("net OnReceive err %v", err)
			return
		}

		//6、给用户发送回去
		conn.Write(wbuf)
	}
}
