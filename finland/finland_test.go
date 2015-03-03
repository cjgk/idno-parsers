package finland_test

import (
	"testing"

	"github.com/steevel/idno-parsers/finland"
)

func TestValid(t *testing.T) {
	tests := []string{"311280-888Y"}

	for _, tt := range tests {
		if p, err := finland.Parse(tt); err != nil {
			t.Errorf("Fail at '%v', out '%+v', err: %+v", tt, p, err)
		}
	}
}

func TestInvalid(t *testing.T) {
	tests := []string{"311280-888X"}

	for _, tt := range tests {
		if p, err := finland.Parse(tt); err == nil {
			t.Errorf("Fail at '%v', out '%+v', should have produced error", tt, p)
		}
	}
}
