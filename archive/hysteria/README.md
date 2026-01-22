# hysteria

镜像仓库: [https://hub.docker.com/repository/docker/wbuntu/hysteria](https://hub.docker.com/repository/docker/wbuntu/hysteria)

镜像地址: wbuntu/hysteria:v1.2.0

基于alpine:3.15构建镜像，版本与hysteria保持一致，目前版本为v1.2.0

## 服务端

证书示例

```shell
➜  cert ls -lh
total 12K
-rw-r--r-- 1 root root 5.5K Sep 19 10:19 tls.crt
-rw-r--r-- 1 root root 1.7K Sep 19 10:19 tls.key
```

config.json示例

```json
{
  "listen": ":443",
  "protocol": "udp",
  "server_name": "hysteria.example.com",
  "cert": "/etc/cert/tls.crt",
  "key": "/etc/cert/tls.key",
  "auth": {
      "mode": "passwords",
      "config": [
        "ffff4eYWCVUImlBVtoWcfBCLRCSY"
      ]
  },
  "obfs": "eeeebzh8SHcw2d7lQBs",
  "alpn": "hysteria",
  "up_mbps": 100,
  "down_mbps": 100
}
```

使用宿主机网络，配置文件与证书挂载至容器内运行

```shell
docker run -d --restart unless-stopped --network host --name hysteria -v $PWD/config.json:/etc/hysteria/config.json -v $PWD/cert:/etc/cert wbuntu/hysteria:v1.2.0
```

## 客户端

config.json示例

```json
{
    "protocol": "udp",
    "server": "server_ip:443",
    "server_name": "hysteria.example.com",
    "auth_str": "ffff4eYWCVUImlBVtoWcfBCLRCSY",
    "obfs": "eeeebzh8SHcw2d7lQBs",
    "alpn": "hysteria",
    "up_mbps": 10,
    "down_mbps": 50,
    "http": {
        "listen": "127.0.0.1:2080"
    },
    "socks5": {
        "listen": "127.0.0.1:2085"
    }
}
```

使用宿主机网络，配置文件挂载至容器内运行

```shell
docker run -d --restart unless-stopped --network host --name hysteria -v $PWD/config.json:/etc/hysteria/config.json wbuntu/hysteria:v1.2.0 /usr/bin/hysteria -c /etc/hysteria/config.json client
```

