# 使用官方 Go 镜像作为基础镜像
FROM golang:1.21.3-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用轻量级的 alpine 镜像作为最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 从 builder 阶段复制编译好的应用
COPY --from=builder /app/main .

# 创建上传目录
RUN mkdir -p ./uploads

# 复制 upload.html 文件
COPY upload.html .

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]
