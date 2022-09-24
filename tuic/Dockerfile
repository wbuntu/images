FROM alpine:3.15
ENV TZ=Asia/Shanghai
ARG VERSION
ARG TARGETARCH
RUN apk add --no-cache wget
COPY download.sh /usr/bin/download.sh
RUN /usr/bin/download.sh $VERSION $TARGETARCH
CMD ["/usr/bin/tuic-server","-c","/etc/tuic/config.json"]
