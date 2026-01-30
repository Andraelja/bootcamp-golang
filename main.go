package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"task-session-1/database"
	"task-session-1/handlers"
	"task-session-1/repositories"
	"task-session-1/services"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// Read ENV
	viper.AutomaticEnv()

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Validasi config
	if config.Port == "" {
		log.Fatal("PORT is required")
	}
	if config.DBConn == "" {
		log.Fatal("DB_CONN is required")
	}

	// Init DB
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("/api/category", categoryHandler.HandleCategory)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)
	http.HandleFunc("/api/product", productHandler.HandleProduct)

	// Start server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("gagal running server:", err)
	}
}
