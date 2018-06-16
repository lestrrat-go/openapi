package es6flow

import "github.com/lestrrat-go/openapi/internal/option"

type Option = option.Interface

const (
	optkeyDefaultServiceName = "serviceName"
	optkeyDirectory = "directory"
)

func WithDir(s string) Option {
	return option.New(optkeyDirectory, s)
}

func WithDefaultServiceName(s string) Option {
	return option.New(optkeyDefaultServiceName, s)
}
