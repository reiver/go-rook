package impr

import (
	"github.com/reiver/go-rook/internal/verbs"
)

// Pull returns a /PULL imperation token.
func Pull(value string) Imperation {
	return Something(verbs.PULL, value)
}
