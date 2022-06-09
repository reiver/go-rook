package intr

import (
	"testing"
)

func TestInterrogation_String(t *testing.T) {

	tests := []struct{
		Verb string
		Object string
		Expected string
	}{
		{
			Verb: "",
			Object: "",
			Expected: "\r\n",
		},



		{
			Verb: "",
			Object:   "joeblow",
			Expected: "joeblow\r\n",
		},
		{
			Verb: "",
			Object:   "joeblow@example.com",
			Expected: "joeblow@example.com\r\n",
		},



		{
			Verb:      "W",
			Object: "",
			Expected: "?W\r\n",
		},
		{
			Verb:      "W",
			Object:      "joeblow",
			Expected: "?W joeblow\r\n",
		},
		{
			Verb:      "W",
			Object:      "joeblow@example.com",
			Expected: "?W joeblow@example.com\r\n",
		},
		{
			Verb:      "W",
			Object:      "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64",
			Expected: "?W did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\r\n",
		},



		{
			Verb:      "PUSHING",
			Object: "",
			Expected: "?PUSHING\r\n",
		},
		{
			Verb:      "PUSHING",
			Object:            "joeblow",
			Expected: "?PUSHING joeblow\r\n",
		},
		{
			Verb:      "PUSHING",
			Object:            "joeblow@example.com",
			Expected: "?PUSHING joeblow@example.com\r\n",
		},
		{
			Verb:      "W",
			Object:      "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64",
			Expected: "?W did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\r\n",
		},
	}

	for testNumber, test := range tests {

		var interrogation Interrogation

		interrogation.loaded = true
		interrogation.verb   = test.Verb
		interrogation.object = test.Object

		var actual   string = interrogation.String()
		var expected string = test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
