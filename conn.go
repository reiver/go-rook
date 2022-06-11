package rook

import (
	"github.com/reiver/go-rook/internal/sntn"

	"io"
	"net"
)

// Conn represents a connection.
type Conn struct {
	conn Connor
}

// Close calls .Close() on the underlying Connor (which has its own .Close() too).
//
// Close makes rook.Conn match the io.Closer interface.
func (receiver *Conn) Close() error {
	if nil == receiver {
		return nil
	}

	var conn Connor
	{
		conn = receiver.conn

		if nil == conn {
			return errInternalError
		}
	}

	return conn.Close()
}

// DLEWriter return an io.Writer that deals with DLE byte-stuffing.
// With DLE byte-stuffing it escapes any bytes with a value of of DLE (0x10) or SYN (0x16).
//
// So, for example, if the original bytes is:
//
//      //                                               DLE                   SYN
//      []byte{0x00,0x02,0x04,0x08,0x1a,0x0c,0x0e,       0x10,0x12,0x14,       0x16,0x18,0x1a,0x1c,0x1e,0x20}
//
// Then DLE byte-stuffed version would be.
//
//      //                                         DLE   DLE             DLE   SYN
//      []byte{0x00,0x02,0x04,0x08,0x1a,0x0c,0x0e, 0x10, 0x10,0x12,0x14, 0x10, 0x16,0x18,0x1a,0x1c,0x1e,0x20}
func (receiver *Conn) DLEWriter() (io.Writer, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	var conn Connor
	{
		conn = receiver.conn

		if nil == conn {
			return nil, errInternalError
		}
	}

	return dleWriter{conn}, nil
}

// LocalAddr returns the local address.
//
// LocalAddr is common in the "net" package, and including it here is meant to make this similar to its usage there.
func (receiver *Conn) LocalAddr() net.Addr {
	if nil == receiver {
		return nil
	}

	var conn Connor
	{
		conn = receiver.conn

		if nil == conn {
			return nil
		}
	}

	return conn.LocalAddr()
}

// RemoteAddr returns the remote address.
//
// RemoteAddr is common in the "net" package, and including it here is meant to make this similar to its usage there.
func (receiver *Conn) RemoteAddr() net.Addr {
	if nil == receiver {
		return nil
	}

	var conn Connor
	{
		conn = receiver.conn

		if nil == conn {
			return nil
		}
	}

	return conn.RemoteAddr()
}
