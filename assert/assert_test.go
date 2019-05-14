package assert_test

import (
	"github.com/jdxcode/goplay/assert"
	"testing"
)

func TestAssert(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("a", "b")
}
