##  生成pb文件
```
protoc -I=api/v1/news --go_out=./ api/v1/news/news.proto
```
##  生成GRPC pb文件
```
protoc -I=api/v1/news --go-grpc_out=./ api/v1/news/news.proto 
```