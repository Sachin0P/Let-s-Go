# Snippetbox

A web application for creating, viewing, and sharing text snippets — built in Go.

This project follows along with **[Let's Go](https://lets-go.alexedwards.net/)** by [Alex Edwards](https://www.alexedwards.net/), a book that teaches how to build production-grade web applications with Go. It's a learning project, not a production app, and serves as a personal record of working through the book.

## Features

- Browse the most recently created snippets
- View an individual snippet by ID
- Create new snippets with a configurable expiry (in days)
- Snippets automatically "expire" and are excluded from listings once their expiry date passes

## Tech Stack

- [Go](https://go.dev/) (standard `net/http`, `html/template`)
- [MySQL](https://www.mysql.com/) for storage, via [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- Plain HTML/CSS/JS for the frontend (no frontend framework)

## Project Structure

```
.
├── cmd/web/              # Application entry point and HTTP handlers
│   ├── main.go           # Wiring: config flags, DB connection, server startup
│   ├── routes.go         # Route definitions
│   ├── handlers.go       # HTTP handler functions
│   ├── helpers.go        # Shared helper functions
│   └── templates.go      # HTML template rendering
├── internal/models/      # Data access layer
│   ├── snippets.go       # SnippetModel: Insert, Get, Latest
│   └── errors.go         # Shared model errors
└── ui/
    ├── html/             # Page and partial templates
    └── static/           # CSS, JS, images
```

## Getting Started

### Prerequisites

- Go (see [go.mod](go.mod) for the required version)
- A running MySQL instance

### Database Setup

Create a database and a `snippets` table, for example:

```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4;

USE snippetbox;

CREATE TABLE snippets (
    id      INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title   VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);
```

Create a MySQL user with the appropriate privileges on this database, and update the DSN accordingly when running the app.

### Running the App

```bash
go run ./cmd/web -dsn="web:<password>@/snippetbox?parseTime=true" -addr=":4000"
```

Flags:

| Flag   | Default                                            | Description                  |
|--------|-----------------------------------------------------|-------------------------------|
| `-addr` | `:4000`                                             | HTTP network address          |
| `-dsn`  | `web:2401@/snippetbox?parseTime=true`               | MySQL data source name        |

Then visit [http://localhost:4000](http://localhost:4000) in your browser.

## Routes

| Method | Pattern                  | Handler            | Description                  |
|--------|---------------------------|---------------------|-------------------------------|
| GET    | `/`                        | `home`              | Show latest snippets          |
| GET    | `/snippet/view/{id}`       | `snippetView`       | View a specific snippet       |
| GET    | `/snippet/create`          | `snippetCreate`     | Show the snippet creation form|
| POST   | `/snippet/create`          | `snippetCreatePost` | Create a new snippet          |
| GET    | `/static/...`               | —                   | Serve static assets           |

## Acknowledgements

This project is built while reading and following **[Let's Go: Learn to Build Professional Web Applications with Go](https://lets-go.alexedwards.net/)** by Alex Edwards. All credit for the architecture and teaching approach goes to the book — this repo is a personal implementation for learning purposes.
