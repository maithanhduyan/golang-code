# Golang CRUD with Postgresql

####
- Create go module
> go mod init exmaple.com/go-crud-postgresql

#### Installing Dependencies
- gorilla/mux router: Implements a request router and dispatcher for matching incoming requests to their respective handler.
> go get -u github.com/gorilla/mux

- lib/pq driver: A pure Go PostgreSQL driver for Go’s database/sql package.
> go get github.com/lib/pq

- joho/godotenv: We will use the godotenv package to read the .env file, which is used to save environment variables for keeping sensitive data safe.
> go get github.com/joho/godotenv

### Database
~~~~
CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    name TEXT,
    age INT,
    location TEXT
);
~~~~




