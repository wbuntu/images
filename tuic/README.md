# tuic

镜像仓库: [https://hub.docker.com/repository/docker/wbuntu/tuic](https://hub.docker.com/repository/docker/wbuntu/tuic)

镜像地址: wbuntu/hysteria:0.8.5

基于alpine:3.15构建镜像，版本与tuic保持一致，目前版本为0.8.5

## 运行容器

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
    "port": 443,
    "token": ["vXdoOCKg120MuGPCg8jtk="],
    "certificate": "/etc/cert/tls.crt",
    "private_key": "/etc/cert/tls.key",
    "ip": "0.0.0.0",
    "congestion_controller": "bbr",
    "max_idle_time": 15000,
    "authentication_timeout": 1000,
    "alpn": ["h3"],
    "max_udp_relay_packet_size": 1500,
    "log_level": "info"
}
```

使用宿主机网络，配置文件与证书挂载至容器内运行

```shell
docker run -d --restart unless-stopped --network host --name tuic -v $PWD/config.json:/etc/tuic/config.json -v $PWD/cert:/etc/cert wbuntu/tuic:0.8.5
```
