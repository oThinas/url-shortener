
# URL Shortener

A simple and efficient URL shortener service built with Go. This project allows users to shorten long URLs and redirect to the original URLs using a generated short code.

## Features

- Shorten any valid URL
- Fast redirection using short codes
- Simple RESTful API
- Persistent storage using Redis (via Docker)

## Tech Stack

- Go (Golang)
- Docker & Docker Compose
- Redis
- [Chi](https://go-chi.io/#/)

## Project Structure

```plaintext
├── cmd       # Entry points for running the application (main packages)
│   └── api
├── internal  # Application code (API handlers, business logic, storage)
│   ├── api   # HTTP handlers and API logic
└─  └── store # Storage logic and Redis integration
```

## API Reference

#### Shorten a long URL

```http
  POST /api/shorten
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `url` | `string` | **Required**. The original URL to shorten |

#### Get the original URL using the short code

```http
  GET /api/${code}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `code` | `string` | **Required**. The short code generated for the URL |

## Run Locally

1. **Clone the repository:**

   ```fish
   git clone https://github.com/oThinas/url-shortener
   cd url-shortener
   ```

2. **Run with Docker Compose:**

   ```fish
   docker compose up --build
   ```

3. **Run the application:**

   ```fish
   go run ./cmd/api/main.go
   ```
