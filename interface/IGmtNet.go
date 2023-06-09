package _interface

type IGmtNet interface {
	OnStart(conn interface{}) error

	OnConnect(conn interface{}) error

	OnReceive(conn interface{}, bytes []byte) ([]byte,error)

	OnShutdown(conn interface{}) error

	OnClose(conn interface{}, err error) error

	// todo to add more
}
