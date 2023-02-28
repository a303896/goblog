#构建go编译环境
FROM golang:latest
#设置go环境变量
ENV GOPROXY https://goproxy.cn,direct
#指定工作目录
WORKDIR /build
#将对应的源代码copy至镜像
COPY . /build
#交叉编译 '-s -w': 压缩编译后的体积 -o 输出为blog
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o blog
#构建运行时环境
FROM alpine
#安装时区依赖
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
	&& apk add --no-cache tzdata
#copy编译好的go程序至镜像
COPY --from=0 /build/blog /
#挂载配置文件
VOLUME /conf
#挂载日志文件
VOLUME /runtime
#暴露端口
EXPOSE 8888
# 指定容器运行时入口程序 blog
ENTRYPOINT ["/blog"]