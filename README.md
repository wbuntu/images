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

#### 2.2.1 创建 buildx 构建实例

```shell
docker buildx create --use --name multiarch --platform linux/amd64,linux/arm64
```

#### 2.2.2 安装 QEMU

```shell
docker run --privileged --rm tonistiigi/binfmt --install all
```

#### 2.2.3 多架构变量

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

### 2.4 GitHub Actions 自动构建

项目已配置 GitHub Actions 自动构建镜像，在创建 Git Tag 时自动触发。

#### 2.4.1 配置步骤

1. 在 GitHub 仓库设置中添加以下 Secrets：

   - `DOCKERHUB_USERNAME`: Docker Hub 用户名
   - `DOCKERHUB_TOKEN`: Docker Hub 访问令牌（在 [Docker Hub Settings](https://hub.docker.com/settings/security) 生成）

2. 创建并推送 Tag 触发构建：

   ```shell
   # 创建版本标签
   git tag v1.0.0

   # 推送标签到远程仓库
   git push origin v1.0.0
   ```

3. 工作流会自动：

   - 按顺序构建：os → dev → proxy/apps
   - 读取各个镜像目录 Makefile 中的 `TARGETS` 变量
   - 为每个版本构建 linux/amd64 和 linux/arm64 多架构镜像
   - 推送到 Docker Hub

   **注意**：不会构建 `ai` 和 `archive` 目录下的镜像

#### 2.4.2 示例

创建 Tag `v1.0.0` 后，会按以下顺序构建镜像：

**第一步（系统镜像）**：

- `wbuntu/alpine:3.20`, `wbuntu/alpine:3.22`

**第二步（开发环境）**：

- `wbuntu/node:22`, `wbuntu/node:24`
- `wbuntu/golang:1.24`, `wbuntu/golang:1.25`

**第三步（代理工具和通用应用）**：

- `wbuntu/gost`, `wbuntu/shadowray`, `wbuntu/mtproxy`
- `wbuntu/caddy:2.8.4`
