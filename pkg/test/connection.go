package test

import (
	"net"
	"time"
)

type TestAddr struct {
	addr string
}

func (a *TestAddr) Network() string {
	return a.addr
}

func (a *TestAddr) String() string {
	return a.addr
}

type TestConn struct {
	addr TestAddr
}

func (c *TestConn) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (c *TestConn) Write(b []byte) (n int, err error) {
	return 0, nil
}

func (c *TestConn) Close() error {
	return nil
}

func (c *TestConn) LocalAddr() net.Addr {
	return &c.addr
}

func (c *TestConn) RemoteAddr() net.Addr {
	return &c.addr
}

func (c *TestConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *TestConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *TestConn) SetWriteDeadline(t time.Time) error {
	return nil
}
