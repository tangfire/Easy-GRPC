protoc -I . --go_out=plugins=grpc:.\hello_grpc .\hello_grpc\hello.proto

protoc -I . --go_out=plugins=grpc:.\hello_grpc .\hello_grpc\type.proto

protoc -I . --go_out=plugins=grpc:.\hello_grpc\multi_service .\hello_grpc\multi_service\multiservice.proto

protoc -I .\hello_grpc\service_proto --go_out=plugins=grpc:.\hello_grpc\service_proto .\hello_grpc\service_proto\common.proto
protoc -I .\hello_grpc\service_proto --go_out=plugins=grpc:.\hello_grpc\service_proto .\hello_grpc\service_proto\order.proto
protoc -I .\hello_grpc\service_proto --go_out=plugins=grpc:.\hello_grpc\service_proto .\hello_grpc\service_proto\video.proto


protoc -I . --go_out=plugins=grpc:.\stream_grpc .\stream_grpc\stream.proto
