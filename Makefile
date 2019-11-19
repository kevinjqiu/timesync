build: protogen
	go build

protogen:
	protoc -I pkg/ pkg/timesync.proto --go_out=plugins=grpc:pkg
