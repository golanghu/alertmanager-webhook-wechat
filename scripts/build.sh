#!/bin/sh

export CGO_ENABLED=0
export GOARCH=amd64
export GOOS=linux
export GO111MODULE=on

case $1 in
"")
  if [ $(uname) == "Darwin" ]
  then
    export GOOS=darwin
  fi
  ;;
"darwin")
  export GOOS=darwin
  ;;
"windows")
  export GOOS=windows
  ;;
esac

GIT_BRANCH=$(git symbolic-ref --short -q HEAD)
GIT_REV=$(git rev-parse --short HEAD)
BUILD_DATE=$(date "+%Y-%m-%d %H:%M:%S")
GO_VERSION=$(go version)

echo "###############"
echo $GO_VERSION

FLAGS="-extldflags -static -X 'main.goVersion=${GO_VERSION: 11}' -X 'main.gitRev=${GIT_BRANCH} ${GIT_REV}' -X 'main.buildDate=${BUILD_DATE}'"

go build -ldflags "$FLAGS" -a -v -o ./bin/alertmanager-webhook-wechat ./cmd/
