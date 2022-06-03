package cache

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCacheSet(t *testing.T) {
	c := New[string, int](10)
	assert.NoError(t, c.Set("a", 5))
}