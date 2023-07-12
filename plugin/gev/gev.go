package gev

import (
	"context"
	"fmt"
	"github.com/Allenxuxu/gev"
	"github.com/baickl/logger"
	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
)

type GevServer struct {
	Ctx       context.Context
	addr      string

	handler _interface.IEasyNet
	InputStreamMap map[string]_interface.IInputStream

}

func NewGevServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *GevServer {
	return &GevServer{
		Ctx:       ctx,
		addr:      fmt.Sprintf("%s://%s:%d", config.GetProtocol(), config.GetIp(), config.GetPort()),
		handler:   handler,
		InputStreamMap: make(map[string]_interface.IInputStream),
	}
}

func (s *GevServer) OnConnect(c *gev.Connection) {
	logger.Infoln("Gev server OnConnect ")
	s.InputStreamMap[c.PeerAddr()] = &base.InputStream{}
	err := s.handler.OnConnect(c)
	if err != nil {
		logger.Errorf("Gnet OnConnect err %v\n", err)
	}
	return
}

func (s *GevServer) OnMessage(c *gev.Connection, ctx interface{}, data []byte) (out interface{}) {
	s.InputStreamMap[c.PeerAddr()].Begin(data)
	data, err := s.handler.OnReceive(c, s.InputStreamMap[c.PeerAddr()])
	if err != nil {
		logger.Errorf("Gnet OnMessage err %v\n", err)
	}
	return data
}

func (s *GevServer) OnClose(c *gev.Connection) {
	s.InputStreamMap[c.PeerAddr()] = nil
	err := s.handler.OnClose(c, nil)
	if err != nil {
		logger.Errorf("Gnet OnClose err %v\n", err)
	}
	return
}
