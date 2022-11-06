package rook_test

import (
	"github.com/reiver/go-rook"

	"testing"
)

func TestSentence_String(t *testing.T) {

	tests := []struct{
		Value rook.Sentence
		Expected string
	}{
		{
			Value: rook.Sentence{
				Punctuation: "",
				Verb:        "",
				Object:      "",
			},
			Expected: "\r\n",
		},



		{
			Value: rook.Sentence{
				Punctuation: "!",
				Verb:        "DOROOD",
				Object:      "",
			},
			Expected: "!DOROOD\u0085",
		},



		{
			Value: rook.Sentence{
				Punctuation: "!",
				Verb:        "IDLING",
				Object:      "",
			},
			Expected: "!IDLING\u0085",
		},



		{
			Value: rook.Sentence{
				Punctuation: "/",
				Verb:        "PULL",
				Object:      "apple",
			},
			Expected: "/PULL apple\u0085",
		},
		{
			Value: rook.Sentence{
				Punctuation: "/",
				Verb:        "PULL",
				Object:      "banana",
			},
			Expected: "/PULL banana\u0085",
		},
		{
			Value: rook.Sentence{
				Punctuation: "/",
				Verb:        "PULL",
				Object:      "cherry",
			},
			Expected: "/PULL cherry\u0085",
		},



		{
			Value: rook.Sentence{
				Punctuation: "/",
				Verb:        "W",
				Object:      "",
			},
			Expected: "/W\r\n",
		},
		{
			Value: rook.Sentence{
				Punctuation: "/",
				Verb:        "W",
				Object:      "joeblow",
			},
			Expected: "/W joeblow\r\n",
		},



		{
			Value: rook.Sentence{
				Punctuation: "",
				Verb:        "",
				Object:      "joeblow",
			},
			Expected: "joeblow\r\n",
		},



		{
			Value: rook.Sentence{
				Punctuation: "/",
				Verb:        "GET",
				Object:      "https://www.example.com/apple/banana/cherry.html?one=1&two=2",
			},
			Expected: "/GET https://www.example.com/apple/banana/cherry.html?one=1&two=2\u0085",
		},
		{
			Value: rook.Sentence{
				Punctuation: "",
				Verb:        "GET",
				Object:      "https://www.example.com/apple/banana/cherry.html?one=1&two=2",
			},
			Expected:
				"GET /apple/banana/cherry.html?one=1&two=2 HTTP/1.1" +"\r\n"+
				"Host: www.example.com"                              +"\r\n"+
				"Connection: close"                                  +"\r\n"+
				                                                      "\r\n",
		},
	}

	for testNumber, test := range tests {

		var actual   string = test.Value.String()
		var expected string = test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PUNCTUATION: %q", test.Value.Punctuation)
			t.Logf("VERB:        %q", test.Value.Verb)
			t.Logf("OBJECT:      %q", test.Value.Object)
			continue
		}
	}
}
