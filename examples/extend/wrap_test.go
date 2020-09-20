package extend

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nikandfor/tlog"
)

func TestWrap(t *testing.T) {
	var buf bytes.Buffer

	l := tlog.New(tlog.NewConsoleWriter(&buf, 0))

	w := LoggerWith(l, Attrs{{"field", "value"}, {"int_filed", 12}})
	w.Printw("message", Attrs{{"add", "one more"}})

	assert.Equal(t, `message                                 field=value  int_filed=12  add=one more
`, buf.String())
}

func BenchmarkPrintf(b *testing.B) {
	b.ReportAllocs()

	var w tlog.CountableIODiscard

	l := tlog.New(tlog.NewConsoleWriter(&w, 0))

	for i := 0; i < b.N; i++ {
		w := LoggerWith(l, Attrs{{"i", i}})
		w.Printw("message", Attrs{{"j", i}})
	}
}
