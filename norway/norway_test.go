package norway_test

import (
	"testing"

	"github.com/steevel/idno-parsers/norway"
)

func TestValid(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"08128149946", norway.Regular},
	}

	for _, tt := range tests {
		if p, _ := norway.Parse(tt.in); !p.Valid || p.Data["type"] != tt.out {
			t.Errorf("Invalid idno, Expected %v Got %v", p.Data["type"], tt.out)
		}
	}
}

func TestInvalid(t *testing.T) {
	tests := []string{"08128149945", "08128149947"}

	for _, tt := range tests {
		if p, _ := norway.Parse(tt); p.Valid {
			t.Errorf("Valid idno, Expected an invalid")
		}
	}
}
