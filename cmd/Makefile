.PHONY:proto,build
proto:
	protoc --go_out=../genproto --go_opt=paths=source_relative \
	--go-grpc_out=../genproto --go-grpc_opt=paths=source_relative \
	-I=../pb ../pb/addressbook.proto

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o grpc_test ./main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -a -o us-backend ./main.go