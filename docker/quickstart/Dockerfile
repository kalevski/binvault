FROM nginx:stable-alpine

ARG TARGETARCH

RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && apk add --no-cache supervisor libstdc++ \
    && apk add --no-cache imagemagick

WORKDIR /service

RUN mkdir data && mkdir data/public

COPY ./www ./data/public
COPY ./conf.d/* /etc/nginx/conf.d/
COPY ./nginx.conf /etc/nginx/nginx.conf

COPY ./supervisord.conf /etc/supervisor/conf.d/supervisord.conf

COPY ./processors.cfg .

COPY ./binvault-${TARGETARCH} /usr/local/bin/binvault
COPY ./binvault-server-${TARGETARCH} binvault-server

RUN chmod +x /service/binvault-server

EXPOSE 8080
EXPOSE 80

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]