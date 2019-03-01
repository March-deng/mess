package gpool

import "time"

type ExecFunc func() error
type SourceID int64

type SourceGroup []SourceID

type SourceReader interface {
	GetJobAmout(start, end time.Time) (int, error)
	GetSources(size uint32) (SourceGroup, error)
}
