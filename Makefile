build:
	cd pb && protoc hello.proto --go_out=plugins=grpc:./
	go build  -o server server.go
	go build  -o client client.go
clean:
	rm server client pb/*.go
