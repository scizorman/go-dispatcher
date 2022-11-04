package dispatcher

const (
	defaultMaxWorkers   = 4
	defaultJobQueueSize = 100
)

// OptionsFunc is a type alias for the Options functional option.
type OptionsFunc func(*Options)

// Options are discrete options for the Dispatcher.
type Options struct {
	maxWorkers   int
	jobQueueSize int
}

// WithMaxQueue sets the maximum number of workers that can be running at the same time.
func WithMaxWorkers(n int) OptionsFunc {
	return func(o *Options) {
		o.maxWorkers = n
	}
}

// WithJobQueueSize sets the maximum number of jobs that can be queued.
func WithJobQueueSize(n int) OptionsFunc {
	return func(o *Options) {
		o.jobQueueSize = n
	}
}
