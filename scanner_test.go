package daniel_test

import (
	"strings"
	"testing"

	"github.com/benbjohnson/sql-parser"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok daniel.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: daniel.EOF},
		{s: `#`, tok: daniel.ILLEGAL, lit: `#`},
		{s: ` `, tok: daniel.WS, lit: " "},
		{s: "\t", tok: daniel.WS, lit: "\t"},
		{s: "\n", tok: daniel.WS, lit: "\n"},

		// Misc characters
		{s: `*`, tok: daniel.ASTERISK, lit: "*"},

		// Identifiers
		{s: `foo`, tok: daniel.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: daniel.IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `FROM`, tok: daniel.FROM, lit: "FROM"},
		{s: `SELECT`, tok: daniel.SELECT, lit: "SELECT"},
	}

	for i, tt := range tests {
		s := daniel.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}

