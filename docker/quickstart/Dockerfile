FROM nginx:mainline

ARG TARGETARCH

RUN apt-get update && apt-get upgrade -y \
    && apt-get install -y --no-install-recommends dumb-init ca-certificates \
    && apt-get install -y --no-install-recommends supervisor libstdc++6 \
    && apt-get install -y --no-install-recommends imagemagick \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

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

EXPOSE 80

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]