ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /mentor-matcher .


FROM debian:bookworm
# We need root CAs to be installed in order to talk to downstream
# services over HTTPS (like auth0)
RUN apt-get update && apt-get install -y ca-certificates
RUN mkdir -p /app/web
WORKDIR /app
COPY --from=builder /mentor-matcher /usr/local/bin/
COPY web/static /app/web/static
COPY web/template /app/web/template
# fly.io expects us to listen on 0.0.0.0:8080
CMD ["mentor-matcher"]
