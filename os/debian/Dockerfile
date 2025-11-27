ARG VERSION
FROM debian:$VERSION
ENV TZ=Asia/Shanghai
RUN apt-get update && apt-get install -y zsh curl jq wget git git-lfs make htop vim && rm -rf /var/lib/apt/lists/*
RUN sed -i 's#^\(root:[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:\).*#\1/usr/bin/zsh#' /etc/passwd \
    && wget -qO /tmp/omz_install.sh https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh \
    && sh /tmp/omz_install.sh --unattended || true \
    && sed -i 's/ZSH_THEME="robbyrussell"/ZSH_THEME="tonotdo"/g' $HOME/.zshrc \
    && rm -f /tmp/omz_install.sh
WORKDIR /root
CMD [ "zsh" ]
