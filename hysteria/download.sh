#!/bin/sh
set -xe
VERSION=$1
TARGETARCH=$2
wget -q https://github.com/HyNetwork/hysteria/releases/download/$VERSION/hysteria-linux-$TARGETARCH -O /usr/bin/hysteria
chmod a+x /usr/bin/hysteria
