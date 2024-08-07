#!/bin/bash

VERSION=$(cat version)
GIT_COMMIT=$(git rev-parse HEAD)
BUILD_DATE=$(date '+%Y%m%d%H%M%S')

echo "Version: $VERSION"
echo "Commit: $GIT_COMMIT"
echo "BuildDate: $BUILD_DATE"

LDFLAGS="-X 'github.com/QQGoblin/kmask/pkg/version.MainVersion=${VERSION}' -X 'github.com/QQGoblin/kmask/pkg/version.GitCommit=${GIT_COMMIT}' -X 'github.com/QQGoblin/kmask/pkg/version.BuildDate=${BUILD_DATE}'"

go run pkg/codegen/main.go
go build -ldflags "${LDFLAGS}" -o kmask ./main.go