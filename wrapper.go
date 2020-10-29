package worker

import (
	"context"
	"time"

	"github.com/im-kulikov/helium/service"
)

type (
	options struct {
		atStart bool
		timer   time.Duration
	}

	callerFunc func(context.Context)

	Option func(o *options)
)

func WithTimer(v time.Duration) Option {
	return func(o *options) {
		o.timer = v
	}
}

func WithImmediately() Option {
	return func(o *options) {
		o.atStart = true
	}
}

func defaultOptions() *options {
	return &options{
		atStart: false,
		timer:   time.Second * 10,
	}
}

func WrapJob(name string, caller callerFunc, opts ...Option) service.Service {
	cfg := defaultOptions()

	for _, o := range opts {
		o(cfg)
	}

	return &wrk{
		name: name,

		caller: caller,
		cancel: func() {},

		timer:   cfg.timer,
		atStart: cfg.atStart,
	}
}
