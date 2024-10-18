FROM wbuntu/ollama:v0.3.10-cpu AS ollama
FROM caddy:2.8.4 AS caddy
FROM debian:12
ENV TZ=Asia/Shanghai
RUN apt-get update && apt-get install -y tree ca-certificates supervisor && apt-get clean
COPY --from=ollama /usr/bin/ollama /usr/bin/ollama
COPY --from=caddy /usr/bin/caddy /usr/bin/caddy
COPY Caddyfile /backup/Caddyfile
COPY supervisord.conf /etc/supervisor/supervisord.conf
COPY supervisor.d /etc/supervisor.d
COPY entrypoint.sh /entrypoint.sh
EXPOSE 8080
VOLUME ["/data"]
ENTRYPOINT [ "/entrypoint.sh" ]