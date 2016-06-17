#There is no Go installed in Alpine by default.
#Building the binary on the host using :
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client.bin client.go

FROM alpine:3.3

MAINTAINER tomwillfixit 

RUN apk update && rm -rf /var/cache/apk/*

#Copy pre-built client binary into image
COPY client.bin /opt/client.bin 

WORKDIR /opt

CMD /bin/sh
