package rook

type Connor interface {
	Close() error
	LocalAddr() net.Addr
	Read(b []byte) (n int, err error)
	RemoteAddr() net.Addr
	Write(b []byte) (n int, err error)
}
