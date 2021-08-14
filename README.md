# Gomu

Gomu is a helper tool for developing [Go Micro][1] projects.

## Getting Started

[Download][2] and install **Go**. Version `1.16` or higher is required

Installation is done by using the `go get` command.

```sh
go get github.com/auditemarlow/gomu
```

Let's create a new project using the `new` command.

```sh
gomu new helloworld
```

Follow the on-screen instructions. Next, we can run the program.

```sh
cd example
go run main.go
```

Finally we can call the service.

```sh
gomu call helloworld Helloworld.Call '{"name": "John"}'
```

We can also stream server responses.

```sh
gomu stream helloworld Helloworld.Stream '{"count": 10}'
```

That's all you need to know to start. Refer to the [Go Micro][1] documentation
for more info on developing services.

## License

This software is published under the [MIT license][3].

[1]: https://github.com/asim/go-micro
[2]: https://golang.org/dl/
[3]: LICENSE
