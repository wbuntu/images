#!/bin/sh
set -xe

VERSION=$1
TARGETARCH=$2

TUICARCH=unknown

if [[ "$TARGETARCH" == "amd64"  ]];then
    TUICARCH=x86_64
elif [[ "$TARGETARCH" == "arm64" ]];then
    TUICARCH=aarch64
else
    echo "unknown TARGETARCH $TARGETARCH"
    exit -1
fi

mkdir -p /etc/tuic
wget -q https://github.com/EAimTY/tuic/releases/download/$VERSION/tuic-server-$VERSION-$TUICARCH-linux-musl -O /usr/bin/tuic-server
chmod a+x /usr/bin/tuic-server
wget -q https://github.com/EAimTY/tuic/releases/download/$VERSION/tuic-client-$VERSION-$TUICARCH-linux-musl -O /usr/bin/tuic-client
chmod a+x /usr/bin/tuic-client
