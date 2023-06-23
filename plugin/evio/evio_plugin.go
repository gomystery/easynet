package evio

import (
	"context"
	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
	"net"
)

type EvioEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	Server *EvioServer

	Handler _interface.IEasyNet
}

func NewEvioGmtNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IEasyNet) *EvioEasyNetPlugin {
	evioEasyNetPlugin := &EvioEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	server := NewGnetServer(ctx, config, handler)
	evioEasyNetPlugin.Server = server

	return evioEasyNetPlugin
}

func (g EvioEasyNetPlugin) Run() error {
	return g.Server.Run()
}
