# 注册3个服务

cd /Users/penny/go/testGoMicro/main

# error
#go run main.go --server_address :8081
#go run main.go --server_address :8082
#go run main.go --server_address :8083

# 一定要&连接，不能分开来执行，因为shell命令是同步的，执行完一下才能执行下一个。
go run main.go --server_address :8081 & go run main.go --server_address :8082 & go run main.go --server_address :8083