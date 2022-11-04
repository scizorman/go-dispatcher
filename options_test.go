package dispatcher

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWithMaxWorkers(t *testing.T) {
	const maxWorkers = 10

	opt := &Options{
		maxWorkers:   defaultMaxWorkers,
		jobQueueSize: defaultJobQueueSize,
	}
	WithMaxWorkers(maxWorkers)(opt)

	want := &Options{
		maxWorkers:   maxWorkers,
		jobQueueSize: defaultJobQueueSize,
	}
	if diff := cmp.Diff(want, opt, cmp.AllowUnexported(Options{})); diff != "" {
		t.Errorf("WithMaxWorkers() mismatch (-want, +got):\n%s", diff)
	}
}

func TestWithJobQueueSize(t *testing.T) {
	const jobQueueSize = 10000

	opt := &Options{
		maxWorkers:   defaultMaxWorkers,
		jobQueueSize: defaultJobQueueSize,
	}
	WithJobQueueSize(jobQueueSize)(opt)

	want := &Options{
		maxWorkers:   defaultMaxWorkers,
		jobQueueSize: jobQueueSize,
	}
	if diff := cmp.Diff(want, opt, cmp.AllowUnexported(Options{})); diff != "" {
		t.Errorf("WithJobQueueSize() mismatch (-want, +got):\n%s", diff)
	}
}
