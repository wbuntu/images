FROM --platform=$TARGETPLATFORM golang:1.18-alpine as builder
ARG VERSION
ENV CGO_ENABLED=0
RUN apk add --no-cache git
RUN git clone https://github.com/ginuerzh/gost.git /gost 
RUN cd /gost && git checkout $VERSION && cd cmd/gost && go env && go build -v
FROM --platform=$TARGETPLATFORM alpine:3.15
COPY --from=builder /gost/cmd/gost/gost /usr/bin/gost
CMD ["/usr/bin/gost","-C","/etc/gost/config.json"]
