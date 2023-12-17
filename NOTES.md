# Notes

This file is intended to describe the making process of this project

# Get external dotenv module

Our first step was to add the following external module using this command:

```bash
go get github.com/joho/godotenv
```

With this module we can read environment files `.env` and set the variables into the current environment.

Usage: 

```go
godotenv.Load()
```

Then we downloaded external modules into our local machine using:

```bash
go mod vendor
```

This will create the vendor file with every source code needed to compile the project.