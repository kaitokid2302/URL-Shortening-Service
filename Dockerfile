# alpine lastest go lastes

FROM golang:alpine

# Set the Current Working Directory inside the container

WORKDIR /app

# Copy go mod and sum files

COPY . .

# run go mod tidy

RUN go mod tidy

# Build the Go app

CMD ["go", "run", "./cmd/main.go"]