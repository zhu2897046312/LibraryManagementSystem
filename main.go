package main

import (
	"LibraryManagementSystem/routers"
)


func main() {
	r := routers.Router()
	r.Run(":8081")
}