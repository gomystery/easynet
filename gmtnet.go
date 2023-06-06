package gmtnet

import (
	"context"
	"net"

	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/plugin/gnet"
)


type GmtNet struct {

	handler base.IGmtNet

	Conn net.Conn

	Ctx context.Context

	GmtNetImpl base.IGmtNet

	Config *base.NetConfig

}


func NewGmtNet(ctx context.Context,netName string,config *base.NetConfig,handler base.IGmtNet) base.IGmtNet {
	gmtnet := &GmtNet{
		Ctx:     ctx,
		handler: handler,
	}

	// todo new GmtNetImpl
	switch netName {
	case "Gnet":
		gmtnet.GmtNetImpl = gnet.NewGnetGmtNetImpl(ctx,config,gmtnet)
	}

	return gmtnet
}

func (g GmtNet) OnStart(conn net.Conn) error {
	return g.handler.OnStart(conn)
}

func (g GmtNet) OnConnect( conn net.Conn) error {
	return g.handler.OnConnect(conn)
}

func (g GmtNet) OnReceive(conn net.Conn, bytes []byte) error {
	return g.handler.OnReceive(conn,bytes)
}

func (g GmtNet) OnShutdown(conn net.Conn) error {
	return g.handler.OnShutdown(conn)
}

func (g GmtNet) OnClose(conn net.Conn, err error) error {
	return g.handler.OnClose(conn,err)
}





