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
		"go get -u google.golang.org/protobuf/proto",
		"go install github.com/golang/protobuf/protoc-gen-go@latest",
		"go install github.com/asim/go-micro/cmd/protoc-gen-micro/v3@latest",
		"\ncompile the proto file " + alias + ".proto:\n",
		"cd " + alias,
		"make proto tidy\n",
	}
}

func create(files []file, c config) error {
	for _, file := range files {
		fp := filepath.Join(c.Alias, file.Path)
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
			"dehyphen": func(s string) string {
				return strings.ReplaceAll(s, "-", "")
			},
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

func NewCommand(alias string) *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "Create a project template",
		Subcommands: []*cli.Command{
			{
				Name:   "function",
				Usage:  "Create a function template, e.g. " + alias + " new function greeter",
				Action: Function,
			},
			{
				Name:   "service",
				Usage:  "Create a service template, e.g. " + alias + " new service greeter",
				Action: Service,
			},
		},
	}
}

func Function(ctx *cli.Context) error {
	function := ctx.Args().First()
	if len(function) == 0 {
		return cli.ShowSubcommandHelp(ctx)
	}

	if path.IsAbs(function) {
		fmt.Println("must provide a relative path as function name")
		return nil
	}

	if _, err := os.Stat(function); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", function)
	}

	fmt.Printf("creating function %s\n", function)

	files := []file{
		{".gitignore", tmpl.GitIgnore},
		{"Dockerfile", tmpl.Dockerfile},
		{"Makefile", tmpl.Makefile},
		{"go.mod", tmpl.Module},
		{"handler/" + function + ".go", tmpl.HandlerFNC},
		{"main.go", tmpl.MainFNC},
		{"proto/" + function + ".proto", tmpl.ProtoFNC},
	}
	c := config{
		Alias:    function,
		Comments: protoComments(function),
		Dir:      function,
	}

	return create(files, c)
}

func Service(ctx *cli.Context) error {
	service := ctx.Args().First()
	if len(service) == 0 {
		return cli.ShowSubcommandHelp(ctx)
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
		{"handler/" + service + ".go", tmpl.HandlerSRV},
		{"main.go", tmpl.MainSRV},
		{"proto/" + service + ".proto", tmpl.ProtoSRV},
	}
	c := config{
		Alias:    service,
		Comments: protoComments(service),
		Dir:      service,
	}

	return create(files, c)
}
