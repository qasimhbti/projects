# Posts Endpoints

Last but not least, let's hook up the posts endpoints.

You'll need similar `/posts` and `/posts/` registrations for the new posts endpoint that you can setup the same way we did the users endpoint using a switch/case statement.

## handlerCreatePost

`func (apiCfg apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request)`

This handler should be invoked by the client sending a `POST` request to the `/posts` endpoint. It will take the following request body parameters.

```go
type parameters struct {
    UserEmail string `json:"userEmail"`
    Text      string `json:"text"`
}
```

Use `apiCfg.dbClient.CreatePost` and return the created data following the same pattern we used for the users endpoints.

## handlerDeletePost

This one works in a similar manner to get user. The `uuid` of the post is passed in as a route parameter. The client would send a `DELETE` request to `/posts/92620acc-6c31-48da-84e7-693ac7e59e48` for example.

## handlerRetrievePosts

This one will return an array of posts, more specifically, all the posts of the given user. Fear not, a slice of `Posts` will marshal to JSON just as you would expect.

## Assignment

To pass off this step, make more test requests to your api using your favorite REST client. Test test test!
