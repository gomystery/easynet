package gev

import (
	"context"
	"github.com/Allenxuxu/gev"
	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"
	"net"
	"strconv"
	"time"
)

type GevGmtNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	Server *GevServer

	Handler _interface.IGmtNet
}

func NewGevGmtNetPlugin(ctx context.Context, config *base.NetConfig, handler _interface.IGmtNet) *GevGmtNetPlugin {
	GmtNetPlugin := &GevGmtNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	Server := NewGevServer(ctx, config, handler)
	GmtNetPlugin.Server = Server

	return GmtNetPlugin
}

func (g GevGmtNetPlugin) Run() error {
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
