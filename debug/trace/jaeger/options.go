package jaeger

import (
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-lib/metrics"
)

var (
	DefaultFromEnv      = false
	DefaultGlobalTracer = true
	DefaultLogger       = jaeger.StdLogger
	DefaultMetrics      = metrics.NullFactory
)

type Options struct {
	Name         string
	FromEnv      bool
	GlobalTracer bool
	Logger       jaeger.Logger
	Metrics      metrics.Factory
}

type Option func(o *Options)

func newOptions(opts ...Option) Options {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	if options.Logger == nil {
		options.Logger = DefaultLogger
	}

	if options.Metrics == nil {
		options.Metrics = DefaultMetrics
	}

	return options
}

func Name(s string) Option {
	return func(o *Options) {
		o.Name = s
	}
}

func FromEnv(e bool) Option {
	return func(o *Options) {
		o.FromEnv = e
	}
}

func GlobalTracer(e bool) Option {
	return func(o *Options) {
		o.GlobalTracer = e
	}
}

func Logger(l jaeger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

func Metrics(m metrics.Factory) Option {
	return func(o *Options) {
		o.Metrics = m
	}
}
