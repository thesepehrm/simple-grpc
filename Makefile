run-client:
	go run cmd/grpc-client/main.go

run-server:
	go run cmd/grpc-server/main.go

gen:
	protoc --proto_path=proto proto/*.proto --go_opt=module=github.com/thesepehrm/simple-grpc --go_out=plugins=grpc:.

clean:
	rm -rf pb/*