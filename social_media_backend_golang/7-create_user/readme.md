# Create user endpoint

Let's get the endpoint for creating users working!

## Ensure the database

At the beginning of `main()`, setup a new database client. I use a `db.json` file in the root of the repo as the path. Be sure to add it to the gitignore!

## apiConfig

Passing api configuration to handlers is best done by creating a "configuration" type, and making the handlers methods on that type, rather than just functions.

I make a simple struct:

```go
type apiConfig struct {
	dbClient database.Client
}
```

Then create an instance of that struct with your database client. We'll use that in our handlers.

## The endpoints

The default `HandleFunc` function doesn't break down requests by HTTP method, and doesn't have explicit support for path variables. We'll be building from scratch for learning purposes, but know that many people use mux libraries like [gorilla](https://github.com/gorilla/mux) in production apps to make routing simpler.

Let's make an "endpoint" method. We'll actually register it twice so that it will trigger on `/users`, the other on `/users/` - note the extra `/`. The extra slash means that requests with a variable in the path will also be routed to the method.

Your method signatures will look something like this:

* `func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request)`

And registering:

```go
m.HandleFunc("/users", apiCfg.endpointUsersHandler)
m.HandleFunc("/users/", apiCfg.endpointUsersHandler)
```

The `endpointUsersHandler` method is just a switch/case statement that routes to a specific "handler" function depending on the HTTP method in the `http.Request` data.

```go
switch r.Method {
case http.MethodGet:
    // call GET handler
case http.MethodPost:
    // call POST handler
case http.MethodPut:
    // call PUT handler
case http.MethodDelete:
    // call DELETE handler
default:
    respondWithError(w, 404, errors.New("method not supported"))
}
```

## The handlers

Let's write the actual handlers, then we'll plug them into the endpoint's switch/case.

### handlerCreateUser

`func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request)`

This handler should be invoked by the client sending a `POST` request to the `/users` endpoint. It takes a JSON body as parameters:

```json
{
    "email": "test@example.com",
    "password": "12345",
    "name": "john doe",
    "age": 18
}
```

I usually unmarshal parameters at the top of the handler method using a [json decoder](https://pkg.go.dev/encoding/json#Decoder):

```go
type parameters struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
    Age      int    `json:"age"`
}
decoder := json.NewDecoder(r.Body)
params := parameters{}
err := decoder.Decode(&params)
if err != nil {
    respondWithError(w, http.StatusBadRequest, err)
    return
}
```

After that, use the parameters to call `apiCfg.dbClient.CreateUser()`. If all goes well, respond with `http.StatusCreated` (201), and a copy of the new user struct marshalled to JSON. Otherwise, at any point an error occurs, call `respondWithError` instead.

## Assignment

To pass off this step, make a `POST` request to http://localhost:8080/users with some valid body data. Make sure your database gets updated as you would expect!
