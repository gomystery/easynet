package gnet

import (
	"context"
	"github.com/gomystery/easynet/interface"
	"github.com/panjf2000/gnet/v2"
	"log"
	"net"
)

type GnetEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *YamlConfig

	Server *GnetServer

	Handler _interface.IEasyNet
}

func NewGnetEasyNetPlugin(ctx context.Context, iconfig _interface.IConfig, handler _interface.IEasyNet) *GnetEasyNetPlugin {

	var config *YamlConfig
	var ok bool
	if config,ok=iconfig.(*YamlConfig);!ok{
		log.Printf("gnet yaml error \n")
	}

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
