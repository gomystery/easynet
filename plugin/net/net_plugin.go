package net

import (
	"context"
	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
	"net"
)

type NetEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	Server *NetServer

	Handler _interface.IEasyNet
}

func NewNetEasyNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IEasyNet) *NetEasyNetPlugin {
	easyNetPlugin := &NetEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	Server := NewNetPollServer(ctx, config, handler)
	easyNetPlugin.Server = Server

	return easyNetPlugin
}

func (g NetEasyNetPlugin) Run() error {
	return g.Server.Run()
}
