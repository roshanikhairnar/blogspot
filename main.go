package main

import (
	"context"
	"entdemo/ent"
	auth "entdemo/ent/authentication"
	start "entdemo/ent/start"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	public.POST("/register", auth.Register)
	public.POST("/login", auth.Login)
	admin := r.Group("/api/admin")
	admin.Use(auth.JwtAuthMiddleware())
	admin.GET("/user", auth.CurrentUser)

	r.GET("/blogs/:id", getBlog)
	r.GET("/blogs", getAllBlogs)
	admin.DELETE("/deleteBlog/:id", deleteBlog)
	admin.PUT("/updateblog/:id", updateBlog)
	admin.POST("/createblog", createBlog)

	r.Run()
}
func getBlog(gctx *gin.Context) {
	ctx := context.Background()
	id := gctx.Param("id")
	blogID, _ := strconv.Atoi(id)
	blog, err := start.QueryBlog(ctx, ent.EntClient, blogID)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"data": blog})
}
func getAllBlogs(gctx *gin.Context) {
	ctx := context.Background()

	blogs, err := start.QueryAllBlogs(ctx, ent.EntClient)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"data": blogs})
}

func deleteBlog(gctx *gin.Context) {
	id := gctx.Param("id")
	blogID, _ := strconv.Atoi(id)
	err := start.DeleteBlog(ent.EntClient, blogID)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, "Success")
}
func updateBlog(gctx *gin.Context) {
	id := gctx.Param("id")
	blogID, _ := strconv.Atoi(id)
	blog := new(struct {
		BlogTitle string
	})
	if err := gctx.ShouldBindJSON(&blog); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blogupdated, err := start.UpdateBlog(ent.EntClient, blogID, blog.BlogTitle)
	fmt.Println("updatedblog", blogupdated)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
func createBlog(gctx *gin.Context) {
	newBlog := new(ent.Blogs)
	if err := gctx.ShouldBindJSON(&newBlog); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBlog, err := start.CreateBlog(ent.EntClient, newBlog)
	fmt.Println("newblog", newBlog)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
