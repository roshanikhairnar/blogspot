package main

import (
	"context"
	"entdemo/ent"
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
	ctx := context.Background()
	// if _, err = start.CreateUser(ctx, client); err != nil {
	// 	log.Fatal(err)
	// }
	// if _, err = start.QueryUser(ctx, client); err != nil {
	// 	log.Fatal(err)
	// }
	// if _, err = start.CreateBlog(ctx, client); err != nil {
	// 	log.Fatal(err)
	// }
	// blog, err := start.QueryBlog(ctx, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	r := gin.Default()
	initRoutes(r)

	r.Run()
}
func initRoutes(r *gin.Engine) {

}
