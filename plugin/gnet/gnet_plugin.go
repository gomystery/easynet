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

	GnetServer *GnetServer

	Handler _interface.IEasyNet
}

func NewGnetEasyNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IEasyNet) *GnetEasyNetPlugin {
	gnetEasyNetPlugin := &GnetEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	gnetServer := NewGnetServer(ctx, config, handler)
	gnetEasyNetPlugin.GnetServer = gnetServer

	return gnetEasyNetPlugin
}

func (g GnetEasyNetPlugin) Run() error {
	err := gnet.Run(
		g.GnetServer,
		g.GnetServer.addr,
		gnet.WithMulticore(g.GnetServer.multicore))
	return err
}
