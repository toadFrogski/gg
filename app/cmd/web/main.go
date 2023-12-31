package main

import (
	"flag"
	"fmt"
	config "gg/conf"
	"gg/middlewares"
	"gg/modules/book"
	"gg/modules/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	defaultConfigPath = "config.yml"
)

// @title GG backend API
// @version 1.0
//
// @BasePath /v1
func main() {
	settings := getConfig()
	r := gin.Default()
	db := initDB(&settings.Database)

	initMiddlewares(r)
	initRoutes(r, db)
	r.Run(settings.Server.GetListenAddr())
}

func getConfig() *config.ApplicationConfig {
	configPath := flag.String("config", defaultConfigPath, "config file")
	flag.Parse()
	config, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		panic(err)
	}

	fmt.Printf("Config loaded: %+v\n", config)
	return config
}

func initDB(dbconfig *config.DatabaseConfig) *gorm.DB {
	db, err := config.ConnectToDB(dbconfig)
	if err != nil {
		fmt.Printf("Error connect to DB: %+v\n", err)
		panic(err)
	}

	return db
}

func initRoutes(r *gin.Engine, db *gorm.DB) {
	user.GetAuthRoutes(r, db)
	v1 := r.Group("/v1")
	{
		book.GetBooksRoutes(v1, db)
		user.GetUserRoutes(v1, db)
	}
}

func initMiddlewares(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.PanicHandleMiddleware())
}
