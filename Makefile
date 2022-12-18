start:
	docker compose -f docker-compose.yml up -d --build

stop:
	docker compose -f docker-compose.yml down

compile-proto:
	protoc --go_out=src/server/grpc --proto_path=src/server/grpc/proto src/server/grpc/proto/*_message.proto
	protoc --go-grpc_out=src/server/grpc --proto_path=src/server/grpc/proto src/server/grpc/proto/*_service.proto
