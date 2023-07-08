package netpoll

import (
	"context"
	"fmt"

	"github.com/baickl/logger"
	"github.com/cloudwego/netpoll"
	"github.com/gomystery/easynet/interface"
)

type NetPollServer struct {
	Ctx context.Context

	config *YamlConfig

	handler _interface.IEasyNet
}

func NewNetPollServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *NetPollServer {
	server := &NetPollServer{
		Ctx:     ctx,
		config:  config,
		handler: handler,
	}
	return server

}

func (s *NetPollServer) Run() error {

	var eventLoop netpoll.EventLoop

	listener, err := netpoll.CreateListener(s.config.GetProtocol(), s.getAddr())
	if err != nil {
		logger.Errorf("create netpoll listener failed err:%v", err)
		return err
	}
	err = s.handler.OnStart(nil)
	if err != nil {
		logger.Errorf("create netpoll OnStart failed err:%v", err)
		return err
	}
	logger.Infof("create netpoll OnStart,Protocol:%v ,addr:%v", s.config.GetProtocol(), s.getAddr())

	//type OnRequest func(ctx context.Context, connection Connection) error
	handle := func(ctx context.Context, connection netpoll.Connection) error {

		var reader, writer = connection.Reader(), connection.Writer()

		// reading
		buf, _ := reader.Next(reader.Len())
		reader.Release()
		//... parse the read data ...
		//var write_data []byte
		write_data, err := s.handler.OnReceive(connection, buf)
		if err != nil {
			logger.Errorf("netpoll server OnReceive ,err=$v \n", err)
			return err
		}

		// writing
		//... make the write data ...
		alloc, err := writer.Malloc(len(write_data))
		copy(alloc, write_data) // write data
		err=writer.Flush()
		if err != nil {
			logger.Errorf("netpoll server writing %s,err:%v \n", string(alloc),err)
		}
		return err
	}
	// todo is right
	prepare := func(connection netpoll.Connection) context.Context {
		logger.Infoln("netpoll server OnStart")
		s.handler.OnStart(connection)
		return s.Ctx
	}

	//type OnConnect func(ctx context.Context, connection Connection) context.Context
	connect := func(ctx context.Context, connection netpoll.Connection) context.Context {
		logger.Infoln("netpoll server OnConnect")
		err := s.handler.OnConnect(connection)
		if err != nil {
			logger.Errorf("netpoll server OnConnect ,err=$v \n", err)
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

func (s NetPollServer) getAddr() string {
	return fmt.Sprintf("%s:%d", s.config.GetIp(), s.config.GetPort())
}
