package dclr

import (
	"github.com/reiver/go-rook/internal/verbs"
)

// Pushing returns a !PUSHING declaration token.
func Pushing(value string) Declaration {
	return Something(verbs.PUSHING, value)
}
