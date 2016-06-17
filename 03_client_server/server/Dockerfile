#There is no Go installed in Alpine by default.
#Building the binary on the host using :
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server.bin server.go

FROM alpine:3.3

MAINTAINER tomwillfixit 

RUN apk update && rm -rf /var/cache/apk/*

#Copy pre-built server binary into image
COPY server.bin /opt/server.bin

WORKDIR /opt

ENTRYPOINT ["/opt/server.bin"]

