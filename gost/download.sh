#!/bin/sh
set -xe

VERSION=$1
TARGETARCH=$2

GOSTVERSION="${VERSION:1}"
GOSTARCH=unknown

if [[ "$TARGETARCH" == "amd64"  ]];then
    GOSTARCH=amd64
elif [[ "$TARGETARCH" == "arm64" ]];then
    GOSTARCH=armv8
else
    echo "unknown TARGETARCH $TARGETARCH"
    exit -1
fi

mkdir -p /etc/gost
wget -q https://github.com/ginuerzh/gost/releases/download/$VERSION/gost-linux-$GOSTARCH-$GOSTVERSION.gz -O gost.gz
gunzip gost.gz
mv gost /usr/bin/gost
chmod a+x /usr/bin/gost
