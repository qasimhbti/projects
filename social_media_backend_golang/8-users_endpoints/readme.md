# Remaining users endpoints

Let's finish the rest of the endpoints for dealing with users, we've done most of the hard work already!

### handlerDeleteUser

`func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request)`

This handler should be invoked by the client sending a `DELETE` request to the `/users/EMAIL` endpoint. Notice the `EMAIL` variable in the path. Conventionally `DELETE` resuests don't have a body, so parameters are often passed in the path instead. If a user want's to delete `test@example.com`'s user record, they would send a `DELETE` request to `/users/test@example.com`.

The full path of the request, including the `/users/` part will be in the `r.URL.Path` variable in your handler. You can use the [strings.TrimPrefix](https://pkg.go.dev/strings#TrimPrefix) function to remove the prefix and get access to the email in the path.

After that, use the parameters to call `apiCfg.dbClient.DeleteUser()`. If all goes well, respond with `http.StatusOK` (200), and an empty JSON body `{}`.

```go
respondWithJSON(w, http.StatusOK, struct{}{})
```

### handlerGetUser

`func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request)`

This handler should be invoked by the client sending a `GET` request to the `/users/EMAIL` endpoint, similar to the delete handler.

Once you get the email from the path in the same way, call `apiCfg.dbClient.GetUser()`. If all goes well, respond with `http.StatusOK` (200), and the marshalled JSON of the user record.

### handlerUpdateUser

`func (apiCfg apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request)`

This one will also take the `email` variable via the path parameter, but it will also take the other parameters via the body. This is a common pattern that signals to the client that the `email` is the primary key (indicates which record is changing), and the stuff in the body is the data that will be updated on the record.

```go
type parameters struct {
    Password string `json:"password"`
    Name     string `json:"name"`
    Age      int    `json:"age"`
}
```

Once you've gathered all the input, call `apiCfg.dbClient.UpdateUser`. If all goes well, return a `http.StatusOK` (200) and the updated user record.

## Assignment

To pass off this step, make all the various requests to your API using your favorite rest client and make sure it does what you expect. Test test test!
