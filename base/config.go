package base

import (
	"fmt"
	"github.com/gomystery/easynet/interface"
	"github.com/gomystery/easynet/plugin/evio"
	"github.com/gomystery/easynet/plugin/gev"
	"github.com/gomystery/easynet/plugin/gnet"
	"github.com/gomystery/easynet/plugin/net"
	"github.com/gomystery/easynet/plugin/netpoll"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/*
	{
		"protocol":"tcp",
		"ip":"127.0.0.1",
		"port":80
	}
*/
type DeFaultNetConfig struct {
	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     int32  `json:"port"`
}

func NewDefaultNetConfig(Protocol string, Ip string, Port int32) _interface.IConfig  {
	return &DeFaultNetConfig{
		Protocol: Protocol,
		Ip:       Ip,
		Port:     Port,
	}
}

// todo yaml
func NewNetConfigWithConfig(path string,netName string) _interface.IConfig {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	//var _config *config.Config
	var config _interface.IConfig
	switch netName {
	case "Gnet":
		config = &gnet.YamlConfig{}
	case "Gev":
		config = &gev.YamlConfig{}
	case "Net":
		config = &net.YamlConfig{}
	case "NetPoll":
		config = &netpoll.YamlConfig{}
	case "Evio":
		config = &evio.YamlConfig{}

	default:
		fmt.Println("no expected net name")
	}
	//将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}

func (n *DeFaultNetConfig) GetProtocol() string {
	return n.Protocol
}

func (n *DeFaultNetConfig) GetIp() string {
	return n.Ip
}

func (n *DeFaultNetConfig) GetPort() int32 {
	return n.Port
}
