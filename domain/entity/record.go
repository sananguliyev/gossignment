package entity

import "time"

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	Counts     []int     `json:"counts"`
	Value      string    `json:"value"`
	TotalCount *int      `json:"totalCount"`
}
