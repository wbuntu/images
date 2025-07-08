ARG VERSION
FROM wbuntu/alpine:3.20 AS tmp
ARG VERSION
RUN wget -L https://github.com/other-blowsnow/n8n-i18n-chinese/releases/download/n8n%40${VERSION}/editor-ui.tar.gz && tar -zxvf editor-ui.tar.gz -C /tmp

ARG VERSION
FROM n8nio/n8n:${VERSION}
COPY --from=tmp /tmp/dist /usr/local/lib/node_modules/n8n/node_modules/n8n-editor-ui/dist
ENV N8N_DEFAULT_LOCALE=zh-CN
ENV N8N_SECURE_COOKIE=false