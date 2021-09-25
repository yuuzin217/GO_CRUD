package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	// defer models.DB.Close()
	fmt.Println(models.DB)
	controllers.StartMainServer()
}
