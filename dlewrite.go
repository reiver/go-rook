package rook

import (
	"bytes"
	"io"
)

const (
	DLE = byte(0x10)
	SYN = byte(0x16)
)

// DLEWrite writers ‘data’ to ‘w’, but does byte-stuffing to escape any bytes with a value DLE (0x10) or SYN (0x16).
func DLEWrite(w io.Writer, data []byte) (int, error) {
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
			case DLE, SYN:
				buffer.WriteByte(DLE)
			}
			buffer.WriteByte(b)
			n++
		}
	}

	{
		w.Write(buffer.Bytes())
	}

	return n, nil
}
