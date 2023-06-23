package evio

import (
	"context"
	"fmt"
	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
	"github.com/tidwall/evio"
	"log"
)

type EvioServer struct {
	Ctx context.Context

	addr      string
	multicore bool

	handler _interface.IEasyNet
}

func NewGnetServer(ctx context.Context, config *base.NetConfig, handler _interface.IEasyNet) *EvioServer {
	return &EvioServer{
		Ctx:       ctx,
		addr:      fmt.Sprintf("%s://%s:%d", config.GetProtocol(), config.GetIp(), config.GetPort()),
		multicore: false,
		handler:   handler,
	}
}

func (s EvioServer) Run() error {
	var events evio.Events
	events.NumLoops = 0
	events.Serving = func(srv evio.Server) (action evio.Action) {
		log.Printf("evio server OnConnect")
		err := s.handler.OnConnect(nil)
		if err != nil {
			log.Printf("evio server OnConnect error %v", err)
		}
		return
	}
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		out, err := s.handler.OnReceive(c, in)
		if err != nil {
			log.Printf("evio server OnReceive err %v", err)
		}
		return
	}

	err := evio.Serve(events, s.addr)
	if err != nil {
		log.Printf("evio Serve error %v", err)
		return err
	}
	return nil
}
