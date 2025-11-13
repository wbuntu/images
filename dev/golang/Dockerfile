ARG VERSION
FROM golang:${VERSION}-alpine
ENV TZ=Asia/Shanghai
RUN apk add --no-cache busybox-extras bash zsh curl jq wget git git-lfs make htop tzdata vim
RUN sed -i 's#^\(root:[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:\).*#\1/bin/zsh#' /etc/passwd \
    && wget -qO /tmp/omz_install.sh https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh \
    && sh /tmp/omz_install.sh --unattended || true \
    && sed -i 's/ZSH_THEME="robbyrussell"/ZSH_THEME="tonotdo"/g' $HOME/.zshrc \
    && rm -f /tmp/omz_install.sh
WORKDIR /root
CMD [ "zsh"]
