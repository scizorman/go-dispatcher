package dispatcher

import (
	"context"
)

// Worker is a job worker.
type Worker[T any] interface {
	// Do does the job.
	Do(ctx context.Context, job T)
}
