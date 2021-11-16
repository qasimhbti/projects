# Unit tests

Unit tests are an important part of any production app, and sanitizing user input even more so!

### userIsEligible

Let's write a function that takes some new user data, and returns an error if anything seems amiss.

`func userIsEligible(email, password string, age int) error`

* If the email is empty, return `errors.New("email can't be empty")`
* If the password is empty, `password can't be empty`
* If the age is less than 18, return `age must be at least AGE years old`, where `AGE` is the actual age.

The [fmt.Errorf](https://pkg.go.dev/fmt#Errorf) function will be useful for that last one.

If nothing is wrong, return `nil`.

### TestUserIsEligible

Now let's write our first unit test! Testing in Go is build in to the toolchain, and started by running `go test` in the package you want to test. Test files in go end in `_test.go`, so our file might be `user_is_eligible.go`. The actual test function should start with the word `Test...` and the function must take a `*testing.T` type as input.

For example: `func TestUserIsEligible(t *testing.T)`

I like to build my tests using a table design. In other words, I specify the test data in a slice of structs, then use a loop to test each case. My test data will look something like this.

```go
var tests = []struct {
    email       string
    password    string
    age         int
    expectedErr error
}{
    {
        email:       "test@example.com",
        password:    "12345",
        age:         18,
        expectedErr: nil,
    },
    {
        email:       "",
        password:    "12345",
        age:         18,
        expectedErr: errors.New("email can't be empty"),
    },
    {
        email:       "test@example.com",
        password:    "",
        age:         18,
        expectedErr: errors.New("password can't be empty"),
    },
    {
        email:       "test@example.com",
        password:    "12345",
        age:         16,
        expectedErr: errors.New("age must be at least 18 years old"),
    },
}
```

Next, I'll loop over each test.

```go
for _, tt := range tests {
    err := userIsEligible(tt.email, tt.password, tt.age)
    errString := ""
    expectedErrString := ""
    if err != nil {
        errString = err.Error()
    }
    if tt.expectedErr != nil {
        expectedErrString = tt.expectedErr.Error()
    }
    if errString != expectedErrString {
        t.Errorf("got %s, want %s", errString, expectedErrString)
    }
}
```

If `e.Errorf()` is ever called, `go test` will fail and print the error.

## Assignment

To pass off this step, get your unit tests to pass! Try temporarily breaking them on purpose as well to make sure it's all working as intended.
