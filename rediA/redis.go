package rediA

import (
	"bufio"
	"io"
)

var (
	arrayPrefixSlice = []byte{'*'}
)

type RespWriter struct {
	*bufio.Writer
}

func NewRespWriter(w io.Writer) *RespWriter {
	return &RespWriter{
		Writer: bufio.NewWriter(w),
	}
}

func (w *RespWriter) WriteCommand(args ...string) error {
	return nil
}
