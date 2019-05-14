package assert

import (
	testify "github.com/stretchr/testify/assert"
	"testing"
)

type Assertions struct {
	*testify.Assertions
}

func New(t *testing.T) *Assertions {
	testify.New(t)
	return &Assertions{
		Assertions: testify.New(t),
	}
}
