#!/bin/bash
set -eu

gox \
    -ldflags "-X github.com/cescoferraro/power/cmd.jwt=${POWER_JWT} -X github.com/cescoferraro/power/cmd.version=${VERSION}" \
    -output="/go/bin/power-{{.Arch}}-${VERSION}" \
    -osarch="linux/armv6" -osarch="linux/armv7" -osarch="linux/amd64" \
    .

cp /usr/local/bin/ngrok /go/bin/ngrok
