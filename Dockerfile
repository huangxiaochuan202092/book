# 使用官方的 Go 基础镜像，明确指定 Go 版本为最新的 1.20 系列（使用 alpine 轻量级基础镜像）
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 配置 Go 代理，加速依赖下载
ENV GOPROXY=https://goproxy.cn

# 清理 Go 模块缓存，避免旧缓存影响依赖下载
RUN go clean -modcache

# 复制 go.mod 和 go.sum 文件到工作目录
COPY go.mod go.sum ./

# 下载项目依赖
RUN go mod download

# 检查依赖是否正确加载
RUN go list -m all

# 复制项目代码到工作目录
COPY . .

# 设置文件权限
RUN chmod -R 755 .

# 构建可执行文件，并将构建日志输出到文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main . 2>&1 | tee build.log

# 检查构建是否成功
RUN if [ ! -f ./main ]; then echo "Build failed, check build.log for details"; exit 1; fi

# 使用轻量级的 Alpine 最新版本镜像作为最终镜像
FROM alpine:latest

# 安装必要的工具，这里安装了证书，用于处理 HTTPS 请求
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /app

# 从 builder 阶段复制可执行文件到当前镜像
COPY --from=builder /app/main .

# 假设项目有静态文件目录 static，复制它到当前镜像
COPY --from=builder /app/static ./static

# 确保可执行文件有执行权限
RUN chmod +x ./main

# 暴露端口，根据你的实际情况修改
EXPOSE 8080

# 启动应用程序
CMD ["./main"]