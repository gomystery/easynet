package evio

import (
	"context"
	"fmt"
	"github.com/gomystery/easynet/interface"
	"github.com/tidwall/evio"
	"log"
)

type EvioServer struct {
	Ctx context.Context

	config *YamlConfig
	addr      string

	handler _interface.IEasyNet
}

func NewEvioServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *EvioServer {

	s :=&EvioServer{
		Ctx:       ctx,
		handler:   handler,
		config:config,
	}
	if s.config != nil {
		s.addr = s.getAddr()
	}
	return s
}

func (s EvioServer) Run() error {
	var events evio.Events
	events.NumLoops = int(s.config.GetLoops())
	events.Opened = func(c evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
		log.Printf("evio Opened OnConnect")
		err := s.handler.OnConnect(c)
		if err != nil {
			log.Printf("evio server OnConnect error %v", err)
		}
		return
	}

	events.Serving = func(srv evio.Server) (action evio.Action) {
		log.Printf("evio server OnStart")
		err := s.handler.OnStart(nil)
		if err != nil {
			log.Printf("evio server OnStart error %v", err)
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
	events.Closed = func(c evio.Conn, inErr error) (action evio.Action) {
		log.Printf("evio Opened OnClose")
		err := s.handler.OnClose(c,inErr)
		if err != nil {
			log.Printf("evio server OnClose error %v", err)
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

func (s EvioServer) getAddr() string {
	if s.config.GetStdlib() {
		ssuf := "-net"
		return fmt.Sprintf("%s%s://%s:%d?reuseport=%t", s.config.GetProtocol(),ssuf, s.config.GetIp(), s.config.GetPort(),s.config.GetReuseport())
	}
	return fmt.Sprintf("%s://%s:%d?reuseport=%t", s.config.GetProtocol(), s.config.GetIp(), s.config.GetPort(),s.config.GetReuseport())
}