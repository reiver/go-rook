package impr

import (
	"strings"
)

type Imperation struct {
	verb string
	object string
	loaded bool
}

func Nothing() Imperation {
	return Imperation{}
}

func Something(verb string, object string) Imperation {
	return Imperation{
		loaded: true,
		verb:verb,
		object:object,
	}
}

func (receiver Imperation) Object() string {
	if Nothing() == receiver {
		return ""
	}

	return receiver.object
}

func (receiver Imperation) Punctuation() string {
	return "/"
}

func (receiver Imperation) String() string {
	if Nothing() == receiver {
		return ""
	}

	var bldr strings.Builder

	if "" != receiver.verb {
		bldr.WriteString(receiver.Punctuation())
		bldr.WriteString(receiver.verb)
	}

	if "" != receiver.object {
		if 0 < bldr.Len() {
			bldr.WriteRune(' ')
		}

		bldr.WriteString(receiver.object)
	}

	bldr.WriteRune('\u0085')

	return bldr.String()
}

func (receiver Imperation) Verb() string {
	if Nothing() == receiver {
		return ""
	}

	return receiver.verb
}
