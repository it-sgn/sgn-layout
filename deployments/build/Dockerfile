FROM registry.cn-shanghai.aliyuncs.com/kbk/golang:1.20 AS builder

COPY . /build
WORKDIR /build

RUN make build && mv /build/bin/* /build/bin/server


FROM registry.cn-shanghai.aliyuncs.com/kbk/alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add tzdata 
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime 
RUN echo "Asia/Jakarta" > /etc/timezone

#Copy binaries from stage 1 to stage 2
WORKDIR /app

COPY --from=builder /build/bin/server .

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf/"]
