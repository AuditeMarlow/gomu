package new

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	tmpl "github.com/auditemarlow/gomu/cmd/cli/new/template"
	"github.com/urfave/cli/v2"
)

type config struct {
	Alias    string
	Comments []string
	Dir      string
}

type file struct {
	Path string
	Tmpl string
}

func protoComments(alias string) []string {
	return []string{
		"\ndownload protoc zip packages (protoc-$VERSION-$PLATFORM.zip) and install:\n",
		"visit https://github.com/protocolbuffers/protobuf/releases/latest",
		"\ndownload protobuf for go-micro:\n",
		"go get -u github.com/golang/protobuf/protoc",
		"go get -u github.com/golang/protobuf/protoc-gen-go",
		"go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3",
		"\ncompile the proto file " + alias + ".proto:\n",
		"cd " + alias,
		"make tidy proto\n",
	}
}

func Run(ctx *cli.Context) error {
	service := ctx.Args().First()
	if len(service) == 0 {
		fmt.Println("must provide a service name")
		return nil
	}

	if path.IsAbs(service) {
		fmt.Println("must provide a relative path as service name")
		return nil
	}

	if _, err := os.Stat(service); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", service)
	}

	fmt.Printf("creating service %s\n", service)

	files := []file{
		{".gitignore", tmpl.GitIgnore},
		{"Dockerfile", tmpl.Dockerfile},
		{"Makefile", tmpl.Makefile},
		{"go.mod", tmpl.Module},
		{"handler/" + service + ".go", tmpl.Handler},
		{"main.go", tmpl.Main},
		{"proto/" + service + ".proto", tmpl.Proto},
	}
	c := config{
		Alias:    service,
		Comments: protoComments(service),
		Dir:      service,
	}

	for _, file := range files {
		fp := filepath.Join(service, file.Path)
		dir := filepath.Dir(fp)

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}

		f, err := os.Create(fp)
		if err != nil {
			return err
		}

		fn := template.FuncMap{
			"lower": strings.ToLower,
			"title": func(s string) string {
				return strings.ReplaceAll(strings.Title(s), "-", "")
			},
		}
		t, err := template.New(fp).Funcs(fn).Parse(file.Tmpl)
		if err != nil {
			return err
		}

		err = t.Execute(f, c)
		if err != nil {
			return err
		}
	}

	for _, comment := range c.Comments {
		fmt.Println(comment)
	}

	return nil
}
