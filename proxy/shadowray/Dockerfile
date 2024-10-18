FROM debian:11
ENV TZ=Asia/Shanghai
ARG TARGETARCH
RUN apt update && apt install shadowsocks-libev curl -y && apt clean && curl -L https://github.com/shadowsocks/v2ray-plugin/releases/download/v1.3.2/v2ray-plugin-linux-$TARGETARCH-v1.3.2.tar.gz | tar -zx \
    && mv v2ray-plugin_linux_$TARGETARCH /usr/bin/v2ray-plugin
CMD ["/usr/bin/ss-server","-c","/etc/shadowsocks-libev/config.json"]
