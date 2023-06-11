package gev

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/Allenxuxu/gev"
	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
)

type GevEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	Server *GevServer

	Handler _interface.IEasyNet
}

func NewGevEasyNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IEasyNet) *GevEasyNetPlugin {
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
		gev.Address(":"+ (strconv.Itoa( int(g.Config.Port)))  ),
		gev.NumLoops(100),
	)
	if err != nil {
		return err
	}

	s.RunEvery(time.Second*2, func() {

	})

	s.Start()
	return nil
}
