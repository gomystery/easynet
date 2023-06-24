package gnet

import (
	"context"
	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
	"github.com/panjf2000/gnet/v2"
	"net"
)

type GnetEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	Server *GnetServer

	Handler _interface.IEasyNet
}

func NewGnetEasyNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IEasyNet) *GnetEasyNetPlugin {
	gnetEasyNetPlugin := &GnetEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	gnetServer := NewGnetServer(ctx, config, handler)
	gnetEasyNetPlugin.Server = gnetServer

	return gnetEasyNetPlugin
}

func (g GnetEasyNetPlugin) Run() error {
	err := gnet.Run(
		g.Server,
		g.Server.addr,
		gnet.WithMulticore(g.Server.multicore))
	return err
}
