FROM golang:1.18-alpine AS builder

# 设置工作目录
WORKDIR /app

ENV GO111MODULE=on  GOPROXY=https://goproxy.cn,direct GIN_MODE=release

RUN cd

COPY . /app

RUN go mod tidy

RUN cd /app/cmd/spg_admin \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o spg_admin
# 获取 Distroless 镜像，只有 650 kB 的大小，是常用的 alpine:latest 的 1/4
FROM alpine

# 将上一阶段构建好的二进制文件复制到本阶段中
COPY --from=builder /app/cmd/spg_admin/ .
COPY --from=builder /app/internal/config/ ./internal/config
COPY --from=builder /app/log ./log

# 设置监听端口
EXPOSE 8081
# 配置启动命令
ENTRYPOINT ["./spg_admin"]