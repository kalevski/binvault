FROM golang:1.24-alpine AS build
WORKDIR /build/binvault
COPY . .
RUN go mod download

RUN apk add --no-cache gcc g++ make musl-dev sqlite-dev
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o binvault-server ./pkg/main.go
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o binvault ./cmd/main.go

ARG GUETZLI_VERSION=1.0.1
ARG GUETZLI_CHECKSUM_SHA256=e52eb417a5c0fb5a3b08a858c8d10fa797627ada5373e203c196162d6a313697

RUN apk add --update --no-cache --virtual .build \
    gflags \
    libpng-dev \
    curl \
    make \
    g++ \
    && apk add --update --no-cache libpng libstdc++ libgcc \
    && curl -L --fail -o /tmp/guetzli.tar.gz \
        https://github.com/google/guetzli/archive/refs/tags/v${GUETZLI_VERSION}.tar.gz \
    && echo "${GUETZLI_CHECKSUM_SHA256}  /tmp/guetzli.tar.gz" > /tmp/guetzli_checksum.txt \
    && sha256sum -w -c /tmp/guetzli_checksum.txt \
    && mkdir /tmp/guetzli \
    && tar -zxv --strip-components=1 -f /tmp/guetzli.tar.gz -C /tmp/guetzli \
    && cd /tmp/guetzli \
    && make \
    && cp bin/Release/guetzli /build/guetzli \
    && rm -f /tmp/guetzli.tar.gz \
    && rm -rf /tmp/guetzli \
    && rm -f /tmp/guetzli_checksum.txt \
    && apk del .build

FROM nginx:stable-alpine

RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && apk add --no-cache pngquant libstdc++ \
    && apk add --no-cache supervisor

WORKDIR /service

RUN mkdir data && mkdir data/public

COPY nginx/www ./data/public
COPY nginx/conf.d/* /etc/nginx/conf.d/
COPY nginx/nginx.conf /etc/nginx/nginx.conf

COPY nginx/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

COPY --from=build /build/binvault/binvault /usr/local/bin
COPY --from=build /build/guetzli /usr/local/bin
COPY --from=build /build/binvault/binvault-server .


RUN chmod +x /service/binvault-server

EXPOSE 8080
EXPOSE 80

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]