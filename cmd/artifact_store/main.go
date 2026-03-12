package main

import (
	"fmt"
	"artifact-store/internal/api"
	"artifact-store/internal/handler"
  "github.com/gin-gonic/gin"
)

func main() {
	handler := &handler.Server{}

	r := gin.Default()

	api.RegisterHandlers(r, handler)

	fmt.Printf("Starting server...")
	r.Run(":8080")
}
