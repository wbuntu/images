ARG VERSION
FROM debian:$VERSION
ENV TZ=Asia/Shanghai
RUN apt update && apt install -y curl wget git make htop
WORKDIR /root
CMD [ "bash" ]