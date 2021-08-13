package template

var Module = `module {{.Dir}}

go 1.16

require (
	github.com/asim/go-micro/v3 v3.5.2
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
`
