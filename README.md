# Dockercon 2016 Warm up : Simple Go examples for beginners

![dockercon](img/dockercon.png)

With Dockercon 2016 just around the corner it feels like the perfect time to learn Go.

** Disclaimer : I'm not a developer. I just like to learn new stuff. **

# First things first  

Install Go : https://golang.org/doc/install

# Examples

This repo contains 3 examples to whet the appetite.  I've only been using Go for a few days but so far so good.  It is simple and most importantly fun :)

In the examples provided we will build a Go binary on the host and copy it into a Docker Container.

You can also install Go inside the Alpine container and build the binary inside the container.  

In the Helloworld and webserver examples the same make commands can be used

## Helloworld / Webserver

cd into 01_helloworld or 02_webserver directory 

Build the Go binary

```
make binary
```

Create a Docker Container image with the binary baked in
```
make container
```

Start the container
```
make run
```

Test the container is running
```
make test
```

Cleanup
```
make clean
```

## Simple Client / Server example using docker-compose and overlay network

This example will build 2 containers. Both containers will be attached to the same overlay network. See the docker-compose.yml for more details.

Build the Client and Server Go binary

```
make binary
```

Build the Client and Server containers with the binaries baked in
```
make build
```

Start the Server container first
```
make start_server
```

Start the Client and send a test message to the server
```
make test_client

Note : Type 'exit' and press return to exit from the client
```

Verify the Server container was actually responding
```
make logs
```

Cleanup
```
make clean
```

