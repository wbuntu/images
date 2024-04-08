FROM yidadaa/chatgpt-next-web:v2.11.3 as chat
FROM ghcr.io/aurora-develop/aurora:latest as aurora
FROM alpine:3.18
ARG TARGETARCH
RUN apk add --no-cache supervisor curl nodejs 
COPY --from=aurora /app /app
COPY --from=chat /app /app/chatgpt-next-web
COPY supervisord.conf /etc/supervisord.conf
COPY supervisor.d /etc/supervisor.d
CMD ["/usr/bin/supervisord","-c","/etc/supervisord.conf"]
