package netpoll

import (
	"context"
	"fmt"
	"github.com/cloudwego/netpoll"
	"time"

	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"
	//"github.com/Allenxuxu/gev/log"
	//"time"
)


type NetPollServer struct {
	Ctx context.Context

	Network      string
	Address string
	multicore bool

	handler _interface.IGmtNet
}

func NewNetPollServer(ctx context.Context, config *base.NetConfig, handler _interface.IGmtNet) *NetPollServer {
	return &NetPollServer{
		Ctx:       ctx,
		Network:   config.GetProtocol(),
		Address:   fmt.Sprintf("%s:%d",config.GetIp(), config.GetPort()),
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
		bytes,err:=s.handler.OnReceive(connection,b)
		connection.Write(bytes)
		return err
	}
	prepare := func(connection netpoll.Connection) context.Context {
		fmt.Println(connection)
		s.handler.OnStart(connection)
		return s.Ctx
	}

	//type OnConnect func(ctx context.Context, connection Connection) context.Context
	connect := func(ctx context.Context, connection netpoll.Connection) context.Context {
		fmt.Println(connection)
		s.handler.OnConnect(connection)
		return s.Ctx
	}

	eventLoop, _ = netpoll.NewEventLoop(
		handle,
		netpoll.WithOnPrepare(prepare),
		netpoll.WithOnConnect(connect),
		netpoll.WithReadTimeout(time.Second),
	)

	// start listen loop ...
	eventLoop.Serve(listener)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	eventLoop.Shutdown(ctx)

	return nil
}