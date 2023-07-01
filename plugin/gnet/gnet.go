package gnet

import (
	"context"
	"fmt"
	"github.com/baickl/logger"
	"time"

	"github.com/gomystery/easynet/interface"
	"github.com/panjf2000/gnet/v2"
	//"time"
)

type GnetServer struct {
	Ctx context.Context
	gnet.BuiltinEventEngine

	eng    gnet.Engine
	addr   string
	config *YamlConfig

	handler _interface.IEasyNet
}

func NewGnetServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *GnetServer {
	server := &GnetServer{
		Ctx:     ctx,
		handler: handler,
		config:  config,
	}

	server.addr = server.getAddr()

	return server

}

func (s *GnetServer) OnBoot(eng gnet.Engine) gnet.Action {
	s.eng = eng
	logger.Infof("Gnet OnStart with multi-core=%t is listening on %s\n", s.config.Multicore, s.addr)
	s.handler.OnStart(nil)
	return gnet.None
}

func (s *GnetServer) OnShutdown(eng gnet.Engine) {
	logger.Infoln("Gnet OnShutdown")
	s.handler.OnShutdown(nil)
}

func (s *GnetServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	logger.Infoln("Gnet OnConnect")
	s.handler.OnConnect(c)

	return nil, gnet.None
}

func (s *GnetServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	logger.Infoln("Gnet OnClose")
	s.handler.OnClose(c, err)

	return gnet.None
}

func (s *GnetServer) OnTick() (delay time.Duration, action gnet.Action) {
	logger.Infoln("Gnet OnTick")
	return 0, gnet.None
}

func (s *GnetServer) OnTraffic(c gnet.Conn) gnet.Action {
	logger.Infoln("Gnet OnReceive")
	buf, _ := c.Next(-1)
	s.handler.OnReceive(c, buf)
	return gnet.None
}

func (s GnetServer) getAddr() string {
	return fmt.Sprintf("%s://%s:%d", s.config.GetProtocol(), s.config.GetIp(), s.config.GetPort())
}
