package start

import (
	"context"
	"entdemo/ent"
	user "entdemo/ent/user"
	"fmt"
	"log"
)

func QueryUserByName(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Username(name)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
func QueryUserById(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.ID(id)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
func CreateBlog(client *ent.Client, newBlog *ent.Blogs) (*ent.Blogs, error) {
	ctx := context.Background()
	u, err := client.Blogs.
		Create().
		SetBlogTitle(*newBlog.BlogTitle).
		SetBlogType(*newBlog.BlogType).
		SetBlogContent(*newBlog.BlogContent).
		SetBlogAuthor(*newBlog.BlogAuthor).
		//SetOwnerID().
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating blog: %w", err)
	}
	log.Println("blog was created: ", u)

	return u, nil
}
func QueryBlog(ctx context.Context, client *ent.Client, id int) (string, error) {
	blog, err := client.Blogs.Get(ctx, id)
	if err != nil {
		return "", fmt.Errorf("failed querying blog: %w", err)
	}
	log.Println("blog returned: ", blog)
	return blog.String(), nil
}
func QueryAllBlogs(ctx context.Context, client *ent.Client) ([]*ent.Blogs, error) {
	blogs, err := client.Blogs.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying blog: %w", err)
	}
	log.Println("blogs returned: ", blogs)
	return blogs, nil
}

func DeleteBlog(client *ent.Client, id int) error {
	err := client.Blogs.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func UpdateBlog(client *ent.Client, id int, title string) (*ent.Blogs, error) {
	ctx := context.Background()
	_, err := QueryBlog(ctx, client, id)
	if err != nil {
		return nil, err
	}
	updatedBlog, err := client.Blogs.UpdateOneID(id).SetBlogTitle(title).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedBlog, nil
}
