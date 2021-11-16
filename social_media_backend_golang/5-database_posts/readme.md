# Database - posts

Let's create a `Post` entity so our users can microblog about their cats!

## UUID dependency

We'll be using [UUIDs](https://qvault.io/clean-code/what-are-uuids-and-should-you-use-them/) as our post IDs. The UUIDs ensure that no two posts have the same ID. We'll use [Google's UUID library](https://pkg.go.dev/github.com/google/uuid). In Go, to add a dependency, you use `go get MODULE_PATH`, where the `MODULE_PATH` is typically the URL of the module's git repo.

`go get github.com/google/uuid`

You'll notice that your `go.mod` file will have a new dependency and a `go.sum` file is created in your project to keep track of your dependencies versions.

## Posts methods

### CreatePost

Make sure to import `"github.com/google/uuid"` at the top of the file you'll be writing this method in.

`func (c Client) CreatePost(userEmail, text string) (Post, error)`

1. Read the current database state
2. Make sure the user exists in the database using the `userEmail`
3. Create a new post struct. Set the `CreatedAt` to now, and create a new uuid using the google package: `id := uuid.New().String()`
4. Update the database on disk

### GetPosts

`func (c Client) GetPosts(userEmail string) ([]Post, error) `

1. Read the current database state
2. Iterate through all the posts and add each post from the given user to a new slice
3. Return the matching posts

### DeletePost

`func (c Client) DeletePost(id string) error`

Given the post id, delete it.

## Assignment

To pass off this step, you should update your sandbox code:

```go
package main

import (
	"fmt"
	"log"

	"github.com/qvault/courses/projects/http_server_golang/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	user, err := c.CreateUser("test@example.com", "password", "john doe", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user created", user)

	updatedUser, err := c.UpdateUser("test@example.com", "password", "john doe", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user updated", updatedUser)

	gotUser, err := c.GetUser("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user got", gotUser)

	err = c.DeleteUser("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user deleted")

	_, err = c.GetUser("test@example.com")
	if err == nil {
		log.Fatal("shouldn't be able to get user that was deleted")
	}
	fmt.Println("user confirmed deleted")

	user, err = c.CreateUser("test@example.com", "password", "john doe", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user recreated", user)

	post, err := c.CreatePost("test@example.com", "my cat is way too fat")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("post created", post)

	secondPost, err := c.CreatePost("test@example.com", "my cat is getting skinny now")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("another post created", secondPost)

	posts, err := c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeletePost(post.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted first post", posts)

	posts, err = c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeletePost(secondPost.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted second post", posts)

	posts, err = c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeleteUser("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user redeleted")
}
```
