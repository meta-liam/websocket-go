# 开发文档

> Go版本的WebSocket库。

## 快速开始

```go
go mod init 
go mod tidy
go run v1/main.go

```

## 发布版本

git tag v1.0.0

## 单元测试

```shell
make utest
```

## 例子

[client测试例子](../examples/v1-client/index.html)

## 打包

```shell
make build-mac
```

Mac测试运行

```shell

cd ./build/mac && chmod a+x ./main && ./main
```

## 参考

[例子1](https://segmentfault.com/a/1190000022034813)
[例子2](https://github.com/gobwas/ws)
[例子3](https://tonybai.com/2019/09/28/how-to-build-websockets-in-go/)
