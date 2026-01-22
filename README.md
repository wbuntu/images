# images

自定义镜像合集，整合容器环境使用的镜像和构建脚本

## 1. 镜像目录

### 1.1 系统

- [alpine](os/alpine)
- [debian](os/debian)

### 1.2 开发环境

- [golang](dev/golang)
- [node](dev/node)

### 1.3 代理工具

- [gost](proxy/gost)
- [shadowray](proxy/shadowray)
- [mtproxy](proxy/mtproxy)

### 1.4 AI

- [next-chat-free](ai/next-chat-free)
- [free-ask-internet](ai/free-ask-internet)
- [next-chat-cf](ai/next-chat-cf)
- [ollama-registry](ai/ollama-registry)

### 1.5 通用应用

- [caddy](apps/caddy)

### 1.6 归档

- [tuic](archive/tuic)
- [hysteria](archive/hysteria)
- [n8n](archive/n8n)

## 2. 备忘录

### 2.1 Makefile

- make image: 默认使用docker和本机架构构建镜像
- make release: 默认使用buildkit构建镜像，默认构建linux/amd64与linux/arm64两种架构（若应用支持）
- make all: 构建多个release版本

### 2.2 Docker

文档链接：[Dockerfile reference](https://docs.docker.com/build/building/multi-platform/)

**创建buildx构建实例**

```	shell
docker buildx create --use --name multiarch --platform linux/amd64,linux/arm64
```

**安装QEMU**

```shell
docker run --privileged --rm tonistiigi/binfmt --install all
```

**多架构变量**

- TARGETPLATFORM: 目标镜像的平台，例如linux/amd64，linux/arm/v7， windows/amd64
- TARGETOS: TARGETPLATFORM 的操作系统部分，如linux，windows
- TARGETARCH: TARGETPLATFORM 的架构部分，如amd64，arm64
- TARGETVARIANT: TARGETPLATFORM 的变体组件，如v7
- BUILDPLATFORM: 执行构建的节点平台
- BUILDOS: BUILDPLATFORM 的操作系统组件
- BUILDARCH: BUILDPLATFORM 的架构组件
- BUILDVARIANT: BUILDPLATFORM 的变体组件

### 2.3 交叉编译

buildkitd使用qemu来模拟目标架构执行多架构编译，但是运行效率很低，若编译工具支持交叉编译，可以使用本机架构完成编译，然后将编译出的二进制文件拷贝到目标架构的基础镜像中制作镜像，以go语言为例，如下：

```dockerfile
# buildkit跨架构编译缓慢，统一使用本机架构进行交叉编译
FROM --platform=$BUILDPLATFORM wbuntu/golang:1.19 AS builder
ARG TARGETARCH
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=$TARGETARCH
WORKDIR /custom-service
COPY . /custom-service
RUN make build
# 编译完成后拷贝到目标架构的基础镜像中
FROM --platform=$TARGETPLATFORM wbuntu/alpine:3.15
COPY --from=builder /custom-service/custom-service /usr/bin/custom-service
CMD ["/usr/bin/custom-service","-c","/etc/custom-service/config.toml"]
```

