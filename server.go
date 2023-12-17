package main

import (
	"github.com/dipeshdulal/clean-gin/bootstrap"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	//err := godotenv.Load(filepath.Join("/mnt/c/Users/84339/GolandProjects/go-techlens", ".env"))
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bootstrap.RootApp.Execute()
}
