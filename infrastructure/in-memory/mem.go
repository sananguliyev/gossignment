package in_memory

import (
	"github.com/SananGuliyev/gossignment/domain/entity"
	"github.com/SananGuliyev/gossignment/domain/repository"
	"github.com/SananGuliyev/gossignment/domain/throwable"
	"github.com/patrickmn/go-cache"
)

type memRepository struct {
	cache *cache.Cache
}

func NewMemRepository(cache *cache.Cache) repository.MemRepository {
	return &memRepository{cache: cache}
}

func (r *memRepository) Create(mem *entity.Mem) error {
	r.cache.Set(mem.Key, mem, cache.NoExpiration)

	return nil
}

func (r *memRepository) GetByKey(key string) (*entity.Mem, error) {
	if x, found := r.cache.Get(key); found {
		res := x.(*entity.Mem)
		return res, nil
	}

	return nil, throwable.NewNotFoundError("Key does not exist")
}
