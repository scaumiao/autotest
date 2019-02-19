# autotest
基于指令的自动化测试服务
```
go get github.com/scaumiao/autotest 
```

cd proto && make stages


## 依赖
#### 1、java环境(略)

#### 2、protoc
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

#### 3、grpc-gateway
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway 
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```
如果遇到网络问题，可采用如下步骤
- 到`$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway`目录
- clone代码`git clone https://github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
- install `go install`


## 启动
- 1、在release中下载vendor.zip，将vendor.zip解压放到项目上
- 2、运行命令 `java -Dwebdriver.chrome.driver=vendor/chromedriver-mac64-2.42 -jar vendor/selenium-server-standalone-3.14.0.jar -port 53188`
- 3、生成相关proto文件`cd proto && make stages`
- 4、进入cmd目录启动`cd cmd && go run main.go`