# 编译 GO 程序
FROM --platform=$BUILDPLATFORM golang:1.21-alpine as build
ARG TARGETARCH
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=$TARGETARCH
COPY workers-ai-to-openai /workers-ai-to-openai
WORKDIR /workers-ai-to-openai
RUN  go build -ldflags "-s -w" -o workers-ai-to-openai main.go
# 构建镜像
FROM --platform=$TARGETPLATFORM yidadaa/chatgpt-next-web:v2.11.3 as chat
FROM --platform=$TARGETPLATFORM alpine:3.18
ARG TARGETARCH
RUN apk add --no-cache supervisor curl nodejs 
COPY --from=build /workers-ai-to-openai/workers-ai-to-openai /app/workers-ai-to-openai/workers-ai-to-openai
COPY --from=chat /app /app/chatgpt-next-web
COPY supervisord.conf /etc/supervisord.conf
COPY supervisor.d /etc/supervisor.d
EXPOSE 3000
EXPOSE 8080
CMD ["/usr/bin/supervisord","-c","/etc/supervisord.conf"]
