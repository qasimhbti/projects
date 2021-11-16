# Database - users

Let's create a `User` entity and add some methods that will allow our API to create, update, delete and get user data.

## Database helper functions

We'll need two more unexported functions that will help us interact with the database on disk.

`func (c Client) updateDB(db databaseSchema) error `

The `updateDB` method should save the data in the given databaseSchema to the filepath specificed in the client, overwriting what was there before.

`func (c Client) readDB() (databaseSchema, error) `

The `readDB` function should return a new databaseSchema populated with the latest data from disk.

## User methods

Now we can build the exported functions that our API will call directly!

### CreateUser

`func (c Client) CreateUser(email, password, name string, age int) (User, error)`

We're using `email` as a primary key, that is, we can't have two users with the same email. If you look at the `databaseSchema` struct you'll notice that the `Users` field is a map: `map[string]User`. The `string` key in the map will be the user's email.

This function should read the current state of the database, create a new user struct, add it to the `Users` map in the schema, then update the data on disk. Don't forget to set the `CreatedAt` field to `time.Now().UTC()`.

### UpdateUser

`func (c Client) UpdateUser(email, password, name string, age int) (User, error)`

This function will be similar to CreateUser, but if the user doesn't already exist, it should return an error: `"user doesn't exist""`. It also won't update the `CreatedAt` timestamp, that should be left alone.

### GetUser

`func (c Client) GetUser(email string) (User, error)`

Given the user's email, find the user in the database and return it.

### DeleteUser

`func (c Client) DeleteUser(email string) error`

Given the user's email, delete the user from the database if it exists.

## Assignment

To pass off this step, you can update your sandbox code:

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
}
```
