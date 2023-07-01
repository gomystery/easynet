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

	Config *base.DeFaultNetConfig
}

func NewEasyNet(ctx context.Context, netName string, config _interface.IConfig, handler _interface.IEasyNet) *EasyNet {
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
		easynet.EasyNetPlugin = evio.NewEvioEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("no expected net plugin")
	}

	return easynet
}



func NewEasyNetWithYamlConfig(ctx context.Context, netName string, handler _interface.IEasyNet,path string) *EasyNet {
	easynet := &EasyNet{
		Ctx:     ctx,
		handler: handler,
	}

	// todo new EasyNetPlugin
	switch netName {
	case "Gnet":
		config := base.NewNetConfigWithConfig(path,netName)
		easynet.EasyNetPlugin = gnet.NewGnetEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Gev":
		config := base.NewNetConfigWithConfig(path,netName)
		easynet.EasyNetPlugin = gev.NewGevEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Net":
		config := base.NewNetConfigWithConfig(path,netName)
		easynet.EasyNetPlugin = np.NewNetEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "NetPoll":
		config := base.NewNetConfigWithConfig(path,netName)
		easynet.EasyNetPlugin = netpoll.NewNetPollEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Evio":
		config := base.NewNetConfigWithConfig(path,netName)
		easynet.EasyNetPlugin = evio.NewEvioEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("no expected net plugin")
	}

	return easynet
}