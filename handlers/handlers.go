package handlers

import (
	"context"
	"entdemo/ent"
	start "entdemo/start"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBlog(gctx *gin.Context) {
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
func GetAllBlogs(gctx *gin.Context) {
	ctx := context.Background()

	blogs, err := start.QueryAllBlogs(ctx, ent.EntClient)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"data": blogs})
}

func DeleteBlog(gctx *gin.Context) {
	id := gctx.Param("id")
	blogID, _ := strconv.Atoi(id)
	err := start.DeleteBlog(ent.EntClient, blogID)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, "Successfully deleted the blog")
}
func UpdateBlog(gctx *gin.Context) {
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
	gctx.JSON(http.StatusOK, "Successfully updated the blog")
}
func CreateBlog(gctx *gin.Context) {
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
	gctx.JSON(http.StatusOK, "Successfully created the blog")
}
