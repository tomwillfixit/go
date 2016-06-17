binary: 
	echo "Building webserver go binary"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webserver.bin webserver.go

container:
	echo "Building Alpine Docker Image with webserver binary included"
	docker build -t go:webserver .
run:
	echo "Starting webserver container"
	docker run -d -p 8888:8888 --name webserver go:webserver

all: clean binary container run test
	echo "Building the Go binary, adding it to a Docker image and starting a webserver container"	

test:
	docker ps |grep webserver || echo "Oops! Webserver container is not running"
	curl http://localhost:8888
clean:
	if [ -f webserver.bin ];then rm webserver.bin;fi
	docker rm -f webserver 
