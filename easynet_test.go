package easynet

import (
	"context"
	"fmt"
	_interface "github.com/gomystery/easynet/interface"
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

func (h Handler) OnReceive(conn interface{}, stream _interface.IInputStream) ([]byte, error) {
	return nil, nil

}

func (h Handler) OnShutdown(conn interface{}) error {
	return nil
}

func (h Handler) OnClose(conn interface{}, err error) error {
	return nil
}

func TestEasyNet(t *testing.T) {
	config := NewDefaultNetConfig("tcp", "127.0.0.1", 9011)
	handler := &Handler{}
	gmet := NewEasyNet(context.Background(), "NetPoll", config, handler)
	fmt.Println(gmet)
}

func TestEasyNetWithYamlConfig(t *testing.T) {
	handler := &Handler{}
	gmet := NewEasyNetWithYamlConfig(context.Background(), "Evio", handler, "./base/confg.yaml")
	fmt.Println(gmet)
}
