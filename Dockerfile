FROM golang:alpine

# Set env variables

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build


# Copy code into container
COPY . .
COPY internal ./

# Copy and download dependencies with go mod
COPY go.mod .
COPY go.sum .
RUN go mod download


# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Command to run when starting the container
CMD ["/dist/main"]
