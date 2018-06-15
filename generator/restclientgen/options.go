package restclientgen

import "github.com/lestrrat-go/openapi/internal/option"

type Option = option.Interface

const (
	optkeyDefaultServiceName = "serviceName"
	optkeyDirectory = "directory"
	optkeyExportNew = "exportNew"
	optkeyPackageName = "packageName"
	optkeyTarget = "target"
)

func WithDir(s string) Option {
	return option.New(optkeyDirectory, s)
}

func WithPackageName(s string) Option {
	return option.New(optkeyPackageName, s)
}

func WithDefaultServiceName(s string) Option {
	return option.New(optkeyDefaultServiceName, s)
}

func WithExportNew(v bool) Option {
	return option.New(optkeyExportNew, v)
}

func WithTarget(s string) Option {
	return option.New(optkeyTarget, s)
}
