FROM golang:1.17-alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
&& apk update && apk add  bash

WORKDIR /alertmanager-webhook
COPY . .
RUN chmod +x ./scripts/build.sh && ./scripts/build.sh

FROM golang:1.17-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
&& apk update && apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone \
&& apk del tzdata

WORKDIR /alertmanager-webhook
COPY --from=0 /alertmanager-webhook/bin/monitor ./
COPY --from=0 /alertmanager-webhook/config ./config

ENTRYPOINT ["./alertmanager-webhook-wechat"]
