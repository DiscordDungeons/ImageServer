FROM golang:alpine as builder

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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

##### Deployment Image #####

FROM scratch

# Gotta copy the SSL certs, or no https
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# Move to /bin directory as the place for resulting binary folder
WORKDIR /bin

COPY --from=builder /build/app .

# Command to run when starting the container
CMD ["./app"]
