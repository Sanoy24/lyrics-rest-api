# Gemius Lyrics Clone REST API

[![Go Report Card](https://goreportcard.com/badge/github.com/Sanoy24/lyrics-rest-api)](https://goreportcard.com/report/github.com/Sanoy24/lyrics-rest-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A robust and scalable RESTful API for managing song lyrics, built with Go. This project demonstrates a modern backend architecture, clean code principles, and a strong focus on security and performance. It serves as a backend foundation for any music or lyrics-based application.

---

## ✨ Features

-   **Full User Management**: Secure user registration, login, and profile management.
-   **JWT-based Authentication**: Stateless and secure authentication using JSON Web Tokens.
-   **Role-Based Access Control (RBAC)**: Differentiated user roles (e.g., `user`, `admin`) to protect endpoints.
-   **CRUD Operations for Lyrics**: Full Create, Read, Update, and Delete functionality for song lyrics.
-   **RESTful API Design**: Clean, intuitive, and well-structured API endpoints.
-   **Structured Logging**: Detailed and structured logging with `slog` for easier debugging and monitoring.
-   **Configuration Management**: Centralized configuration handling for easy setup across different environments.
-   **Graceful Shutdown**: Ensures all active requests are completed before the server shuts down, preventing data corruption.

## 🛠️ Tech Stack & Architecture

This project is built with a modern Go stack, emphasizing best practices and scalability.

### Technologies Used

-   **Language**: **Go**
-   **Database**: **PostgreSQL**
-   **ORM**: **GORM** for elegant and efficient database interactions.
-   **Routing**: **Gin Gonic** (or your choice of router like `chi` or `net/http`) for high-performance HTTP routing.
-   **Logging**: `log/slog` (Structured Logging from Go's standard library).
-   **Authentication**: `golang-jwt` for JWT implementation.
-   **Environment Variables**: `godotenv` for easy management of configuration.
-   **Containerization**: **Docker** & **Docker Compose** for consistent development and deployment environments.

### Project Architecture

The API follows a **Layered Architecture** (similar to Clean or Hexagonal Architecture) to ensure a clear separation of concerns, making the codebase modular, scalable, and easy to maintain.

```
lyrics-rest-api/
├── api/                # API handlers, routes, and middleware
├── cmd/                # Main application entrypoint
├── config/             # Configuration management
├── internal/
│   ├── models/         # Database models (User, Role, Lyric, etc.)
│   ├── services/       # Business logic layer
│   └── repositories/   # Data access layer (interacts with the DB)
├── pkg/                # Shared packages (e.g., apperror, utils)
└── ...
```

-   **Repository Pattern**: The data access logic is abstracted away from the business logic using repositories (as seen in `internal/api/repositories/user/user_repo.go`). This makes it easy to switch database technologies or mock the database for testing.
-   **Service Layer**: The core business logic resides in the service layer, orchestrating data between handlers and repositories.
-   **Handler/Controller Layer**: This layer is responsible for handling HTTP requests, validating input, and returning appropriate HTTP responses.

---

## 🚀 Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

-   Go (version 1.21 or newer)
-   Docker and Docker Compose
-   Make (optional, for using Makefile commands)

### Installation & Running the App

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/Sanoy24/lyrics-rest-api.git
    cd lyrics-rest-api
    ```

2.  **Set up environment variables:**
    Create a `.env` file in the root directory by copying the example file.

    ***

    ## 📚 Swagger / API Documentation

    This project includes Swagger documentation generated with swag (swaggo) and served via `gin-swagger`.

    Where to access

    -   Swagger UI: http://localhost:8080/swagger/index.html
    -   The router mounts the Swagger UI at `/swagger/*any` (see `internal/api/router/router.go`).

    Generating or updating the docs

    1. Install the `swag` CLI (if you haven't already):

    ```sh
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

    2. From the project root, generate the documentation (this writes into `./docs`):

    ```sh
    swag init -g ./cmd/server/main.go -o ./docs
    ```

    Serving the docs

    1. Start the API locally:

    ```sh
    go run ./cmd/server/main.go
    ```

    2. Or start with Docker Compose:

    ```sh
    docker-compose up --build
    ```

    3. Open the Swagger UI in your browser:

    ```
    http://localhost:8080/swagger/index.html
    ```

    Notes & troubleshooting

    -   The Swagger metadata (title, version, host, BasePath) is defined by annotations in `cmd/server/main.go` and by the generated `docs` package (`docs/docs.go`). If you change annotations, re-run `swag init`.
    -   If running inside Docker, ensure the port mapping in `docker-compose.yml` exposes the server port to the host.
    -   If the UI loads but shows no endpoints, regenerate the docs and restart the server.

    If you'd like, I can also add a Make target or npm script to simplify generating the docs (e.g., `make docs`).

    ```sh
    cp .env.example .env
    ```

    Now, open the `.env` file and fill in your database credentials and other required values.

3.  **Run locally :**
    If you have a local PostgreSQL instance running, you can run the app directly.

    ```sh
    # Install dependencies
    go mod tidy

    # Run the application
    go run ./cmd/api/main.go
    ```

---

## ⚙️ Environment Variables

The following environment variables are needed to run the application:

```env
# Server Configuration
HTTP_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=lyrics_db

# JWT Configuration
JWT_SECRET_KEY=your_super_secret_key
JWT_EXPIRATION_HOURS=72
```
