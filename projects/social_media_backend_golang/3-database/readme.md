# Database

Let's build a simple database package. Our database package will simply read and save our user's data in a `JSON` file on disk. If this were a database project we'd do something a bit more scalable!

## Internal folder

In go, [packages](https://qvault.io/golang/how-to-separate-library-packages-in-go/) that aren't meant to be used by other modules are stored in an `internal` folder. The Go toolchain won't allow modules that don't contain that `internal` folder to use that code, so it's good practice to scope packages accordingly.

Create an `internal` directory at the root of your project, then add a `database` directory with a `database.go` file. Make sure to add `package database` at the top of that file.

## The database package

### Client type

Your database package needs to [export](https://www.ardanlabs.com/blog/2014/03/exportedunexported-identifiers-in-go.html) a `Client` struct type. All the database interactions our API will use will be methods on that struct. The struct only needs a single field to hold the file path to the database. For example, I used `./db.json` in the same directory.

## NewClient function

Create a `NewClient` function that creates an instance of a `Client`. It should take a path and return a client struct. This kind of "initializer" pattern shows the user of the database package exactly how to make a new `Client` with no confusion.

## Database schema

We need a structure that will define our JSON schema in the database file. In Go, the easiest way to do this is by using [structs with JSON tags](https://qvault.io/golang/json-golang/). Also, to kill two birds with one stone, we can use these same JSON tags later in our HTTP responses from the API.

```go
type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}
```

We'll export the `User` and `Post` types so that the users of the package can interact with them later.

## createDB method

`func (c Client) createDB() error`

This unexported function creates a mew database at the path in the `Client` instance. It will overwrite any previous data in the file if it existed previously.

You'll want to use [os.WriteFile](https://pkg.go.dev/os#WriteFile) and [json.Marshal](https://qvault.io/golang/json-golang/#marshal-json).

## EnsureDB method

`(c Client) EnsureDB() error`

This exported method should check if a database already exists, you can use [os.ReadFile](https://pkg.go.dev/os#ReadFile). If it does, all is good! Otherwise, it should create a new database using `createDB`.

## Assignment

We'll get to unit tests shortly, while building this `database` package however, I recommend creating a little `sandbox` directory in the root of your project, and adding a new `main.go` file there. You can use this new sandbox program to test out your database code without having to make strange changes to the API itself.

To pass off this step, run your sandbox code to test out the various functionality we added. Play with it a bit! I used something like this:

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
	fmt.Println("database ensured!")
}
```
