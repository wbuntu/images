ARG VERSION
FROM golang:${VERSION}-alpine
ENV TZ=Asia/Shanghai
RUN apk add --no-cache busybox-extras bash curl wget git make htop tzdata
WORKDIR /root
CMD [ "bash" ]