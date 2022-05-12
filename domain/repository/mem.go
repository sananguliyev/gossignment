package repository

import "github.com/SananGuliyev/gossignment/domain/entity"

type MemRepository interface {
	Create(mem *entity.Mem) error
	GetByKey(key string) (*entity.Mem, error)
}
