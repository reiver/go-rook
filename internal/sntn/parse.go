package sntn

import (
	"github.com/reiver/go-rook/internal/token/dclr"
	"github.com/reiver/go-rook/internal/token/impr"
	"github.com/reiver/go-rook/internal/token/intr"
	"github.com/reiver/go-rook/internal/verbs"

	"github.com/reiver/go-utf8s"

	"io"
	"strings"
)

// Parse parses a sentence into tokens from an io.Reader.
func Parse(reader io.Reader) (Sentence, error) {
	if nil == reader {
		return nil, errNilReader
	}

	var r rune
	var n int
	var err error

	// Try to read the first rune.
	//
	// There should be at least one rune.
	// (Because even if it is an empty line, there will still be a line-terminator character.)
	// If there isn't it is an error.
	{
		r, n, err = utf8s.ReadRune(reader)
		if nil != err {
/////////////////////// RETURN
			return nil, err
		}
		if 0 >= n {
/////////////////////// RETURN
			return nil, errReadTooShort
		}
		if utf8s.RuneError == r {
/////////////////////// RETURN
			return nil, errRuneError
		}
	}

	// Check to see if the sentence line is empty (other than the line-terminator character).
	//
	// The shortest possible (valid) sentence line could be one of the following:
	//
	//	"\u000A" // line feed
	//
	//	"\u000B" // vertical tab
	//
	//	"\u000C" // form feed
	//
	//	"\u000D" // carriage return
	//
	//	"\u0085" // next line
	//
	//	"\u2028" // line separator
	//
	//	"\u2029" // paragraph separator
	//
	// If we had the following empty sentence line:
	//
	//	"\u000A" // line feed
	//
	// Then we treat this the same as:
	//
	//	"!IDLING\u0085"
	//
	// If we have any of the following empty sentence lines:
	//
	//	"\u000B" // vertical tab
	//
	//	"\u000C" // form feed
	//
	//	"\u000D" // carriage return
	//
	//	"\u0085" // next line
	//
	//	"\u2028" // line separator
	//
	//	"\u2029" // paragraph separator
	//
	// Then we treat this the same as:
	//
	//	"/DOROOD\u0085"
	//
	// The reason we make "\u000A" different is so we can simplify the parsing of lines that end in "\r\n".
	// I want this protocol to be something a software developer earlier in their learning-experience to be able to implement from scratch,
	// and hopefully this way of doing it makes it easier.
	{
		switch r {
		case '\u000A': // line feed
/////////////////////// RETURN
			return dclr.Idling(), nil

		case '\u000B', // vertical tab
		     '\u000C', // form feed
		     '\u000D', // carriage return
		     '\u0085', // next line
		     '\u2028', // line separator
		     '\u2029': // paragraph separator
/////////////////////// RETURN
			return dclr.Dorood(), nil
		}
	}

	var verb   strings.Builder
	var object strings.Builder

	// Figure out whether this is a declaration, an imperation, or an interrogation.
	//
	// We use the first character of the line to figure this out.
	// (Which should hopefully make this easier to parse for software developers earlier in their learning-experience.)
	//
	// '!' ≡  declaration
	//
	// '/' ≡  imperation (i.e., a command)
	//
	// '?' ≡  interrogation (i.e., a question, or a query)
	//
	// If the line does not begin with on of these, then we default this to the imperation /PULL
	var punctuation rune
	{
		switch r {
		case '!','/','?':
			punctuation = r
		default:
			punctuation = '/'
			verb.WriteString(verbs.PULL)
			object.WriteRune(r)
		}
	}

	// Figure out the verb.
	if 0 >= verb.Len() {
		loop1: for {
			r, n, err = utf8s.ReadRune(reader)
			if nil != err {
/////////////////////////////// RETURN
				return nil, err
			}
			if 0 >= n {
/////////////////////////////// RETURN
				return nil, errReadTooShort
			}
			if utf8s.RuneError == r {
/////////////////////////////// RETURN
				return nil, errRuneError
			}

			switch r {
			case '\u0020': // space
	/////////////////////// BREAK
				break loop1

			case '\u000A', // line feed
			     '\u000B', // vertical tab
			     '\u000C', // form feed
			     '\u000D', // carriage return
			     '\u0085', // next line
			     '\u2028', // line separator
			     '\u2029': // paragraph separator

				switch punctuation {
				case '!':
					return dclr.Something(verb.String(), ""), nil
				case '/':
					return impr.Something(verb.String(), ""), nil
				case '?':
					return intr.Something(verb.String(), ""), nil
				default:
					return nil, errInternalError
				}
			default:
				verb.WriteRune(r)
			}
		}

	}

	{
		loop2: for {
			r, n, err = utf8s.ReadRune(reader)
			if nil != err {
/////////////////////////////// RETURN
				return nil, err
			}
			if 0 >= n {
/////////////////////////////// RETURN
				return nil, errReadTooShort
			}
			if utf8s.RuneError == r {
/////////////////////////////// RETURN
				return nil, errRuneError
			}

			switch r {
			case '\u000A', // line feed
			     '\u000B', // vertical tab
			     '\u000C', // form feed
			     '\u000D', // carriage return
			     '\u0085', // next line
			     '\u2028', // line separator
			     '\u2029': // paragraph separator
	/////////////////////// BREAK
				break loop2

			default:
				object.WriteRune(r)
			}
		}

	}

	{
		switch punctuation {
		case '!':
			return dclr.Something(verb.String(), object.String()), nil
		case '/':
			return impr.Something(verb.String(), object.String()), nil
		case '?':
			return intr.Something(verb.String(), object.String()), nil
		default:
			return nil, errInternalError
		}
	}
}
