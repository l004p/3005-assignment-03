# 3005 Assignment 3 Question 1

This project was created for COMP 3005: Database Management Systems in order to work with connecting to PostgreSQL and implement basic CRUD commands.

It is created with [Golang](https://go.dev), [Cobra](https://github.com/spf13/cobra), and [sqlc](https://github.com/sqlc-dev/sqlc).

Watch the video demo [here](https://youtu.be/6IBSIobI0Og)

## About the folders

**sql:** Where the SQL lives. This includes the database schema, the provided information to populate the database, and the queries I wrote. 

**cmd:** Where the commands live. Root is the root command. The other commands are the respective CRUD operations as outlined by the assignment specs.

**db:** I used sqlc to generate Go code from my hand written SQL queries. Any code generated with the tool lives here.

## Running the code

(This assumes you have Go setup)

1. Create a `.env` file at the root of the directory. Put your own postgres URI into this .env with `DATABASE_URL="your-uri"`

2. Run `go mod tidy` to install dependencies.

3. Run `go run main.go [command] [args]` to compile and run the code in one step, or compile the code seperately with `go build`


