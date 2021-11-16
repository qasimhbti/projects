# HTTP Response Helpers

Let's shift our focus back to the main API application now. We should make some response helpers, they'll simplify the logic in our endpoints when we get to those.

## respondWithJSON

Let's build a little function that will handle the standard HTTP stuff for us.

`func respondWithJSON(w http.ResponseWriter, code int, payload interface{})`

An endpoint would call this function once per request, right before returning. It will write some standard HTTP headers, marshal an interface into JSON bytes, and write all that to the response along with a status code.

### Add headers first

`w.Header().Set(key, value)`

### Write JSON body

```go
response, err := json.Marshal(payload)
// deal with err ...
w.Write(dat)
```

### Write status code

`w.WriteHeader(code)`

## Updating the handler

Let's make sure it works by altering the `testHandler` we wrote:

```go
func testHandler(w http.ResponseWriter, r *http.Request) {
    // you can use any compatible type, but let's use our database package's User type for practice
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}
```

## respondWithError

I like to wrap `respondWithJSON` with another function to standardize API errors.

`respondWithError(w http.ResponseWriter, code int, err error)`

You'll want to make a standard error format. I use the following:

```go
type errorBody struct {
	Error string `json:"error"`
}
```

This function should take the `err`, create a new `errorBody`, then call `respondWithJSON`.

## New error handler for testing

Let's test `respondWithError` by adding a new handler function called `testErrHandler` and connecting it to a new path in our API:

`m.HandleFunc("/err", testErrHandler)`

## Assignment

To pass off this step, make GET requests to [http://localhost:8080](http://localhost:8080) and [http://localhost:8080/err](http://localhost:8080/err). Make sure you get the responses you expect!
