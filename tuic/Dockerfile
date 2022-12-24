FROM wbuntu/alpine:3.15
ARG VERSION
ARG TARGETARCH
COPY download.sh /usr/bin/download.sh
RUN /usr/bin/download.sh $VERSION $TARGETARCH
CMD ["/usr/bin/tuic-server","-c","/etc/tuic/config.json"]
