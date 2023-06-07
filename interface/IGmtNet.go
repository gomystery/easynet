package _interface

import "net"

type IGmtNet interface {

	OnStart(conn net.Conn) error


	OnConnect(conn net.Conn) error


	OnReceive(conn net.Conn,bytes []byte) error


	OnShutdown(conn net.Conn) error


	OnClose(conn net.Conn, err error) error


	// todo to add more
}