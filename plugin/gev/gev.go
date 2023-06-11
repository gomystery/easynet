package gev

import (
	"context"
	"fmt"
	"log"

	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
	"github.com/Allenxuxu/gev"

)


type GevServer struct {
	Ctx context.Context
	addr      string
	multicore bool

	handler _interface.IEasyNet
}

func NewGevServer(ctx context.Context, config *base.NetConfig, handler _interface.IEasyNet) *GevServer {
	return &GevServer{
		Ctx:       ctx,
		addr:      fmt.Sprintf("%s://%s:%d", config.GetProtocol(), config.GetIp(), config.GetPort()),
		multicore: false,
		handler:   handler,
	}
}

func (s *GevServer) OnConnect(c *gev.Connection) {
	log.Printf("Gnet server with multi-core=%t is listening on %s\n", s.multicore, s.addr)
	s.handler.OnConnect(c)
	return
}


func (s *GevServer) OnMessage(c *gev.Connection, ctx interface{}, data []byte) (out interface{}) {
	s.handler.OnReceive(c,data)
	return
}


func (s *GevServer) OnClose(c *gev.Connection) {
	s.handler.OnClose(c,nil)
	return
}
