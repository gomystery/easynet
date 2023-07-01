package gev

import (
	"context"
	"github.com/Allenxuxu/gev"
	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
	"log"
	"net"
	"strconv"
)

type GevEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *YamlConfig

	Server *GevServer

	Handler _interface.IEasyNet
}

func NewGevEasyNetPlugin(ctx context.Context, iconfig _interface.IConfig, handler _interface.IEasyNet) *GevEasyNetPlugin {

	var config *YamlConfig
	var ok bool
	if config,ok=iconfig.(*YamlConfig);!ok{
		log.Printf("gev yaml error \n")
	}

	easyNetPlugin := &GevEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	Server := NewGevServer(ctx, config, handler)
	easyNetPlugin.Server = Server

	return easyNetPlugin
}

func (g GevEasyNetPlugin) Run() error {
	s, err := gev.NewServer(g.Server,
		gev.Network(g.Config.Protocol),
		gev.Address(":"+(strconv.Itoa(int(g.Config.Port)))),
		gev.NumLoops(100),
	)
	if err != nil {
		return err
	}

	s.Start()
	return nil
}
