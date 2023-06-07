package gnet

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"
	"github.com/panjf2000/gnet/v2"
	//"time"
)

type GnetServer struct {
	Ctx context.Context
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool

	handler _interface.IGmtNet
}

func NewGnetServer(ctx context.Context, config *base.NetConfig, handler _interface.IGmtNet) *GnetServer {
	return &GnetServer{
		Ctx:       ctx,
		addr:      fmt.Sprintf("%s://%s:%d", config.GetProtocol(), config.GetIp(), config.GetPort()),
		multicore: false,
		handler:   handler,
	}
}

func (s *GnetServer) OnBoot(eng gnet.Engine) gnet.Action {
	s.eng = eng
	log.Printf("Gnet server with multi-core=%t is listening on %s\n", s.multicore, s.addr)
	s.handler.OnStart(nil)
	return gnet.None
}

func (s *GnetServer) OnShutdown(eng gnet.Engine) {
	s.handler.OnShutdown(nil)
}

func (s *GnetServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	s.handler.OnConnect(c)

	return nil, gnet.None
}

func (s *GnetServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	s.handler.OnClose(c, err)

	return gnet.None
}

func (s *GnetServer) OnTick() (delay time.Duration, action gnet.Action) {
	return 0, gnet.None
}

func (s *GnetServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	s.handler.OnReceive(c, buf)
	return gnet.None
}
