package template

var Dockerfile = `FROM golang:alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/{{.Alias}}
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc
COPY . /go/src/{{.Alias}}
RUN make init proto tidy build

FROM scratch
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/{{.Alias}}/{{.Alias}} /{{.Alias}}
ENTRYPOINT ["/{{.Alias}}"]
CMD []
`

var DockerIgnore = `.gitignore
Dockerfile{{if .Skaffold}}
resources/
skaffold.yaml{{end}}
`
