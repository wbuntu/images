FROM alpine:3.15
ARG VERSION
ARG TARGETARCH
RUN apk add --no-cache wget
RUN wget -q https://github.com/HyNetwork/hysteria/releases/download/$VERSION/hysteria-linux-$TARGETARCH -O /usr/bin/hysteria
RUN chmod a+x /usr/bin/hysteria
CMD ["/usr/bin/hysteria","-c","/etc/hysteria/config.json","server"]
