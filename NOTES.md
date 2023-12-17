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