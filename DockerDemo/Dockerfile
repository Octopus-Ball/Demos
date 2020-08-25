# 指定基础镜像
FROM golang as build

# 通过设置环境变量来设置go代理开启go mod支持
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

# 指定了容器中的一个工作目录，容器启动时执行的命令会在该目录下执行
# 设置工作目录后，之后的RUN、CMD、ENTRYPOINT、ADD、COPY等命令都会在该目录1下执行
# 可使用多个WORKDIR指令(后续指令若为相对路径，则会基于之前命令指定的路径)
# 使用docker run运行容器时，可以通过-w参数覆盖构建时所设置的工作目录
WORKDIR /go/cache

# 将本地文件添加到容器中(tar类型文件会自动解压)，也可以从指定的网络资源加载
# COPY功能与ADD类似，但不会自动解压文件，也不能访问网络资源
ADD go.mod .
ADD go.sum .

# 构建镜像时执行的指令(用来执行命令行命令的)
RUN go mod download

WORKDIR /go/release

ADD . .

# -w 去掉DWARF调试信息(得到的程序就不能用gdb调试了)
# -s 去掉符号表(panic时的stack trace就没有任何文件名/行号信息了)
# -X 设置包中的变量值
RUN go build -ldflags="-s -w" -installsuffix cgo -o app main.go

FROM scratch as prod

# 从之前构建的build拷贝东西到新镜像
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/release/app /
COPY --from=build /go/release/conf.yaml /

CMD ["/app"]

# https://www.cnblogs.com/likeli/p/10521941.html
