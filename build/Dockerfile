FROM golang:1.18-alpine as builder

WORKDIR /app

# Fetch dependencies.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary.
RUN cd cmd/data-structure-algoritms \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/data-structure-algoritms .

FROM alpine:3.16

ARG ROOT_FOLDER="app"
ARG BUILD_ID=0
ARG VERSION=0.0.1

ENV ROOT_FOLDER=${ROOT_FOLDER} \
    BUILD_ID=${BUILD_ID} \
    VERSION=${VERSION}

# Copy static executable
COPY --from=builder /go/bin/data-structure-algoritms app/data-structure-algoritms

EXPOSE 90

# Run binary.
ENTRYPOINT ["app/data-structure-algoritms"]
