package main

import (
	"blogapi/cmd/database"
	"blogapi/cmd/internal/http"

	"github.com/gin-gonic/gin"
)

func main() {
	connectionString := "postgresql://postgres:postgres@localhost:5432/blog_api"
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	g := gin.Default()
	http.Configure(conn)
	http.SetRoutes(g)
	g.Run(":3000")
}
