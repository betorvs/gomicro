#!/usr/bin/env bash

VERSION="$1"
if [ -z "${VERSION}" ]; then
    VERSION="0.0.1"
fi
COMMITHASH=$(git rev-parse HEAD)

go build -v -ldflags "-X \"github.com/betorvs/gomicro/version.Version=${VERSION}\" -X \"github.com/betorvs/gomicro/version.BuildInfo=${COMMITHASH}\""