package dispatcher

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

// Dispatcher is a job dispatcher.
type Dispatcher[T any] struct {
	worker Worker[T]
	queue  chan T
	sem    *semaphore.Weighted
	wg     sync.WaitGroup
}

// New returns a new Dispatcher.
func New[T any](w Worker[T], optFns ...OptionsFunc) *Dispatcher[T] {
	opt := &Options{
		maxWorkers:   defaultMaxWorkers,
		jobQueueSize: defaultJobQueueSize,
	}
	for _, optFn := range optFns {
		optFn(opt)
	}
	return &Dispatcher[T]{
		worker: w,
		queue:  make(chan T, opt.jobQueueSize),
		sem:    semaphore.NewWeighted(int64(opt.maxWorkers)),
	}
}

// Start starts the Dispatcher.
func (d *Dispatcher[T]) Start(ctx context.Context) {
	d.wg.Add(1)
	go func(ctx context.Context) {
		var wg sync.WaitGroup
	LOOP:
		for {
			select {
			case <-ctx.Done():
				wg.Wait()
				break LOOP
			case j := <-d.queue:
				wg.Add(1)
				if err := d.sem.Acquire(ctx, 1); err != nil {
					break LOOP
				}
				go func(ctx context.Context, job T) {
					defer wg.Done()
					defer d.sem.Release(1)
					d.worker.Do(ctx, job)
				}(ctx, j)
			}
		}
		d.wg.Done()
	}(ctx)
}

// Wait waits for the Dispatcher to stop.
func (d *Dispatcher[T]) Wait() {
	d.wg.Wait()
}

// Dispatch dispatches a job.
func (d *Dispatcher[T]) Dispatch(ctx context.Context, job T) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case d.queue <- job:
		return nil
	}
}
