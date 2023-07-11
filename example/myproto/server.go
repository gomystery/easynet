package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/gomystery/easynet"
	_interface "github.com/gomystery/easynet/interface"
)

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int16
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

type Handler struct {
}

func (h Handler) OnStart(conn interface{}) error {
	fmt.Println("test OnStart")
	return nil
}

func (h Handler) OnConnect(conn interface{}) error {
	return nil

}

// |len|body|
// 0   7    n
func (h Handler) OnReceive(conn interface{}, stream _interface.IInputStream) ([]byte, error) {

	var rspMsg string
	var left []byte
	left= stream.Begin(nil)
	for  {
		if len(left) < 2  {
			break
		}
		strlenBytes := left[:2]
		strlen := BytesToInt(strlenBytes)
		if len(left) >= 2+strlen {
			bodyBytes := left[2:2+strlen]
			fmt.Println("test receive msg ",string(bodyBytes))
			msg := fmt.Sprintf("test receive msg %s \n",string(bodyBytes))
			rspMsg = rspMsg + msg
			left = left[2+strlen:]
		}else {
			break
		}
	}

	stream.End(left)
	return []byte(rspMsg), nil

}

func (h Handler) OnShutdown(conn interface{}) error {
	return nil
}

func (h Handler) OnClose(conn interface{}, err error) error {
	return nil
}

func main() {
	handler := &Handler{}
	easynet.NewEasyNetWithYamlConfig(context.Background(), "Net", handler, "/Users/lili/GolandProjects/go/src/easynet/base/confg.yaml")



}
