package rook

import (
	"bytes"
	"fmt"
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

	var lenData int = len(data)

	if 0 >= lenData {
		return 0, nil
	}

	var bufferLen int
	var buffer bytes.Buffer
	{
		for _, b := range data {
			switch b {
			case dle, syn:
				buffer.WriteByte(dle)
				bufferLen++
			}
			buffer.WriteByte(b)
			bufferLen++
		}
	}

	{
		n, err := w.Write(buffer.Bytes())
		if nil != err {
//@TODO: We shouldn't be using ‘n’ here, since it would have also counted the byte-stuffed DLEs; which is wrong since its value should be from the perspective of how many bytes of ‘data’ were written.
			return n, err
		}
		{
			expected := bufferLen
			actual   := n

			if expected != actual {
//@TODO: We shouldn't be using ‘n’ here, since it would have also counted the byte-stuffed DLEs; which is wrong since its value should be from the perspective of how many bytes of ‘data’ were written.
				return n, fmt.Errorf("actual number of bytes written (%d) is not what was expected (%d)", actual, expected)
			}
		}
	}

	return lenData, nil
}
