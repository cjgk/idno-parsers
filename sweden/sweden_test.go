package sweden_test

import (
	"testing"

	"github.com/steevel/idno-parsers/sweden"
)

func TestValid(t *testing.T) {
	tests := []string{
		"19811208-0091", "811208-0091", "8112080091",
		"19880411-0149", "880411-0149", "8804110149",
	}

	for _, tt := range tests {
		if p, err := sweden.Parse(tt); err != nil {
			t.Errorf("Fail at '%v', out '%+v', err: %+v", tt, p, err)
		}
	}
}
