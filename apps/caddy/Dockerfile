ARG VERSION
FROM --platform=$BUILDPLATFORM caddy:${VERSION}-builder AS builder

ARG VERSION
ARG TARGETARCH
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=${TARGETARCH}
ENV CADDY_VERSION=v${VERSION}
RUN xcaddy build ${CADDY_VERSION} \
--with github.com/mholt/caddy-webdav

FROM --platform=$TARGETPLATFORM caddy:${VERSION}
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
