package gev

import (
	"context"
	"fmt"
	"github.com/Allenxuxu/gev"
	"log"

	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"

	//"github.com/Allenxuxu/gev/log"
	//"time"
)


type GevServer struct {
	Ctx context.Context
	addr      string
	multicore bool

	handler _interface.IGmtNet
}

func NewGevServer(ctx context.Context, config *base.NetConfig, handler _interface.IGmtNet) *GevServer {
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
