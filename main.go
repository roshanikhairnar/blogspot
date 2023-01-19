package main

import (
	"context"
	auth "entdemo/authentication"
	"entdemo/ent"
	handler "entdemo/handlers"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=roshani dbname=entdemo password='roshani@123'")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ent.EntClient = client
	initRoutes(client)
}

func initRoutes(client *ent.Client) {
	r := gin.Default()

	public := r.Group("/api")

	//public.POST("/register", auth.Register)
	public.POST("/login", auth.Login)
	admin := r.Group("/api/admin")
	admin.Use(auth.JwtAuthMiddleware())
	admin.GET("/user", auth.CurrentUser)

	r.GET("/blogs/:id", handler.GetBlog)
	r.GET("/blogs", handler.GetAllBlogs)
	admin.DELETE("/deleteBlog/:id", handler.DeleteBlog)
	admin.PUT("/updateblog/:id", handler.UpdateBlog)
	admin.POST("/createblog", handler.CreateBlog)

	r.Run()
}
