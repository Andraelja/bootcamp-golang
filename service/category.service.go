package service

import (
	"errors"
	"task-session-1/data"
	"task-session-1/entity"
)

var ErrCategoryNotFound = errors.New("Category Not Found")

func GetAllCategory() []entity.Category {
	return data.Categories
}

func GetCategoryById(id int) (entity.Category, error) {
	// looping untuk mencari ID
	for _, category := range data.Categories {
		if category.ID == id {
			return category, nil
		}
	}

	return entity.Category{}, ErrCategoryNotFound
}

func StoreCategory(category entity.Category) entity.Category {
	category.ID = len(data.Categories) + 1
	data.Categories = append(data.Categories, category)
	return category
}
