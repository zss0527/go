package dto

import (
	"bunny-fairy-api/internal/models/home"
)

type CategoryDTO struct {
	ID       string        `json:"id"` // 接收JSON字符串id
	Name     string        `json:"name"`
	Picture  string        `json:"picture"`
	Children []CategoryDTO `json:"children"`
	Goods    []home.Goods  `json:"goods"`
}
