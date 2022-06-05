package cache

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCacheSet(t *testing.T) {
	c := New[string, int](10)
	assert.NoError(t, c.Set("a", 5))
}

func TestCacheGet(t *testing.T) {
	c := New[string, int](10)
	assert.NoError(t, c.Set("a", 5))
	v, ok := c.Get("a")
	assert.Equal(t, ok, true)
	assert.Equal(t, v, 5)
}

func TestCacheRemove(t *testing.T) {
	c := New[string, int](10)
	assert.NoError(t, c.Set("a", 5))
	assert.NoError(t, c.Remove("a"))
	_, ok := c.Get("a")
	assert.Equal(t, ok, false)
}

func TestCacheSizeCapacity(t *testing.T) {
	c := New[string, int](10)
	assert.NoError(t, c.Set("a", 5))
	assert.NoError(t, c.Set("b", 50))
	assert.NoError(t, c.Set("c", 500))
	assert.Equal(t, c.Capacity(), 10)
	assert.Equal(t, c.Size(), 3)
}