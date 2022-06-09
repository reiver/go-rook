package dclr

import (
	"testing"
)

func TestDeclaration_String(t *testing.T) {

	tests := []struct{
		Verb string
		Object string
		Expected string
	}{
		{
			Verb: "",
			Object: "",
			Expected: "\u0085",
		},



		{
			Verb: "",
			Object:   "joeblow",
			Expected: "joeblow\u0085",
		},
		{
			Verb: "",
			Object:   "joeblow@example.com",
			Expected: "joeblow@example.com\u0085",
		},



		{
			Verb:      "W",
			Object: "",
			Expected: "!W\u0085",
		},
		{
			Verb:      "W",
			Object:      "joeblow",
			Expected: "!W joeblow\u0085",
		},
		{
			Verb:      "W",
			Object:      "joeblow@example.com",
			Expected: "!W joeblow@example.com\u0085",
		},
		{
			Verb:      "W",
			Object:      "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64",
			Expected: "!W did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u0085",
		},



		{
			Verb:      "PUSHING",
			Object: "",
			Expected: "!PUSHING\u0085",
		},
		{
			Verb:      "PUSHING",
			Object:            "joeblow",
			Expected: "!PUSHING joeblow\u0085",
		},
		{
			Verb:      "PUSHING",
			Object:            "joeblow@example.com",
			Expected: "!PUSHING joeblow@example.com\u0085",
		},
		{
			Verb:      "W",
			Object:      "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64",
			Expected: "!W did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u0085",
		},
	}

	for testNumber, test := range tests {

		var declaration Declaration

		declaration.loaded = true
		declaration.verb  = test.Verb
		declaration.object = test.Object

		var actual   string = declaration.String()
		var expected string = test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
