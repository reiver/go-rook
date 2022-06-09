package rook

import (
	"io"
	"net"
)

// A Connor represents a lower-level network connection.
//
// It is an io.ReadWriteCloser with the usual .Close(), .Read(), and .Write() methods.
// But also has the .LocalAddr() and .RemoteAddr() methods that are common with Go network programming.
type Connor interface {
	io.ReadWriteCloser
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}
