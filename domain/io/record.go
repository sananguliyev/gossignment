package io

import (
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func (t *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type RecordFilterInput struct {
	StartDate Date  `json:"startDate" validate:"required"`
	EndDate   Date  `json:"endDate" validate:"required"`
	MinCount  int64 `json:"minCount" validate:"required"`
	MaxCount  int64 `json:"maxCount" validate:"required"`
}

type RecordOutput struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

type RecordsOutput struct {
	Records []*RecordOutput `json:"records"`
}
