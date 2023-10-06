# Golang Test Server (with Fiber)

A JSON CRUD API built using the Go programming language and the [Fiber](https://gofiber.io/) web framework.

## About Fiber

Fiber is a web framework for Go that's designed to be fast and flexible. It's inspired by Express, a popular web framework for Node.js, and aims to provide a similar interface for developers while taking advantage of Go's performance benefits. Fiber offers a robust set of features for web development, including middleware support, routing, and templating, while keeping a minimal footprint and optimized performance.

## Project Structure

- `main.go`: The main entry point for the server.
- `src`: Contains additional Go source files or resources.
- `tmp`: Temporary files or resources used by the project.
- `.gitignore`: Specifies patterns that tell Git which files or directories to ignore.
- `Dockerfile`: Used for creating a Docker container for the project.
- `Makefile`: Contains build automation and project management tasks.
- `air.toml`: Configuration file for the [Air](https://github.com/cosmtrek/air) live reload tool for Go.
- `go.mod` and `go.sum`: Files related to Go module dependencies.

## Setup & Running

### Local Development

1. Install Go and set up your Go workspace.
2. Clone the repository.
3. Navigate to the project root and install dependencies using `go mod download`.
4. Build the project using `make build`.
5. Run the server using `make run`.

### Live Reloading (using Air)

For live reloading during development:
1. Ensure you have [Air](https://github.com/cosmtrek/air) installed.
2. Use the following command: `make air`.

### Using Docker

1. Build the Docker image using `make docker-build`.
2. Run the Docker container using `make docker-run`.
