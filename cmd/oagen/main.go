package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lestrrat-go/openapi/generator/grpcgen"
	"github.com/lestrrat-go/openapi/generator/restclientgen"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "\n%s\n\n", err)
		os.Exit(1)
	}
}

func doMainHelp() {
	fmt.Fprintf(os.Stderr, "oagen [subcommand] [options...] specfile")
	fmt.Fprintf(os.Stderr, "\n\nsubcommands:")
	fmt.Fprintf(os.Stderr, "\n    protobuf   - generate protocol buffers definition file")
	fmt.Fprintf(os.Stderr, "\n    restclient - generate REST client")
	fmt.Fprintf(os.Stderr, "\n\n")
}

type mapFlags map[string]string

func (v *mapFlags) Set(value string) error {
	i := strings.IndexByte(value, '=')
	if i < 1 || i > len(value)-1 {
		return errors.New(`expected key=value`)
	}

	(*v)[value[:i]] = value[i+1:]
	return nil
}

func (v *mapFlags) String() string {
	var keys []string
	for _, key := range *v {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for i, key := range keys {
		buf.WriteString(key)
		buf.WriteByte('=')
		buf.WriteString((*v)[key])
		if i < len(keys)-1 {
			buf.WriteString(", ")
		}
	}
	return buf.String()
}

type protobufCmd struct {
	annotate      bool
	globalOptions mapFlags
	output        string
	packageName   string
}

type restClientCmd struct {
	dir         string
	exportNew   bool
	packageName string
	serviceName string
	clientName  string
	target      string
}

func _main() error {
	if len(os.Args) < 2 {
		doMainHelp()
		return errors.New(`insufficient arguments`)
	}

	switch os.Args[1] {
	case "protobuf":
		return doProtobuf(os.Args[2:])
	case "restclient":
		return doRestClient(os.Args[2:])
	default:
		fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		if err := fs.Parse(os.Args[1:]); err == flag.ErrHelp {
			doMainHelp()
			return nil
		}
		return errors.Errorf(`unknown command %s`, os.Args[1])
	}
}

func doProtobuf(args []string) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var cmd protobufCmd
	cmd.globalOptions = mapFlags{}
	fs := flag.NewFlagSet("protobuf", flag.ContinueOnError)
	fs.Var(&cmd.globalOptions, "global-option", "key=value pair of global options")
	fs.StringVar(&cmd.packageName, "package", "", "package name to use in protobuf declaration")
	fs.StringVar(&cmd.output, "output", "", "protobuf file to write result")
	fs.BoolVar(&cmd.annotate, "annotate", true, "place google.api.http annotation")
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return nil
		}
		return errors.Wrap(err, `failed to parse options`)
	}

	var spec openapi.OpenAPI

	fn := fs.Arg(0)
	if cmd.packageName == "" {
		cmd.packageName = filepath.Base(fn)
		if i := strings.LastIndexByte(cmd.packageName, '.'); i > 0 {
			cmd.packageName = cmd.packageName[:i]
		}
	}

	f, err := os.Open(fn)
	if err != nil {
		return errors.Wrap(err, `failed to open file`)
	}
	defer f.Close()

	parsed, err := openapi.ParseYAML(f, openapi.WithValidate(true))
	if err != nil {
		return errors.Wrapf(err, `failed to parse openapi spec in %s`, fn)
	}

	spec = parsed

	var options []grpcgen.Option

	var dst io.Writer
	if cmd.output == "" {
		dst = os.Stdout
	} else {
		tmpfile, err := ioutil.TempFile("", cmd.output)
		if err != nil {
			return errors.Wrap(err, `failed to create temporary file`)
		}
		defer func() {
			if e := os.Rename(tmpfile.Name(), cmd.output); e != nil {
				err = errors.Wrapf(e, `failed to rename temporary file %s to %s`, tmpfile.Name(), cmd.output)
			}
		}()
		defer tmpfile.Close()
		defer os.Remove(tmpfile.Name())
		dst = tmpfile
	}

	options = append(options, grpcgen.WithDestination(dst))
	options = append(options, grpcgen.WithAnnotation(cmd.annotate))
	options = append(options, grpcgen.WithPackageName(cmd.packageName))
	for key, value := range cmd.globalOptions {
		options = append(options, grpcgen.WithGlobalOption(key, value))
	}

	if err := grpcgen.Generate(ctx, spec, options...); err != nil {
		return errors.Wrap(err, `failed to generate code`)
	}

	dir := filepath.Dir(cmd.output)
	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.Wrapf(err, `failed to create directory %s`, dir)
		}
	}

	if cmd.output == "" {
	}
	return nil
}

func doRestClient(args []string) error {
	var cmd restClientCmd
	fs := flag.NewFlagSet("restclient", flag.ContinueOnError)
	fs.StringVar(&(cmd.dir), "directory", "", "directory name to store the new library")
	fs.StringVar(&cmd.packageName, "package", "", "package name of the new library (only applicable if target=go)")
	fs.StringVar(&cmd.serviceName, "service", "", "default service name for unspecified operations")
	fs.StringVar(&cmd.target, "target", "go", "target runtime to generate")
	fs.BoolVar(&cmd.exportNew, "export-new", true, "control if New() should be exported (only for Go client)")
	fs.StringVar(&cmd.clientName, "client-name", "Client", "name of Client class (only for ES6 client)")
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return nil
		}
		return errors.Wrap(err, `failed to parse options`)
	}
	//	flag.StringVar(&cmd.format, "format", "", "format of file")

	var spec openapi.OpenAPI

	fn := fs.Arg(0)
	f, err := os.Open(fn)
	if err != nil {
		return errors.Wrap(err, `failed to open file`)
	}
	defer f.Close()

	parsed, err := openapi.ParseYAML(f)
	if err != nil {
		return errors.Wrapf(err, `failed to parse openapi spec in %s`, fn)
	}
	spec = parsed

	var options []restclientgen.Option

	options = append(options, restclientgen.WithTarget(cmd.target))
	options = append(options, restclientgen.WithExportNew(cmd.exportNew))
	options = append(options, restclientgen.WithClientName(cmd.clientName))

	if v := cmd.dir; v != "" {
		options = append(options, restclientgen.WithDir(v))
	}

	if v := cmd.packageName; v != "" {
		options = append(options, restclientgen.WithPackageName(v))
	}

	if v := cmd.serviceName; v != "" {
		options = append(options, restclientgen.WithDefaultServiceName(v))
	}

	if err := restclientgen.Generate(spec, options...); err != nil {
		return errors.Wrap(err, `failed to generate code`)
	}

	return nil
}
