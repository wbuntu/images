ARG VERSION
FROM debian:$VERSION
ENV TZ=Asia/Shanghai
RUN apt update && apt install -y zsh curl jq wget git git-lfs make htop
RUN sed -i 's#^\(root:[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:\).*#\1/usr/bin/zsh#' /etc/passwd \
    && wget -qO /tmp/omz_install.sh https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh \
    && sh /tmp/omz_install.sh --unattended || true \
    && sed -i 's/ZSH_THEME="robbyrussell"/ZSH_THEME="tonotdo"/g' $HOME/.zshrc \
    && rm -f /tmp/omz_install.sh
WORKDIR /root
CMD [ "zsh" ]
