package gmtnet

import (
	"context"
	"fmt"
	"net"

	"github.com/gomystery/gmtnet/base"
	"github.com/gomystery/gmtnet/interface"
	"github.com/gomystery/gmtnet/plugin/gnet"
)


type GmtNet struct {

	handler _interface.IGmtNet

	Conn net.Conn

	Ctx context.Context

	GmtNetPlugin _interface.IPlugin

	Config *base.NetConfig

}


func NewGmtNet(ctx context.Context,netName string,config *base.NetConfig,handler _interface.IGmtNet) *GmtNet {
	gmtnet := &GmtNet{
		Ctx:     ctx,
		handler: handler,
	}

	// todo new GmtNetPlugin
	switch netName {
	case "Gnet":
		gmtnet.GmtNetPlugin = gnet.NewGnetGmtNetPlugin(ctx,config,handler)
		err := gmtnet.GmtNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	return gmtnet
}




