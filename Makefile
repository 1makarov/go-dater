proto-gen:
	protoc --proto_path=. --go_out=plugins=grpc:internal dater.proto