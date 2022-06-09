package verbs

// Imperation
//
// The actual protocol has imperations start with a '/' character.
// For example:
//
//	/PULL
//
// But these constants do not include the '/' character at the beginning.
const (
	PULL = "PULL"
	W    = "W"
)

// Declarations
//
// The actual protocol has declarations start with a '!' character.
// For example:
//
//	!PUSHING
//
// But these constants do not include the '!' character at the beginning.
const (
	DOROOD  = "DOROOD"
	IDLING  = "IDLING"
	PUSHING = "PUSHING"
)
