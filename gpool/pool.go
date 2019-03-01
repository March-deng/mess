package gpool

import (
	"context"
	"log"
	"time"
)

type Pool struct {
	reader      SourceReader
	workers     []*worker
	montiorTick *time.Ticker
	level       int
}

func NewPool(reader SourceReader) *Pool {
	return &Pool{
		reader: reader,
	}
}

func NewPoolWithFunc(reader SourceReader, Func ExecFunc) *Pool {
	return &Pool{
		reader: reader,
	}
}

func (p *Pool) montior(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-p.montiorTick.C:
			level, err := p.reader.GetJobAmout(time.Now(), time.Now())
			if err != nil {
				log.Println(err)
				continue
			}
			p.level = level
		}
	}
}

func (p *Pool) start(ctx context.Context) {
	for {
		select {
		default:
		}
	}
}
