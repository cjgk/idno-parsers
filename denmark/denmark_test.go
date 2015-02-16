package denmark_test

import (
	"testing"

	"github.com/steevel/idno-parsers/denmark"
)

func TestValid(t *testing.T) {
	tests := []string{
		"081281-0391", "0812810391",
		"110488-9898", "1104889898",
	}

	for _, tt := range tests {
		if p, err := denmark.Parse(tt); err != nil {
			t.Errorf("Fail at '%v', out '%+v' err %+v", tt, p, err)
		}
	}
}
