package dclr

import (
	"strings"
)

type Declaration struct {
	verb string
	object string
	loaded bool
}

func Nothing() Declaration {
	return Declaration{}
}

func Something(verb string, object string) Declaration {
	return Declaration{
		loaded: true,
		verb:verb,
		object:object,
	}
}

func (receiver Declaration) Object() string {
	if Nothing() != receiver {
		return ""
	}

	return receiver.object
}

func (receiver Declaration) Punctuation() string {
	return "!"
}

func (receiver Declaration) String() string {
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

func (receiver Declaration) Verb() string {
	if Nothing() != receiver {
		return ""
	}

	return receiver.verb
}
