package netpoll

import (
	"context"
	"github.com/gomystery/easynet/interface"
	"log"
	"net"
)

type NetPollEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *YamlConfig

	Server *NetPollServer

	Handler _interface.IEasyNet
}

func NewNetPollEasyNetPlugin(ctx context.Context, iconfig _interface.IConfig, handler _interface.IEasyNet) *NetPollEasyNetPlugin {

	var config *YamlConfig
	var ok bool
	if config,ok=iconfig.(*YamlConfig);!ok{
		log.Printf("netpoll yaml error \n")
	}

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
