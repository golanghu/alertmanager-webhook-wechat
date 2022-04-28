PROJECT="alertmanager-webhook-wechat"
DOCKERFILE="Dockerfile"

ifeq ($(GOPATH),)
	PATH := $(HOME)/go/bin:$(PATH)
else
	PATH := $(GOPATH)/bin:$(PATH)
endif

export GO111MODULE=on

default:
	scripts/build.sh linux

image: default
	docker build -t ${PROJECT} -f ${DOCKERFILE} .

darwin:
	scripts/build.sh darwin

clean:
	rm -f bin/*

all: clean image

.PHONY: clean default