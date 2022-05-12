package io

type MemInput struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type MemOutput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
