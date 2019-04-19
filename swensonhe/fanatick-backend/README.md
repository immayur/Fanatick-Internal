# Fanatick Backend

Fanatick Backend is the backend server application for Fanatick.

## Installation

### Go

Go is the language of the application.

Installation: [https://golang.org/doc/install](https://golang.org/doc/install)

### PostgreSQL

PostgreSQL is the database for the application.

```bash
brew install postgresql
```

### Migrate

Migrate is used to run database migrations. 

```bash
brew install golang-migrate
```

## Getting Started

Create a local database:
```bash
createdb fanatick
```

Run the database migrations:
```bash
make migrate DB_URL="psql://localhost/fanatick?sslmode=disable"
```

## Run

Export the environment variables `.envrc-template` using your method of choice. Your IDE is a good option.

Generate swagger specs:
```bash
    swag init -o cmd/http/openapi/
```

Build and run it:   
```bash
go build -o ./bin/http -i ./cmd/http
 ./bin/http
```

Lookout for specs on:
```bash
    yourhost + /swagger/index.html
    
```

## Test

Run tests:
```bash
make test DB_NAME=fanatick_test
```

## Architecture

```
fanatick-backend
├── bin            # application binaries
├── cmd            # application commands
|   └── http       # server application
└── fanatick       # Fanatick domain definitions (models, interfaces)
    ├── algolia    # Algolia client (search)
    ├── api        # Fanatick API implementation
    ├── firebase   # Firebase client (auth)
    ├── mock       # mock interface implementations
    ├── postgres   # PostgreSQL client (DB)
    └── pushwoosh  # Pushwoosh client (push notifications)
```
