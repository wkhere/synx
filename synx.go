package synx

import "sync"

// MarkFirst is a combination of WaitGroup and Once,
// allowing to synchronize all workers and mark
// occurence of some condition in one of them,
// like, the error.
// It is similar to but different from x/sync/errgroup,
// which gives back only the first error occured.
// In the case one would want to do something with all the errors
// separately and be interested in an extra mark that any ever happened,
// this is the way to go.
type MarkFirst struct {
	occured bool
	once    sync.Once
	wg      sync.WaitGroup
}

func (a *MarkFirst) Add() { a.wg.Add(1) }

func (a *MarkFirst) Mark() {
	a.once.Do(func() { a.occured = true })
}

func (a *MarkFirst) Done() { a.wg.Done() }

func (a *MarkFirst) Wait() { a.wg.Wait() }

func (a *MarkFirst) Occured() bool { return a.occured }
