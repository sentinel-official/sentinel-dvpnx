FROM golang:1.24-alpine3.22 AS build

COPY . /root/sentinel-dvpnx/

RUN --mount=target=/go/pkg/mod,type=cache \
    --mount=target=/root/.cache/go-build,type=cache \
    apk add autoconf automake bash file g++ gcc git libtool linux-headers make musl-dev unbound-dev && \
    cd /root/sentinel-dvpnx/ && make --jobs=$(nproc) install && \
    git clone --branch=master --depth=1 https://github.com/handshake-org/hnsd.git /root/hnsd && \
    cd /root/hnsd/ && bash autogen.sh && sh configure && make --jobs=$(nproc)

FROM alpine:3.22

COPY --from=build /go/bin/sentinel-dvpnx /usr/local/bin/sentinel-dvpnx
COPY --from=build /root/hnsd/hnsd /usr/local/bin/hnsd

RUN apk add --no-cache iptables openvpn unbound-libs v2ray wireguard-tools && \
    rm -rf /etc/v2ray/ /usr/share/v2ray/

ENTRYPOINT ["sentinel-dvpnx"]
