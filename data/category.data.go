package data

import "task-session-1/entity"

var Categories = []entity.Category{
	{
		ID: 1, 
		Name: "Hardware", 
		Description: "Kategori pada hardware",
	},
	{
		ID: 2, 
		Name: "Software", 
		Description: "Kategori pada software",
	},
	{
		ID: 3, 
		Name: "Lainnya", 
		Description: "Kategori pada lainnya",
	},
}
