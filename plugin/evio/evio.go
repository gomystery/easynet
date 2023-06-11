package evio

import (
	"context"
	"fmt"
	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
	"github.com/tidwall/evio"

	//"time"
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
		fmt.Println("evio",s.addr)
		s.handler.OnStart(nil)
		return
	}
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		s.handler.OnReceive(c,in)

		out = in
		return
	}

	err:=evio.Serve(events, s.addr)
	if err != nil {
		return err
	}
	//fmt.Println("evio",s.addr)
	return nil
}
