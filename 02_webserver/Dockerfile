FROM alpine:3.3

MAINTAINER tomwillfixit 

RUN apk update && rm -rf /var/cache/apk/*

COPY webserver.bin / 

EXPOSE 8888

ENTRYPOINT ["/webserver.bin"]

