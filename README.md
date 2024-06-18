# douban_api.rixingyike.com
给《小程序从 0 到 1》读者提供的后端 api.rixingyike.com 程序源码，用于替换已经变化的豆瓣接口。

## 从编码编译

依次进行了如下操作：

```bash 
go mod init github.com/rixingyike/douban_api.rixingyike.com
go mod tidy
go install github.com/go-bindata/go-bindata/v3/go-bindata@v3.1.3
go-bindata -o assets.go assets/  # assets.go 文件是生成的
go get github.com/kataras/iris/v12
```

## 启动

./dev.sh 

## 即时编译

./dev.sh 

如果有问题，到公众号「艺述论」找我。