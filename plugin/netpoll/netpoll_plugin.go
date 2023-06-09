package netpoll

import (
	"context"
	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"
	"net"
)

type NetPollGmtNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	Server *NetPollServer

	Handler _interface.IGmtNet
}

func NewNetPollGmtNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IGmtNet) *NetPollGmtNetPlugin {
	GmtNetPlugin := &NetPollGmtNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	Server := NewNetPollServer(ctx, config, handler)
	GmtNetPlugin.Server = Server

	return GmtNetPlugin
}

func (g NetPollGmtNetPlugin) Run() error {
	return g.Server.Run()
}
