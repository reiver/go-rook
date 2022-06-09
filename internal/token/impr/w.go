package impr

import (
	"github.com/reiver/go-rook/internal/verbs"
)

// W returns a /W imperation token.
func W(value string) Imperation {
	return Something(verbs.W, value)
}
