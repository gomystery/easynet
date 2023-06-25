package easynet

import (
	"context"
	"fmt"
	"github.com/gomystery/easynet/base"
	"testing"
)

type Handler struct {
}

func (h Handler) OnStart(conn interface{}) error {
	return nil
}

func (h Handler) OnConnect(conn interface{}) error {
	return nil

}

func (h Handler) OnReceive(conn interface{}, bytes []byte) ([]byte, error) {
	return nil, nil

}

func (h Handler) OnShutdown(conn interface{}) error {
	return nil
}

func (h Handler) OnClose(conn interface{}, err error) error {
	return nil
}

func TestGnet(t *testing.T) {
	config := base.NewNetConfig("tcp", "127.0.0.1", 9011)
	handler := &Handler{}
	gmet := NewEasyNet(context.Background(), "NetPoll", config, handler)
	fmt.Println(gmet)
}
