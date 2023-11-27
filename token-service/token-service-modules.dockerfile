# Step 1: Modules caching
FROM golang:alpine as token-service-modules
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download