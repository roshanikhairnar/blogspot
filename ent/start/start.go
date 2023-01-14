package ent

import (
	"context"
	"entdemo/ent"
	user "entdemo/ent/user"
	"fmt"
	"log"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetName("ayush").
		SetType("admin").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
func CreateBlog(ctx context.Context, client *ent.Client) (*ent.Blogs, error) {
	u, err := client.Blogs.
		Create().
		SetBlogTitle("new trends").
		SetBlogType("fashion").
		SetBlogContent("jeans top are out of fashion....saree in").
		SetBlogAuthor("ayush").
		SetOwnerID(1).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating blog: %w", err)
	}
	log.Println("blog was created: ", u)
	return u, nil
}
func QueryBlog(ctx context.Context, client *ent.Client) (string, error) {
	// u, err := client.Blogs.
	// 	Query().
	// 	Where((blogs.BlogTitleIn("new trends"))).
	// 	Only(ctx)
	v, err := client.Blogs.Get(ctx, 1)
	if err != nil {
		return "", fmt.Errorf("failed querying blog: %w", err)
	}
	log.Println("blogs returned: ", v)
	return v.String(), nil
}
