FROM alpine:latest

ARG TARGETARCH

WORKDIR /service

COPY ./binvault-${TARGETARCH} /usr/local/bin/binvault
COPY ./binvault-server-${TARGETARCH} binvault-server

RUN chmod +x /usr/local/bin/binvault binvault-server

CMD ["binvault-server"]