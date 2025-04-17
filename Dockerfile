FROM golang:1.24-alpine AS build
WORKDIR /build/binvault
COPY . .
RUN go mod download

RUN apk add --no-cache gcc g++ make musl-dev sqlite-dev
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o binvault-server ./pkg/main.go
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o binvault ./cmd/main.go

FROM nginx:stable-alpine

RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && apk add --no-cache supervisor libstdc++ \
    && apk add --no-cache imagemagick

WORKDIR /service

RUN mkdir data && mkdir data/public

COPY nginx/www ./data/public
COPY nginx/conf.d/* /etc/nginx/conf.d/
COPY nginx/nginx.conf /etc/nginx/nginx.conf

COPY nginx/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

COPY nginx/processors.cfg .

COPY --from=build /build/binvault/binvault /usr/local/bin
COPY --from=build /build/binvault/binvault-server .

RUN chmod +x /service/binvault-server

EXPOSE 8080
EXPOSE 80

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]