FROM --platform=$BUILDPLATFORM alpine:3.15
ARG VERSION
ARG TARGETARCH
RUN wget -q https://github.com/HyNetwork/hysteria/releases/download/$VERSION/hysteria-linux-$TARGETARCH -O /usr/bin/hysteria && chmod a+x /usr/bin/hysteria
CMD ["/usr/bin/hysteria","-c","/etc/hysteria/config.json","server"]
