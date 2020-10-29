package worker

import (
	"context"
	"sync"
	"time"
)

type wrk struct {
	name string
	once sync.Once

	caller func(context.Context)
	cancel context.CancelFunc

	atStart bool
	timer   time.Duration
}

func (w *wrk) Start(ctx context.Context) error {
	w.once.Do(func() {
		ctx, w.cancel = context.WithCancel(ctx)

		go func() {
			if w.atStart {
				w.caller(ctx)
			}

			done := ctx.Done()
			tick := time.NewTimer(w.timer)

		loop:
			for {
				select {
				case <-done:
					break loop
				case <-tick.C:
					w.caller(ctx)
					tick.Reset(w.timer)
				}
			}
		}()
	})

	return nil
}

func (w *wrk) Stop() error {
	w.cancel()

	return nil
}

func (w *wrk) Name() string { return w.name }
