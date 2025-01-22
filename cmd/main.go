package main

import (
	"log"

	"github.com/Fairuzzzzz/taskapp/internal/configs"
	favoritesHandler "github.com/Fairuzzzzz/taskapp/internal/handlers/favorites"
	recipesHandler "github.com/Fairuzzzzz/taskapp/internal/handlers/recipes"
	"github.com/Fairuzzzzz/taskapp/internal/middleware"
	"github.com/Fairuzzzzz/taskapp/internal/models/favorites"
	"github.com/Fairuzzzzz/taskapp/internal/models/recipes"
	"github.com/Fairuzzzzz/taskapp/internal/models/users"
	favoritesRepo "github.com/Fairuzzzzz/taskapp/internal/repository/favorites"
	recipesRepo "github.com/Fairuzzzzz/taskapp/internal/repository/recipes"
	favoritesSvc "github.com/Fairuzzzzz/taskapp/internal/services/favorites"
	recipesSvc "github.com/Fairuzzzzz/taskapp/internal/services/recipes"
	"github.com/Fairuzzzzz/taskapp/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.SetDefaultUser())

	var cfg *configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Initialize error", err)
	}

	cfg = configs.Get()
	db, err := internalsql.Connect(cfg.Database.Datasourcename)
	if err != nil {
		log.Fatalf("failed to connect to database, err : %+v\n", err)
	}

	db.AutoMigrate(&recipes.Recipe{})
	db.AutoMigrate(&favorites.Favorites{})
	db.AutoMigrate(&users.User{})

	var defaultUser users.User
	if err := db.First(&defaultUser).Error; err != nil {
		defaultUser = users.User{
			Name: "Default User",
		}
		db.Create(&defaultUser)
	}

	recipeRepo := recipesRepo.NewRepository(db)
	favoriteRepo := favoritesRepo.NewRepository(db)

	recipeSvc := recipesSvc.NewService(cfg, recipeRepo)
	favoriteSvc := favoritesSvc.NewService(cfg, favoriteRepo)

	recipeHandler := recipesHandler.NewHandler(r, recipeSvc)
	recipeHandler.RegisterRoute()

	favoriteHandler := favoritesHandler.NewHandler(r, favoriteSvc)
	favoriteHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
