# images

自定义镜像合集，整合容器环境使用的镜像和构建脚本

## 1. 镜像目录

### 1.1 系统镜像

- [alpine](alpine)
- [debian](debian)

### 1.2 开发环境

- [go](go)

### 1.3 代理工具

- [gost](gost)
- [hysteria](hysteria)
- [mtproxy](mtproxy)
- [shadowray](shadowray)
- [tuic](tuic)

## 2. 备忘录

### 2.1 Makefile

- make image: 默认使用docker和本机架构构建镜像
- make release: 默认使用buildkit构建镜像，默认构建linux/amd64与linux/arm64两种架构（若应用支持）

### 2.2 Docker

文档链接：[Dockerfile reference](https://docs.docker.com/engine/reference/builder/)

**创建buildx构建实例**

```	shell
docker buildx create --use --name multiarch --platform linux/amd64,linux/arm64
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

