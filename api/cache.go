package api

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func init() {
	c = cache.New(100*time.Minute, 101*time.Minute)
}

// SetCache function
func SetCache(data Post) {
	c.Set(data.ID, data, cache.DefaultExpiration)
}

// GetCache function
func GetCache(id string) (interface{}, bool) {
	return c.Get(id)
}
