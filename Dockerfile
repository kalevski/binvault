FROM golang:1.19 AS build
WORKDIR /go/src/binvault
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .
FROM alpine:latest as release
WORKDIR /app
COPY --from=build /go/src/binvault/app .
RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && chmod +x /app/app

EXPOSE 3000

ENTRYPOINT ["/usr/bin/dumb-init", "--"]