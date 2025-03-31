package main

import (
	"fmt"

	"backend/db"
	"backend/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
