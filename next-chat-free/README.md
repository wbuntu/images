# next-chat-free

整合 [chatgpt-next-web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web) 前端与 [aurora](https://github.com/aurora-develop/aurora) 后端，部署后可使用免费的 gpt-3.5-turbo 模型聊天。

默认使用 3000 端口提供网页服务，8080 端口提供 aurora 接口服务。

**运行容器**

这里启动一个容器暴露网页端口：

```shell
docker run -d --name next-chat-free -p 3000:3000 wbuntu/next-chat-free:v0.0.1
```

**验证 chatgpt 接口**

这里在容器内执行一个 curl 命令，验证当前的 IP 是否可访问免费的 chatgpt 网页接口，正常情况下可以看到返回的消息中 role 为 assistant 消息 content 输出了 **this is a test!**

```shell
➜  ~ docker exec next-chat-free curl -s --location 'http://127.0.0.1:8080/v1/chat/completions' --data '{"model": "gpt-3.5-turbo","messages":[{"role": "user", "content": "Say this is a test!"}]}'
{"id":"chatcmpl-QXlha2FBbmROaXhpZUFyZUF3ZXNvbWUK","object":"chat.completion","created":0,"model":"gpt-3.5-turbo-0125","usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0},"choices":[{"index":0,"message":{"role":"assistant","content":"This is a test!"},"finish_reason":null}]}
```

**访问网页**

![](img-01.png)

**设置访问密码**

可以使用 CODE 环境变量设置访问网页的密码添加防护：

```shell
docker run -d --name next-chat-free -p 3000:3000 -e CODE=RkdJ1r0+B9zkksS5S wbuntu/next-chat-free:v0.0.1
```

> [!NOTE]  
> chatgpt-web-next 与 aurora 使用了一些同名的环境变量，可以在 ini 文件中单独为程序设置环境变量，避免冲突产生问题