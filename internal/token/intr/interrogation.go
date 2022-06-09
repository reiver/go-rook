package intr

import (
	"strings"
)

type Interrogation struct {
	verb string
	object string
	loaded bool
}

func Nothing() Interrogation {
	return Interrogation{}
}

func Something(verb string, object string) Interrogation {
	return Interrogation{
		loaded: true,
		verb:verb,
		object:object,
	}
}

func (receiver Interrogation) Object() string {
	if Nothing() == receiver {
		return ""
	}

	return receiver.object
}

func (receiver Interrogation) Punctuation() string {
	return "?"
}

func (receiver Interrogation) String() string {
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

	bldr.WriteString("\r\n")

	return bldr.String()
}

func (receiver Interrogation) Verb() string {
	if Nothing() == receiver {
		return ""
	}

	return receiver.object
}
