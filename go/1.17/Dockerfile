FROM wbuntu/alpine:3.15
ARG GOVERSION
ARG TARGETARCH
ENV GOROOT=/root/go/root
ENV GOPATH=/root/go/path
ENV PATH="${GOROOT}/bin:${GOPATH}/bin:${PATH}"
RUN mkdir -p /root/go /root/work && curl -sL https://go.dev/dl/go$GOVERSION.linux-$TARGETARCH.tar.gz | tar -zx && mv go $GOROOT
WORKDIR /root/work