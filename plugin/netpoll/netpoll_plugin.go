package netpoll

import (
	"context"
	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
	"net"
)

type NetPollEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	Server *NetPollServer

	Handler _interface.IEasyNet
}

func NewNetPollEasyNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IEasyNet) *NetPollEasyNetPlugin {
	easyNetPlugin := &NetPollEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	Server := NewNetPollServer(ctx, config, handler)
	easyNetPlugin.Server = Server

	return easyNetPlugin
}

func (g NetPollEasyNetPlugin) Run() error {
	return g.Server.Run()
}
