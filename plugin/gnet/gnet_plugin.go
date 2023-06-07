package gnet

import (
	"context"
	"github.com/gomystery/gmtnet/base"
	"github.com/panjf2000/gnet/v2"
	"net"
)

type GnetGmtNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	GnetServer *GnetServer

	Handler base.IGmtNet
}


func NewGnetGmtNetPlugin(ctx context.Context,config *base.NetConfig, handler base.IGmtNet) *GnetGmtNetPlugin {
	gnetGmtNetPlugin := &GnetGmtNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	gnetServer := NewGnetServer(ctx,config, handler)
	gnetGmtNetPlugin.GnetServer = gnetServer

	return gnetGmtNetPlugin
}

func (g GnetGmtNetPlugin) Run() error {
	err := gnet.Run(
		g.GnetServer,
		g.GnetServer.addr,
		gnet.WithMulticore(g.GnetServer.multicore))
	return err
}
