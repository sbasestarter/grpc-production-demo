# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# We specify the base image we need for our go application
# https://hub.docker.com/_/golang?tab=tags
FROM golang:1.14.15-alpine3.11 AS build

ARG GITHUB_ACCESS_TOKEN

# ENV GOPROXY https://mirrors.aliyun.com/goproxy/
# Set the Current Working Directory inside the container
WORKDIR /go/release

# Copy go mod and sum files
COPY go.mod go.sum ./

ENV GOPROXY=https://goproxy.io,direct

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

ENV GOPRIVATE=github.com/sbasestarter/grpc-production-demo

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# govvv
RUN apk update && apk add git

# git auth
RUN git config --global url."https://${GITHUB_ACCESS_TOKEN}@github.com".insteadOf "https://github.com"

# RUN go clean -modcache
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o bserver backserver/cmd/server.go

# Multi-stage build - final stage
FROM alpine:3.13.4

RUN apk add --no-cache bash && addgroup -S backserver && adduser -S backserver -G backserver -u 1000
USER backserver
ENV GOROOT=/home/backserver/go

COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /home/demo/go/lib/time/zoneinfo.zip
COPY --from=build /go/release/bserver /home/backserver

ENTRYPOINT ["/home/backserver/bserver"]
