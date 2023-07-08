package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/netpoll"

	"github.com/gomystery/easynet"
)

type Handler struct {
}

func (h Handler) OnStart(conn interface{}) error {
	fmt.Println("test OnStart")
	return nil
}

func (h Handler) OnConnect(conn interface{}) error {
	netpollConn, ok:= conn.(netpoll.Connection)
	if !ok {
		fmt.Println("test conn err")
	}

	fmt.Println("test conn LocalAddr",netpollConn.LocalAddr())
	fmt.Println("test conn RemoteAddr",netpollConn.RemoteAddr())

	return nil

}

func (h Handler) OnReceive(conn interface{}, bytes []byte) ([]byte, error) {
	//netpollConn, ok:= conn.(netpoll.Connection)
	//if !ok {
	//	fmt.Println("test conn err")
	//}

	fmt.Println("test receive msg ",string(bytes))

	return bytes, nil

}

func (h Handler) OnShutdown(conn interface{}) error {
	return nil
}

func (h Handler) OnClose(conn interface{}, err error) error {
	return nil
}

func main() {
	handler := &Handler{}
	easynet.NewEasyNetWithYamlConfig(context.Background(), "NetPoll", handler, "../../base/confg.yaml")


}
