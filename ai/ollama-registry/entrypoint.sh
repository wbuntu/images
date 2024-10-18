#!/bin/bash

mkdir -p /data/ollama/models /data/caddy
ln -sfv /data/ollama /root/.ollama

if [ ! -f /data/caddy/Caddyfile ]; then
    cp /backup/Caddyfile /data/caddy/Caddyfile
fi

/usr/bin/supervisord -c /etc/supervisor/supervisord.conf