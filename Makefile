gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=. --go_out=.

clean:
	rm pb/*

server:
	go run cmd/server/main.go -port 5000

server-tls:
	go run cmd/server/main.go -port 5005 -tls

client:
	go run cmd/client/main.go -address 0.0.0.0:5000

client-tls:
	go run cmd/client/main.go -address 0.0.0.0:5005 -tls

test:
	go test -cover -race ./...

cert:
	cd cert; ./gen.sh; cd ..

update-packages:
	go get -u ./...

.PHONY: gen clean server client test cert update-packages