.PHONY: build run clean docker-build docker-run

# Run air for live reloading
air:
	@air

# Build the Go project
build:
	@go build -o ./tmp/main .

# Run the Go project
run:
	@./tmp/main

# Remove build artifacts
clean:
	@rm -rf ./tmp/main

# Build a Docker image
docker-build:
	@docker build -t golang-server .

# Run the Docker image
docker-run:
	@docker run -d -p 3000:3000 --name go-server golang-server
