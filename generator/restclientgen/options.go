package restclientgen

import "github.com/lestrrat-go/openapi/internal/option"

type Option = option.Interface

const (
	optkeyDirectory = "directory"
	optkeyPackageName = "packageName"
)

func WithDir(s string) Option {
	return option.New(optkeyDirectory, s)
}

func WithPackageName(s string) Option {
	return option.New(optkeyPackageName, s)
}
