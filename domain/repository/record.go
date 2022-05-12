package repository

import (
	"github.com/SananGuliyev/gossignment/domain/entity"
	"github.com/SananGuliyev/gossignment/domain/io"
)

type RecordRepository interface {
	FilterByTimeAndAmount(startDate, endDate io.Date, minCount, maxCount int64) ([]*entity.Record, error)
}
