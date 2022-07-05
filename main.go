package main

import (
	"api-redeem-point/api"
	"api-redeem-point/app/modules"
	"api-redeem-point/config"
	"api-redeem-point/repository"
	"api-redeem-point/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "api-redeem-point/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title API Poins ID
// @version 1.0
// @description Berikut API Poins ID
// @host api-poins-id.herokuapp.com/v1
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	port := os.Getenv("PORT")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := config.GetConfig()
	dbCon := utils.NewConnectionDatabase(config)

	defer dbCon.CloseConnection()

	controllers := modules.RegistrationModules(dbCon, config)
	dbCon.Postgres.AutoMigrate(&repository.History_Transaction{})
	dbCon.Postgres.AutoMigrate(&repository.Store{})
	dbCon.Postgres.AutoMigrate(&repository.Customer{})
	dbCon.Postgres.AutoMigrate(&repository.StockProduct{})
	dbCon.Postgres.AutoMigrate(&repository.Admin{})

	e := echo.New()
	handleSwagger := echoSwagger.WrapHandler
	e.GET("/swagger/*", handleSwagger)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})
	api.RegistrationPath(e, controllers)

	go func() {
		if port == "" {
			port = "8080"
		}
		address := fmt.Sprintf(":%s", port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit
}
