# Building

By the end of this step you'll have a working HTTP server! Don't worry, it's not much code! We don't even need to use any frameworks, just the good 'ol standard library.

## Boilerplate

You'll need a `main.go` file. It should be part of `package main` and have a main function. You should be able to run `go build` with just that and get a runnable program that does nothing but exit with a succesful status code.

Next, you'll need to [create a Go module](https://golang.org/doc/tutorial/create-module) at the root of your project. You can do this by:

`go mod init MODULE_NAME`

I recommend naming the module by it's remote git locaiton (you should store all your projects in Git!). For example, my GitHub name is `wagslane` so my module name might be `github.com/wagslane/http_server_golang`.

## The serve mux

Within you're `main` function, you'll need to create a [ServeMux](https://pkg.go.dev/net/http#ServeMux) using [http.NewServeMux](https://pkg.go.dev/net/http#NewServeMux)

A `ServeMux` is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

## The handler

Next, you'll add a single handler function and register it to be called at the root path `"/"`. You'll use the ServeMux's [HandleFunc method](https://pkg.go.dev/net/http#ServeMux.HandleFunc).

You can use this code for the handler function itself:

```go
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
```

This simple handler function will add a `Content-Type` header, a status code, and am empty JSON body to the response.

## The server

Lastly, we'll create a simple [http.Server](https://pkg.go.dev/net/http#Server)

```go
const addr = "localhost:8080"
srv := http.Server{
    Handler:      m,
    Addr:         addr,
    WriteTimeout: 30 * time.Second,
    ReadTimeout:  30 * time.Second,
}
```

Once all that's done, call the [ListenAndServe](https://pkg.go.dev/net/http#Server.ListenAndServe) method on the server struct. It will block forever and serve your HTTP handler!

## Assignment

To pass off this step, make a request to [http://localhost:8080](http://localhost:8080) and see if you get an empty JSON body.
