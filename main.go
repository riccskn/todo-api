package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"todo-api/config"
	"todo-api/router"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("ERROR: could not load the .env file")
	}

	db := config.Initialize(os.Getenv("DB_DNS"))

	r := gin.Default()

	router.Initialize(r, db)

	err = r.Run()
	if err != nil {
		log.Fatal("ERROR: could not run the server")
	}

}
