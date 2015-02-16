package parsers

import "testing"

func TestNew(t *testing.T) {
	foo := New("foo")
	if foo.Data == nil {
		t.Errorf("Data not initialized")
	}
}
