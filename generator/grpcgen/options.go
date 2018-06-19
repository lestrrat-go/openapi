package grpcgen

import (
	"io"

	"github.com/lestrrat-go/openapi/internal/option"
)

func WithDestination(dst io.Writer) Option {
	return option.New(optkeyDestination, dst)
}

func WithPackageName(name string) Option {
	return option.New(optkeyPackageName, name)
}

func WithAnnotation(b bool) Option {
	return option.New(optkeyAnnotation, b)
}

func WithGlobalOption(key, value string) Option {
	return option.New(optkeyGlobalOption, &globalOption{name: key, value: value})
}
