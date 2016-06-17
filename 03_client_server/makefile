binary:
	echo "Building client and server binaries"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client/client.bin client/client.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server/server.bin server/server.go 

build:
	docker-compose build --no-cache

start_server:
	echo "Starting Server"
	docker-compose up -d

test_client:
	echo "Starting client"
	docker-compose run client /opt/client.bin
logs:
	docker-compose logs

all: clean binary build start_server test_client

clean:
	docker-compose rm -f
