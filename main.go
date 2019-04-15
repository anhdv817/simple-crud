package main

import (
	"github.com/gin-gonic/gin"
	"simple-crud/pkg/controllers/users"
)

func main() {
	// TODO: Init database
	r := gin.Default()
	api := r.Group("/api")
	users.Register(api.Group("/users"))

	r.Run()
}