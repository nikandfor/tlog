package tlog

import (
	"io"
	"testing"
)

func NewTestLogger(t testing.TB, v string, tostderr io.Writer) *Logger {
	w := tostderr
	ff := LdetFlags

	if t != nil && w == nil {
		w = newTestingWriter(t)
		ff = 0
	}

	tl := New(NewConsoleWriter(w, ff))

	if v != "" {
		tl.SetFilter(v)
	}

	return tl
}
