package sntn_test

import (
	"github.com/reiver/go-rook/internal/token/dclr"
	"github.com/reiver/go-rook/internal/token/impr"
	"github.com/reiver/go-rook/internal/token/intr"
	"github.com/reiver/go-rook/internal/sntn"

	"io"
	"strings"

	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct{
		Input string
		Expected sntn.Sentence
	}{
		{
			"\u000A",
			dclr.Something("IDLING", ""),
		},
		{
			"\u000B",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u000C",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u000D",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u000D\u000A",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u0085",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u2028",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u2029",
			dclr.Something("DOROOD", ""),
		},



		{
			"\u000Aapple banana cherry",
			dclr.Something("IDLING", ""),
		},
		{
			"\u000Bapple banana cherry",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u000Capple banana cherry",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u000Dapple banana cherry",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u000D\u000Aapple banana cherry",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u0085apple banana cherry",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u2028apple banana cherry",
			dclr.Something("DOROOD", ""),
		},
		{
			"\u2029apple banana cherry",
			dclr.Something("DOROOD", ""),
		},



		{
			"!SOMETHING\u000A",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u000B",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u000C",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u000D",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u000D\u000A",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u0085",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u2028",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u2029",
			dclr.Something("SOMETHING", ""),
		},



		{
			"!SOMETHING\u000Aapple banana cherry",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u000Bapple banana cherry",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u000Capple banana cherry",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u000Dapple banana cherry",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u000D\u000Aapple banana cherry",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u0085apple banana cherry",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u2028apple banana cherry",
			dclr.Something("SOMETHING", ""),
		},
		{
			"!SOMETHING\u2029apple banana cherry",
			dclr.Something("SOMETHING", ""),
		},



		{
			"/SOMETHING\u000A",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u000B",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u000C",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u000D",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u000D\u000A",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u0085",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u2028",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u2029",
			impr.Something("SOMETHING", ""),
		},



		{
			"/SOMETHING\u000Aapple banana cherry",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u000Bapple banana cherry",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u000Capple banana cherry",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u000Dapple banana cherry",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u000D\u000Aapple banana cherry",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u0085apple banana cherry",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u2028apple banana cherry",
			impr.Something("SOMETHING", ""),
		},
		{
			"/SOMETHING\u2029apple banana cherry",
			impr.Something("SOMETHING", ""),
		},



		{
			"?SOMETHING\u000A",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u000B",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u000C",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u000D",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u000D\u000A",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u0085",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u2028",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u2029",
			intr.Something("SOMETHING", ""),
		},



		{
			"?SOMETHING\u000Aapple banana cherry",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u000Bapple banana cherry",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u000Capple banana cherry",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u000Dapple banana cherry",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u000D\u000Aapple banana cherry",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u0085apple banana cherry",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u2028apple banana cherry",
			intr.Something("SOMETHING", ""),
		},
		{
			"?SOMETHING\u2029apple banana cherry",
			intr.Something("SOMETHING", ""),
		},



		{
			"!SOMETHING one\u000A",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u000B",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u000C",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u000D",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u000D\u000A",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u0085",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u2028",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u2029",
			dclr.Something("SOMETHING", "one"),
		},



		{
			"!SOMETHING one\u000Aapple banana cherry",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u000Bapple banana cherry",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u000Capple banana cherry",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u000Dapple banana cherry",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u000D\u000Aapple banana cherry",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u0085apple banana cherry",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u2028apple banana cherry",
			dclr.Something("SOMETHING", "one"),
		},
		{
			"!SOMETHING one\u2029apple banana cherry",
			dclr.Something("SOMETHING", "one"),
		},



		{
			"/SOMETHING one\u000A",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u000B",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u000C",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u000D",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u000D\u000A",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u0085",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u2028",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u2029",
			impr.Something("SOMETHING", "one"),
		},



		{
			"/SOMETHING one\u000Aapple banana cherry",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u000Bapple banana cherry",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u000Capple banana cherry",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u000Dapple banana cherry",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u000D\u000Aapple banana cherry",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u0085apple banana cherry",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u2028apple banana cherry",
			impr.Something("SOMETHING", "one"),
		},
		{
			"/SOMETHING one\u2029apple banana cherry",
			impr.Something("SOMETHING", "one"),
		},



		{
			"?SOMETHING one\u000A",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u000B",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u000C",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u000D",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u000D\u000A",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u0085",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u2028",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u2029",
			intr.Something("SOMETHING", "one"),
		},



		{
			"?SOMETHING one\u000Aapple banana cherry",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u000Bapple banana cherry",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u000Capple banana cherry",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u000Dapple banana cherry",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u000D\u000Aapple banana cherry",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u0085apple banana cherry",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u2028apple banana cherry",
			intr.Something("SOMETHING", "one"),
		},
		{
			"?SOMETHING one\u2029apple banana cherry",
			intr.Something("SOMETHING", "one"),
		},



		{
			"!SOMETHING one two\u000A",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u000B",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u000C",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u000D",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u000D\u000A",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u0085",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u2028",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u2029",
			dclr.Something("SOMETHING", "one two"),
		},



		{
			"!SOMETHING one two\u000Aapple banana cherry",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u000Bapple banana cherry",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u000Capple banana cherry",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u000Dapple banana cherry",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u000D\u000Aapple banana cherry",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u0085apple banana cherry",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u2028apple banana cherry",
			dclr.Something("SOMETHING", "one two"),
		},
		{
			"!SOMETHING one two\u2029apple banana cherry",
			dclr.Something("SOMETHING", "one two"),
		},



		{
			"/SOMETHING one two\u000A",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u000B",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u000C",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u000D",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u000D\u000A",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u0085",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u2028",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u2029",
			impr.Something("SOMETHING", "one two"),
		},



		{
			"/SOMETHING one two\u000Aapple banana cherry",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u000Bapple banana cherry",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u000Capple banana cherry",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u000Dapple banana cherry",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u000D\u000Aapple banana cherry",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u0085apple banana cherry",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u2028apple banana cherry",
			impr.Something("SOMETHING", "one two"),
		},
		{
			"/SOMETHING one two\u2029apple banana cherry",
			impr.Something("SOMETHING", "one two"),
		},



		{
			"?SOMETHING one two\u000A",
			intr.Something("SOMETHING", "one two"),
		},
		{
			"?SOMETHING one two\u000B",
			intr.Something("SOMETHING", "one two"),
		},
		{
			"?SOMETHING one two\u000C",
			intr.Something("SOMETHING", "one two"),
		},
		{
			"?SOMETHING one two\u000D",
			intr.Something("SOMETHING", "one two"),
		},
		{
			"?SOMETHING one two\u000D\u000A",
			intr.Something("SOMETHING", "one two"),
		},
		{
			"?SOMETHING one two\u0085",
			intr.Something("SOMETHING", "one two"),
		},
		{
			"?SOMETHING one two\u2028",
			intr.Something("SOMETHING", "one two"),
		},
		{
			"?SOMETHING one two\u2029",
			intr.Something("SOMETHING", "one two"),
		},



		{
			"joeblow\u000A",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u000B",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u000C",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u000D",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u000D\u000A",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u0085",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u2028",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u2029",
			impr.Something("PULL", "joeblow"),
		},



		{
			"joeblow\u000Aapple banana cherry",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u000Bapple banana cherry",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u000Capple banana cherry",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u000Dapple banana cherry",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u000D\u000Aapple banana cherry",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u0085apple banana cherry",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u2028apple banana cherry",
			impr.Something("PULL", "joeblow"),
		},
		{
			"joeblow\u2029apple banana cherry",
			impr.Something("PULL", "joeblow"),
		},



		{
			"joeblow@example.com\u000A",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u000B",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u000C",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u000D",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u000D\u000A",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u0085",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u2028",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u2029",
			impr.Something("PULL", "joeblow@example.com"),
		},



		{
			"joeblow@example.com\u000Aapple banana cherry",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u000Bapple banana cherry",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u000Capple banana cherry",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u000Dapple banana cherry",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u000D\u000Aapple banana cherry",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u0085apple banana cherry",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u2028apple banana cherry",
			impr.Something("PULL", "joeblow@example.com"),
		},
		{
			"joeblow@example.com\u2029apple banana cherry",
			impr.Something("PULL", "joeblow@example.com"),
		},



		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000A",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000B",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000C",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000D",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000D\u000A",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u0085",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u2028",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u2029",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},



		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000Aapple banana cherry",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000Bapple banana cherry",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000Capple banana cherry",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000Dapple banana cherry",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u000D\u000Aapple banana cherry",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u0085apple banana cherry",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u2028apple banana cherry",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
		{
			"did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64\u2029apple banana cherry",
			impr.Something("PULL", "did:ed25519:SFmJz897zApIZaAsTql4mAELQ5mH7nDZkq8t32i_Zr8.base64"),
		},
	}


	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Input)
		if nil == reader {
			t.Errorf("For test #%d, nil reader", testNumber)
			t.Logf("INPUT: %q", test.Input)
			continue
		}

		var sentence sntn.Sentence
		{
			var err error

			sentence, err = sntn.Parse(reader)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("INPUT: %q", test.Input)
				continue
			}
			if nil == sentence {
				t.Errorf("For test #%d, nil sentence", testNumber)
				t.Logf("INPUT: %q", test.Input)
				continue
			}
		}

		{
			var expected string = test.Expected.String()
			var actual   string = sentence.String()

			if expected != actual {
				t.Errorf("For test #%d, actual value is not what was expected", testNumber)
				t.Logf("INPUT: %q", test.Input)
				t.Logf("SENTENCE: %q", sentence)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}
