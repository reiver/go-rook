package sntn

import (
	"github.com/reiver/go-fck"
)

const (
	errInternalError = fck.Error("internal error")
	errNilReader     = fck.Error("nil reader")
	errReadTooShort  = fck.Error("read too short")
	errRuneError     = fck.Error("rune error")
)
