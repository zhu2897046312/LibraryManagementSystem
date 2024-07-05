package main

import (
	"LibraryManagementSystem/models"
	"LibraryManagementSystem/routers"
	"LibraryManagementSystem/utils"
)


func main() {
	utils.DB_MySQL.AutoMigrate(&models.Borrower{})
	utils.DB_MySQL.AutoMigrate(&models.Borrowing{})
	r := routers.Router()
	r.Run(":8081")
}