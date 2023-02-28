#.proto文件生成.go文件的执行命令
gen:
	protoc --proto_path=./proto ./proto/*.proto --go_out=plugins=grpc:./

clean:
	del ./pb/*.go

run:
	go run main.go
