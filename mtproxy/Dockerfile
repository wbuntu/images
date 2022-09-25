FROM wbuntu/debian:10

VOLUME ["/data"]

RUN apt-get update \
    && apt-get install git curl build-essential libssl-dev zlib1g-dev xxd -y \
    && git clone https://github.com/TelegramMessenger/MTProxy && cd MTProxy \
    && make && mv objs/bin/mtproto-proxy /usr/bin/mtproto-proxy && cd - && rm -rf MTProxy \
    && apt-get autoremove build-essential -y && apt-get clean

COPY mtproxy.sh /usr/bin/mtproxy.sh

CMD  ["/usr/bin/mtproxy.sh"]
