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

func UpdateCategory(id int, category entity.Category) (entity.Category, error) {
	// looping untuk ambil data yang mau diedit
	for i, v := range data.Categories {
		if v.ID == id {
			category.ID = id
			// ambil array yg akan diedit, kemudia set
			data.Categories[i] = category
			return category, nil
		}
	}

	return entity.Category{}, ErrCategoryNotFound
}
