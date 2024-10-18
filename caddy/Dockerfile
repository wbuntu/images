ARG TARGETARCH
ARG VERSION

FROM --platform=$BUILDPLATFORM caddy:${VERSION}-builder AS builder
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} xcaddy build ${VERSION} \
--with github.com/mholt/caddy-webdav

FROM --platform=$TARGETPLATFORM caddy:${VERSION}
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
