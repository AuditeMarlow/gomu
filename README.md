# Gomu

Gomu is a helper tool for developing [Go Micro][1] projects.

## Getting Started

[Download][2] and install **Go**. Version `1.16` or higher is required

Installation is done by using the [`go install`][3] command.

```bash
go install github.com/auditemarlow/gomu@latest
```

Let's create a new service using the `new` command.

```bash
gomu new service helloworld
```

Follow the on-screen instructions. Next, we can run the program.

```bash
cd example
go run main.go
```

Finally we can call the service.

```bash
gomu call helloworld Helloworld.Call '{"name": "John"}'
```

That's all you need to know to get started. Refer to the [Go Micro][1]
documentation for more info on developing services.

## Dependencies

You will need protoc-gen-micro for code generation

* [protobuf][4]
* [protoc-gen-go][5]
* [protoc-gen-micro][6]

```bash
# Download latest proto release
# https://github.com/protocolbuffers/protobuf/releases
go get -u github.com/golang/protobuf/protoc
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3
```

## Creating A Service

To create a new service, use the `gomu new service` command.

```bash
$ gomu new service helloworld
creating service helloworld

download protoc zip packages (protoc-$VERSION-$PLATFORM.zip) and install:

visit https://github.com/protocolbuffers/protobuf/releases/latest

download protobuf for go-micro:

go get -u github.com/golang/protobuf/protoc
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3

compile the proto file helloworld.proto:

cd helloworld
make proto tidy
```

To create a new function, use the `gomu new function` command. Functions differ
from services in that they exit after returning.

```bash
$ gomu new function helloworld
creating function helloworld

download protoc zip packages (protoc-$VERSION-$PLATFORM.zip) and install:

visit https://github.com/protocolbuffers/protobuf/releases/latest

download protobuf for go-micro:

go get -u github.com/golang/protobuf/protoc
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3

compile the proto file helloworld.proto:

cd helloworld
make proto tidy
```

## Running A Service

To run a service, `cd` into its directory, generate the protobuf code, install
its dependencies and run the program.

```bash
make proto tidy
go run main.go
```

### With Docker

To run a service with Docker, `cd` into its directory, generate the protobuf
code, build the Docker image and run the Docker container.

```bash
make proto docker
docker run helloworld:latest
```

## Calling A Service

To call a service, use the `gomu call` command. This will send a single request
and expect a single response.

```bash
$ gomu call helloworld Helloworld.Call '{"name": "John"}'
{"msg":"Hello John"}
```

To call a service's server stream, use the `gomu stream server` command. This
will send a single request and expect a stream of responses.

```bash
$ gomu stream server helloworld Helloworld.ServerStream '{"count": 10}'
{"count":0}
{"count":1}
{"count":2}
{"count":3}
{"count":4}
{"count":5}
{"count":6}
{"count":7}
{"count":8}
{"count":9}
```

To call a service's bidirectional stream, use the `gomu stream bidi` command.
This will send a stream of requests and expect a stream of responses.

```bash
$ gomu stream bidi helloworld Helloworld.BidiStream '{"stroke": 1}' '{"stroke": 2}' '{"stroke": 3}'
{"stroke":1}
{"stroke":2}
{"stroke":3}
```

## License

This software is published under the [MIT license][7].

[1]: https://github.com/asim/go-micro
[2]: https://golang.org/dl/
[3]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[4]: https://grpc.io/docs/protoc-installation/
[5]: https://micro.mu/github.com/golang/protobuf/protoc-gen-go
[6]: https://github.com/asim/go-micro/tree/master/cmd/protoc-gen-micro
[7]: LICENSE
