package finland

import (
	"fmt"

	"github.com/steevel/idno-parsers"
)

func init() {
	parsers.Register("finland", Parse)
}

func Parse(id string) (*parsers.IdNo, error) {
	return nil, fmt.Errorf("Not implemented yet")
}
