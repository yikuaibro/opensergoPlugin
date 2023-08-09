# opensergoPlugin
opensergoPlugin POC Version

在pkg/plugin中执行 make 构建 pluginserver

go run main/main.go或者build

测试plugin是stream proto
```
pkg/plugin
    ├── assets
    │   └── README.md
    ├── assets.go
    ├── client
    │   └── client.go
    ├── config
    │   ├── config.go
    │   └── config.yaml
    ├── main
    │   ├── main
    │   └── main.go
    ├── Makefile
    ├── pl
    │   ├── builtin
    │   │   ├── const.go
    │   │   └── stream
    │   ├── plugin
    │   │   ├── load.go
    │   │   ├── options.go
    │   │   ├── plugin.go
    │   │   └── store
    │   └── plugin.go
    ├── proto
    │   ├── stream
    │   │   ├── stream_grpc.pb.go
    │   │   └── stream.pb.go
    │   └── stream.proto
    ├── scripts
    │   └── plugin.sh
    ├── server
    │   └── stream
    │       └── main.go
    └── util
        └── const.go
