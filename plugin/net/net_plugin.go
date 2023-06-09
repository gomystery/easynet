package net

import (
	"context"
	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"
	"net"
)

type NetGmtNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	Server *NetServer

	Handler _interface.IGmtNet
}

func NewNetGmtNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IGmtNet) *NetGmtNetPlugin {
	GmtNetPlugin := &NetGmtNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	Server := NewNetPollServer(ctx, config, handler)
	GmtNetPlugin.Server = Server

	return GmtNetPlugin
}

func (g NetGmtNetPlugin) Run() error {
	return g.Server.Run()
}
