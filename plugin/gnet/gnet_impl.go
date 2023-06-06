package gnet

import (
	"context"
	"net"

	"github.com/gomystery/gmtnet/base"
)

type GnetGmtNetImpl struct {
	Conn net.Conn

	Ctx context.Context

	Config *base.NetConfig

	GnetServer *GnetServer

	Gmtnet base.IGmtNet
}


func NewGnetGmtNetImpl(ctx context.Context,config *base.NetConfig,gmtnet base.IGmtNet) *GnetGmtNetImpl {

	gnetGmtNetImpl := &GnetGmtNetImpl{
		Ctx:    ctx,
		Config: config,
		Gmtnet: gmtnet,
	}

	gnetServer := NewGnetServer(ctx,config,gnetGmtNetImpl)
	gnetGmtNetImpl.GnetServer = gnetServer

	return gnetGmtNetImpl
}

func (g GnetGmtNetImpl) OnStart(conn net.Conn) error {
	return g.Gmtnet.OnStart(conn)
}

func (g GnetGmtNetImpl) OnConnect(conn net.Conn) error {
	return g.Gmtnet.OnConnect(conn)
}

func (g GnetGmtNetImpl) OnReceive(conn net.Conn, bytes []byte) error {
	return g.Gmtnet.OnReceive(conn,bytes)
}

func (g GnetGmtNetImpl) OnShutdown(conn net.Conn) error {
	return g.Gmtnet.OnShutdown(conn)
}

func (g GnetGmtNetImpl) OnClose(conn net.Conn, err error) error {
	return g.Gmtnet.OnClose(conn,err)
}

