package sntn

type Sentence interface {
	Verb() string
	Object() string
	Punctuation() string
	String() string
}
