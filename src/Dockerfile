FROM golang:1.17.3-alpine3.15 AS base
WORKDIR /src

# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o /app

# Now create separate deployment image
FROM gcr.io/distroless/base

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

WORKDIR /src
COPY --from=0 /app .
ENTRYPOINT ["./app"]
