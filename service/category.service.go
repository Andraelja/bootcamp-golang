package service

import (
	"errors"
	"task-session-1/entity"
	"task-session-1/data"
)

var ErrCategoryNotFound = errors.New("Category Not Found")

func GetAllCategory() []entity.Category {
	return data.Categories
}