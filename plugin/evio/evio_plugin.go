package evio

import (
	"context"
	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"
	"net"
)

type EvioGmtNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	server *EvioServer

	Handler _interface.IGmtNet
}

func NewEvioGmtNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IGmtNet) *EvioGmtNetPlugin {
	evioGmtNetPlugin := &EvioGmtNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	server := NewGnetServer(ctx, config, handler)
	evioGmtNetPlugin.server = server

	return evioGmtNetPlugin
}

func (g EvioGmtNetPlugin) Run() error {
	return g.server.Run()
}
