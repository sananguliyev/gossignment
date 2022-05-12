package in_memory

import (
	"github.com/patrickmn/go-cache"
	"time"
)

func NewStorage() *cache.Cache {
	db := cache.New(5*time.Minute, 10*time.Minute)
	return db
}
