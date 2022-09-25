# mtproxy

镜像仓库: [https://hub.docker.com/repository/docker/wbuntu/mtproxy](https://hub.docker.com/repository/docker/wbuntu/mtproxy)

镜像地址: wbuntu/mtproxy:v0.3

来自telegram的mtproto-proxy容器化版本，基于debian:10构建镜像，目前版本为v0.3

## 构建镜像

```shell
git clone https://github.com/mikumaycry/mtproxy.git
cd mtproxy
make image
➜  ~ docker images
REPOSITORY                TAG       IMAGE ID       CREATED        SIZE
wbuntu/mtproxy            v0.3      435b1e0d14e2   4 hours ago    276MB
```

## 运行容器

- mtproto-proxy需要获取客户端IP，若服务器网络环境存在NAT时，需要提供NAT信息，启动脚本已集成NAT参数检测
- 启动脚本默认读取三个环境变量
  - PORT：默认 **8888**，用于本地获取统计信息
  - HTTPPORT：默认 **8443**，用于供客户端通过公网连接代理
  - DOMAIN：默认 **cloudflare.com**，用于启用FAKE-TLS，模拟TLS连接
- 容器启动时，自动打印代理链接，格式为 **tg://proxy?server=xx.xx.xx.xx&port=xxxx&secret=xxxxxxxxxx**，复制到telegram中打开即可使用
- 默认使用 **/data** 目录保存动态生成的代理配置文件与密钥文件

### 使用默认配置运行容器

**运行容器**

```shell
docker run -d --restart always  --name mtproxy -v /data:/data -p 8443:8443 wbuntu/mtproxy:v0.3
```

**获取代理链接**

```shell
➜  ~ docker logs -f --tail=100 mtproxy

tg_proxy_url generated: tg://proxy?server=xx.xx.xx.xx&port=xxxx&secret=xxxxxxxxxxx

[16][2021-02-07 10:31:35.501762 local] Invoking engine mtproxy-0.01 compiled at Feb  7 2021 02:31:05 by gcc 8.3.0 64-bit after commit dc0c7f3de40530053189c572936ae4fd1567269b
[16][2021-02-07 10:31:35.502118 local] config_filename = '/data/proxy_multi'
[16][2021-02-07 10:31:35.518014 local] Successfully checked domain cloudflare.com in 0.005 seconds: is_reversed_extension_order = 0, server_hello_encrypted_size = 1873, use_random_encrypted_size = 1
[16][2021-02-07 10:31:35.522009 local] Started as [xx.xx.xx.xx:8888:16:1612693895]
[16][2021-02-07 10:31:35.522504 local] configuration file /data/proxy_multi re-read successfully (752 bytes parsed), new configuration active
[16][2021-02-07 10:31:35.522579 local] main loop
```

### 使用自定义参数运行容器

```shell
docker run -d --restart always --name mtproxy -v /data:/data -p 8443:8443 -e  PORT=8888 -e HTTPPORT=8443 -e DOMAIN=apple.com wbuntu/mtproxy:v0.3
```

### 使用已构建的镜像运行服务

镜像托管于DockerHub，若无法拉取时，可自行构建

```shell
docker run -d --restart always --network host --name mtproxy -v /data:/data -p 8443:8443  -e  PORT=8888 -e HTTPPORT=8443 -e DOMAIN=apple.com wbuntu/mtproxy:v0.3
```
