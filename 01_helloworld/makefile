binary: 
	echo "Building helloworld go binary"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o helloworld.bin helloworld.go

container:
	echo "Building Alpine Docker Image with helloworld binary included"
	docker build -t go:helloworld .
run:
	echo "Starting helloworld container"
	docker run -d -p 8888:8888 --name helloworld go:helloworld

all: clean binary container run test
	echo "Building the Go binary, adding it to a Docker image and starting a helloworld container"	

test:
	docker ps |grep helloworld || echo "Oops! Helloworld container is not running"
	curl http://localhost:8888
clean:
	if [ -f helloworld.bin ];then rm helloworld.bin;fi
	docker rm -f helloworld
