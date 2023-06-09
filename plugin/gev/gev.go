package gev

import (
	"context"
	"fmt"
	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/toolkit/sync/atomic"
	"log"

	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"
	"github.com/panjf2000/gnet/v2"

	//"github.com/Allenxuxu/gev/log"
	//"time"
)

type example struct {
	Count atomic.Int64
}

func (s *example) OnConnect(c *gev.Connection) {
	s.Count.Add(1)
	//log.Println(" OnConnect ： ", c.PeerAddr())
}
func (s *example) OnMessage(c *gev.Connection, ctx interface{}, data []byte) (out interface{}) {
	//log.Println("OnMessage")
	out = data
	return
}

func (s *example) OnClose(c *gev.Connection) {
	s.Count.Add(-1)
	//log.Println("OnClose")
}

type GevServer struct {
	Ctx context.Context
	gnet.BuiltinEventEngine

	eng       gnet.Engine
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
