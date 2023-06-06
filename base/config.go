package base


/*

{
	"protocol":"tcp",
	"ip":"127.0.0.1",
	"port":80
}

*/
type NetConfig struct {
	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     int32    `json:"port"`
}

func NewNetConfig(Protocol string,Ip string,Port int32) *NetConfig {
	return &NetConfig{
		Protocol: Protocol,
		Ip:       Ip,
		Port:     Port,
	}
}

// todo yaml
func NewNetConfigWithConfig(path string) *NetConfig {
	return &NetConfig{
	}
}

func (n *NetConfig) GetProtocol() string {
	return 	n.Protocol
}

func (n *NetConfig) GetIp() string {
	return 	n.Ip
}

func (n *NetConfig) GetPort() int32 {
	return 	n.Port
}



