protoc -I . --go_out=plugins=grpc:.\hello_grpc .\hello_grpc\hello.proto

protoc -I . --go_out=plugins=grpc:.\hello_grpc .\hello_grpc\type.proto

