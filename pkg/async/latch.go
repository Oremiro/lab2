package async

import (
	"sync"
)

var closedChan = make(chan struct{})

func init() {
	close(closedChan)
}

type TaskCancellation struct {
	mu   sync.Mutex
	done chan struct{}
}

func (l *TaskCancellation) Done() <-chan struct{} {
	l.mu.Lock()
	if l.done == nil {
		l.done = make(chan struct{})
	}
	d := l.done
	l.mu.Unlock()
	return d
}

func (l *TaskCancellation) resolve(fn func()) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if fn != nil {
		fn()
	}

	if l.done == nil {
		l.done = closedChan
		return
	}

	select {
	case <-l.done:
		panic(ErrAlreadyResolved)
	default:
		// ...
	}

	close(l.done)
}
