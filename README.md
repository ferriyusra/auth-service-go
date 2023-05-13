# Auth Service

This microservice is the API service for Auth Service.

## Documentation

[Product Requirements Document]()

[Design Flow]()

## Requirements

This project is developed with:

- Go 1.17

- Postgres 14

## Installation

Clone the project

```bash
git clone https://github.com/ferriyusra/auth-service-go
```

Go to the project directory

```bash
cd auth-service
```

This service contains a `.env.example` file that defines environment variables you need to set. Copy and set the variables to a new `.env` file.

```bash
cp .env.example .env
```

Start the app

```bash
go run main.go
nodemon --exec go run main.go --signal SIGTERM
```

## Database

If you have not created the database for auth service, please create one before going to the next step.

We're using [golang-migrate](https://github.com/golang-migrate/migrate) for the migration.

### Without Docker

Install the package

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Run the migration (change the value accordingly)

```bash
migrate -path=migrations -database "postgres://postgres:password@127.0.0.1:5432/database?sslmode=disable" up
```

To rollback

```bash
migrate -path=migrations -database "postgres://postgres:password@127.0.0.1:5432/database?sslmode=disable" down 1
```

### With Docker

Run the migration (change the value accordingly)

```bash
docker run -v "$(pwd)"/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://postgres:password@127.0.0.1:5432/database?sslmode=disable" up
```

To rollback

```bash
docker run -v "$(pwd)"/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://postgres:password@127.0.0.1:5432/database?sslmode=disable" down 1
```

## Deployment

### Without Docker

Build the binary

```bash
go build -o auth
```

Run it

```bash
./auth
```

### With Docker

Build the image

```bash
docker build -t auth-service .
```

Run the container

```bash
docker run -d -p 8080:8080 --name auth1 --network=host --env-file=.env auth-service
```

## Testing

Run testing with coverage

```bash
go test -coverprofile=coverage.out ./...
```

Show coverage detail

```bash
go tool cover -func=coverage.out
```

Show coverage detail as HTML

```bash
go tool cover -html=coverage.out
```
