# Notes

This file is intended to describe the making process of this 
project

# Get external dotenv module

Our first step was to add the following external module using this 
command:

```bash
go get github.com/joho/godotenv
```

With this module we can read environment files `.env` and set the 
variables into 
the current environment.

Usage: 

```go
godotenv.Load()
```

Then we downloaded external modules into our local machine using:

```bash
go mod vendor
```

This will create the vendor file with every source code needed to 
compile the project.

We need to run this later:

```bash
go mod tidy
```

Tidy makes sure go.mod matches the source code in the module.
It adds any missing modules necessary to build the current module's
packages and dependencies, and it removes unused modules that
don't provide any relevant packages. It also adds any missing entries
to go.sum and removes any unnecessary ones.

## Serving an http server

Using `chi` module from `chi-go` and the built in http server module
`net/http` we can create an simple but powerfull http server to handle
requests.

```bash
go get github.com/go-chi/chi/v5/
```

Go code to run an http server:

```go
router := chi.NewRouter()
server := &http.Server{
    Handler: router,
    Addr:    ":" + port,
}
// Start serving an http server
fmt.Printf("Server starting on port %v\n", port)
err := server.ListenAndServe()
```

## Struct tags

A field declaration may be followed by an optional string literal tag, which becomes an attribute for all the fields in the corresponding field declaration. The tags are made visible through a reflection interface but are otherwise ignored.

More on struct tags: 
- https://go.dev/ref/spec#Struct_types
- https://pkg.go.dev/reflect#StructTag

### Tags on json.Marshall()

We use struct tags in this project to indicate the `encoding/json` `Marshall` function some metadata to map certain properties
in some way. For example is used in the `respondWithError` method of [json.go](./json.go) file.

```go
type errResponse struct {
	Error string `json:"error"`
}
```

More examples:

```go
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`
```

More info about struct tags for `Marshall()` [here](https://pkg.go.dev/encoding/json#Marshal).

## Database

We will be using a postgres database in order to store user information, and will install the following go packages:

### sqlc

sqlc generates fully type-safe idiomatic Go code from SQL.sqlc generates fully type-safe idiomatic Go code from SQL.

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### goose

Goose is a database migration tool. Manage your database schema by creating incremental SQL changes or Go functions.

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```