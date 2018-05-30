package grpcgen

import (
	"io"

	"github.com/lestrrat-go/openapi/internal/option"
)

func WithDestination(dst io.Writer) Option {
	return option.New(optkeyDestination, dst)
}

func WithAnnotation(b bool) Option {
	return option.New(optkeyAnnotation, b)
}
