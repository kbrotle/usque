FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# -trimpath helps with reproducible builds and removes local path info from binaries
RUN go build -o usque -ldflags="-s -w" -trimpath .

# scratch won't be enough, because we need a cert store
FROM alpine:latest

# Keep the image up to date with latest security patches
RUN apk --no-cache upgrade

WORKDIR /app

COPY --from=builder /app/usque /bin/usque

ENTRYPOINT ["/bin/usque"]
