make pwdmanager_proto_update:
	protoc --go_out=. ./pwdmanager/v1/structs/passwd.proto
	protoc -I=. pwdmanager/v1/service/service.proto --go-grpc_out=.