#!/bin/bash
set -eu

gox \
    -ldflags "-X github.com/cescoferraro/power/cmd.jwt=${POWER_JWT}" \
    -output="/go/bin/power-{{.Arch}}-${VERSION}" \
    -osarch="linux/arm" -osarch="linux/amd64" \
    .

cp /usr/local/bin/ngrok /go/bin/ngrok