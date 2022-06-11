package rook

import (
	"bytes"
	"io"
)

const (
	dle = byte(0x10)
	syn = byte(0x16)
)

// dleWriter deals with DLE byte-stuffing, that escapes any bytes with a value of of DLE (0x10) or SYN (0x16).
type dleWriter struct {
	writer io.Writer
}

var _ io.Writer = &dleWriter{}

// Write writers ‘data’ to ‘w’, but does DLE byte-stuffing that escapes any bytes with a value of DLE (0x10) or SYN (0x16).
//
// For example, if ‘data’ is:
//
//	//                                               DLE                   SYN
//	[]byte{0x00,0x02,0x04,0x08,0x1a,0x0c,0x0e,       0x10,0x12,0x14,       0x16,0x18,0x1a,0x1c,0x1e,0x20}
//
// Then what would be written to ‘w’ would be:
//
//	//                                         DLE   DLE             DLE   SYN
//	[]byte{0x00,0x02,0x04,0x08,0x1a,0x0c,0x0e, 0x10, 0x10,0x12,0x14, 0x10, 0x16,0x18,0x1a,0x1c,0x1e,0x20}
//
func (receiver dleWriter) Write(data []byte) (int, error) {

	var w io.Writer = receiver.writer
	if nil == w {
		return 0, errNilWriter
	}

	if 0 >= len(data) {
		return 0, nil
	}

	var n int

	var buffer bytes.Buffer
	{
		for _, b := range data {
			switch b {
			case dle, syn:
				buffer.WriteByte(dle)
			}
			buffer.WriteByte(b)
			n++
		}
	}

	{
		n2, err2 := w.Write(buffer.Bytes())
		if nil != err2 {
//@TODO: Shouldn't be using n2 here. Shouldn't could the byte-stuffing.
			return n2, err2
		}
	}

	return n, nil
}
