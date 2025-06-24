ARG VERSION
FROM alpine:$VERSION
ENV TZ=Asia/Shanghai
RUN apk add --no-cache busybox-extras bash curl wget git make htop tzdata
COPY index.html /var/www/index.html
WORKDIR /root
CMD [ "bash" ]