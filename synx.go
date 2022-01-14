package synx

import "sync"

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
