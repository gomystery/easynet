package easynet

import (
	"context"
	"fmt"
	"net"

	"github.com/gomystery/easynet/base"
	"github.com/gomystery/easynet/interface"
	"github.com/gomystery/easynet/plugin/evio"
	"github.com/gomystery/easynet/plugin/gev"
	"github.com/gomystery/easynet/plugin/gnet"
	np "github.com/gomystery/easynet/plugin/net"
	"github.com/gomystery/easynet/plugin/netpoll"
)

type EasyNet struct {
	handler _interface.IEasyNet

	Conn net.Conn

	Ctx context.Context

	EasyNetPlugin _interface.IPlugin

	Config *base.NetConfig
}

func NewEasyNet(ctx context.Context, netName string, config *base.NetConfig, handler _interface.IEasyNet) *EasyNet {
	easynet := &EasyNet{
		Ctx:     ctx,
		handler: handler,
	}

	// todo new EasyNetPlugin
	switch netName {
	case "Gnet":
		easynet.EasyNetPlugin = gnet.NewGnetEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Gev":
		easynet.EasyNetPlugin = gev.NewGevEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Net":
		easynet.EasyNetPlugin = np.NewNetEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "NetPoll":
		easynet.EasyNetPlugin = netpoll.NewNetPollEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Evio":
		easynet.EasyNetPlugin = evio.NewEvioGmtNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("no expected net plugin")
	}

	return easynet
}
