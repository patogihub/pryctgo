package main

import "github.com/patogihub/pryctgo/todo-api/routes"

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
