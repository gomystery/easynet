package gmtnet

import (
	"context"
	"fmt"
	"github.com/gomystery/gmtnet/base"
	"net"
	"testing"
)

type Handler struct {

}

func (h Handler) OnStart(conn net.Conn) error {
	return nil
}

func (h Handler) OnConnect(conn net.Conn) error {
	return nil

}

func (h Handler) OnReceive(conn net.Conn, bytes []byte) error {
	return nil

}

func (h Handler) OnShutdown(conn net.Conn) error {
	return nil

}

func (h Handler) OnClose(conn net.Conn, err error) error {
	return nil

}

func TestGnet(t *testing.T) {
	config := base.NewNetConfig("tcp","127.0.0.1",9009)
	handler := &Handler{}
	gmet := NewGmtNet(context.Background(),"Gnet",config,handler)
	fmt.Println(gmet)
}