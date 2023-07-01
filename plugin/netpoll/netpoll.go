package netpoll

import (
	"context"
	"fmt"
	"github.com/cloudwego/netpoll"
	"github.com/gomystery/easynet/interface"
	"log"
)

type NetPollServer struct {
	Ctx context.Context

	Network   string
	Address   string
	multicore bool

	handler _interface.IEasyNet
}

func NewNetPollServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *NetPollServer {
	return &NetPollServer{
		Ctx:       ctx,
		Network:   config.GetProtocol(),
		Address:   fmt.Sprintf("%s:%d", config.GetIp(), config.GetPort()),
		multicore: false,
		handler:   handler,
	}

}

func (s *NetPollServer) Run() error {

	var eventLoop netpoll.EventLoop

	listener, err := netpoll.CreateListener(s.Network, s.Address)
	if err != nil {
		fmt.Println("create netpoll listener failed")
		return err
	}

	//type OnRequest func(ctx context.Context, connection Connection) error
	handle := func(ctx context.Context, connection netpoll.Connection) error {
		var b []byte
		connection.Read(b)
		bytes, err := s.handler.OnReceive(connection, b)
		if err != nil {
			log.Printf("Gev server OnReceive ,err=$v \n", err)
		}
		connection.Write(bytes)
		return err
	}
	// todo is right
	prepare := func(connection netpoll.Connection) context.Context {
		fmt.Println(connection)
		s.handler.OnStart(connection)
		return s.Ctx
	}

	//type OnConnect func(ctx context.Context, connection Connection) context.Context
	connect := func(ctx context.Context, connection netpoll.Connection) context.Context {
		log.Printf("Gev server OnConnect \n")
		err := s.handler.OnConnect(connection)
		if err != nil {
			log.Printf("Gev server OnConnect ,err=$v \n", err)
		}
		return ctx
	}

	eventLoop, _ = netpoll.NewEventLoop(
		handle,
		netpoll.WithOnPrepare(prepare),
		netpoll.WithOnConnect(connect),
		//netpoll.WithReadTimeout(time.Second),
	)

	// start listen loop ...
	eventLoop.Serve(listener)

	return nil
}
