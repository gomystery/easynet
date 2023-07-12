package _interface

type IEasyNet interface {
	OnStart(conn interface{}) error

	OnConnect(conn interface{}) error

	OnReceive(conn interface{}, ip IInputStream) ([]byte, error)

	OnShutdown(conn interface{}) error

	OnClose(conn interface{}, err error) error

	// todo to add more
}
