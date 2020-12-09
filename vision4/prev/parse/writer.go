package parse

import (
	"io"
	"time"

	"github.com/nikandfor/tlog"
)

type (
	Writer interface {
		Labels(Labels) error
		Frame(Frame) error
		Message(Message) error
		Metric(Metric) error
		Meta(Meta) error
		SpanStart(SpanStart) error
		SpanFinish(SpanFinish) error
	}

	AnyWriter struct {
		tlog.Writer
	}

	ConsoleWriter struct {
		*tlog.ConsoleWriter
		lastTime int64
	}

	ConvertWriter struct {
		w  Writer
		ls map[tlog.PC]struct{}
	}

	DiscardWriter struct{}
)

var (
	_ Writer = AnyWriter{}

	_ tlog.Writer = &ConvertWriter{}

	Discard Writer = DiscardWriter{}
)

func NewAnyWiter(w tlog.Writer) AnyWriter {
	return AnyWriter{Writer: w}
}

func (w AnyWriter) Labels(ls Labels) error {
	return w.Writer.Labels(ls.Labels, ls.Span)
}

func (w AnyWriter) Frame(l Frame) error {
	tlog.PC(l.PC).SetCache(l.Name, l.File, l.Line)

	return nil
}

func (w AnyWriter) Meta(m Meta) error {
	return w.Writer.Meta(
		tlog.Meta{
			Type: m.Type,
			Data: m.Data,
		},
	)
}

func (w AnyWriter) Message(m Message) error {
	return w.Writer.Message(
		tlog.Message{
			PC:    tlog.PC(m.PC),
			Time:  m.Time,
			Text:  m.Text,
			Attrs: m.Attrs,
			Level: m.Level,
		},
		m.Span,
	)
}

func (w AnyWriter) Metric(m Metric) error {
	return w.Writer.Metric(
		tlog.Metric{
			Name:  m.Name,
			Value: m.Value,
		},
		m.Span,
	)
}

func (w AnyWriter) SpanStart(s SpanStart) error {
	return w.Writer.SpanStarted(tlog.SpanStart{
		ID:        s.ID,
		Parent:    s.Parent,
		StartedAt: s.StartedAt,
		PC:        tlog.PC(s.PC),
	})
}

func (w AnyWriter) SpanFinish(f SpanFinish) error {
	return w.Writer.SpanFinished(tlog.SpanFinish{
		ID:      f.ID,
		Elapsed: time.Duration(f.Elapsed),
	})
}

func NewConsoleWriter(w io.Writer, ff int) *ConsoleWriter {
	return &ConsoleWriter{ConsoleWriter: tlog.NewConsoleWriter(w, ff)}
}

func (w *ConsoleWriter) Labels(ls Labels) error {
	b, wr := tlog.Getbuf()
	defer wr.Ret(&b)

	b = append(b, "Labels:"...)
	for _, l := range ls.Labels {
		b = append(b, ' ')
		b = append(b, l...)
	}

	return w.ConsoleWriter.Message(
		tlog.Message{
			Time: w.lastTime,
			Text: tlog.UnsafeBytesToString(b),
		},
		ls.Span,
	)
}

func (w *ConsoleWriter) Frame(l Frame) error {
	tlog.PC(l.PC).SetCache(l.Name, l.File, l.Line)

	return nil
}

func (w *ConsoleWriter) Meta(m Meta) error {
	b, wr := tlog.Getbuf()
	defer wr.Ret(&b)

	b = tlog.AppendPrintf(b, "Meta: %v ", m.Type)

	for _, l := range m.Data {
		b = tlog.AppendPrintf(b, " %q", l)
	}

	return w.ConsoleWriter.Message(
		tlog.Message{
			Time: w.lastTime,
			Text: tlog.UnsafeBytesToString(b),
		},
		tlog.ID{},
	)
}

func (w *ConsoleWriter) Message(m Message) error {
	w.lastTime = m.Time

	return w.ConsoleWriter.Message(
		tlog.Message{
			PC:    tlog.PC(m.PC),
			Time:  m.Time,
			Text:  m.Text,
			Attrs: m.Attrs,
			Level: m.Level,
		},
		m.Span,
	)
}

func (w *ConsoleWriter) Metric(m Metric) error {
	b, wr := tlog.Getbuf()
	defer wr.Ret(&b)

	wh := tlog.DefaultStructuredConfig.MessageWidth
	if cfg := w.StructuredConfig; cfg != nil {
		wh = cfg.MessageWidth
	}

	b = tlog.AppendPrintf(b, "%-*v  %15.5f ", wh, m.Name, m.Value)

	for _, l := range m.Labels {
		b = append(b, ' ')
		b = append(b, l...)
	}

	return w.ConsoleWriter.Message(
		tlog.Message{
			Time: w.lastTime,
			Text: tlog.UnsafeBytesToString(b),
		},
		m.Span,
	)
}

func (w *ConsoleWriter) SpanStart(s SpanStart) error {
	w.lastTime = s.StartedAt

	return w.ConsoleWriter.SpanStarted(tlog.SpanStart{
		ID:        s.ID,
		Parent:    s.Parent,
		StartedAt: s.StartedAt,
		PC:        tlog.PC(s.PC),
	})
}

func (w *ConsoleWriter) SpanFinish(f SpanFinish) error {
	b, wr := tlog.Getbuf()
	defer wr.Ret(&b)

	b = tlog.AppendPrintf(b, "Span finished - elapsed %vms", f.Elapsed)

	return w.ConsoleWriter.Message(
		tlog.Message{
			Time: w.lastTime,
			Text: tlog.UnsafeBytesToString(b),
		},
		f.ID,
	)
}

func NewConvertWriter(w Writer) *ConvertWriter {
	return &ConvertWriter{
		w:  w,
		ls: make(map[tlog.PC]struct{}),
	}
}

func (w *ConvertWriter) Labels(ls tlog.Labels, sid ID) error {
	return w.w.Labels(Labels{Labels: ls, Span: sid})
}

func (w *ConvertWriter) Meta(m tlog.Meta) error {
	return w.w.Meta(
		Meta{
			Type: m.Type,
			Data: m.Data,
		},
	)
}

func (w *ConvertWriter) Message(m tlog.Message, sid tlog.ID) error {
	err := w.location(m.PC)
	if err != nil {
		return err
	}

	return w.w.Message(Message{
		Span: sid,
		PC:   uint64(m.PC),
		Time: m.Time,
		Text: m.Text,
	})
}

func (w *ConvertWriter) Metric(m tlog.Metric, sid tlog.ID) error {
	return w.w.Metric(Metric{
		Span:  sid,
		Name:  m.Name,
		Value: m.Value,
	})
}

func (w *ConvertWriter) SpanStarted(s tlog.SpanStart) error {
	err := w.location(s.PC)
	if err != nil {
		return err
	}

	return w.w.SpanStart(SpanStart{
		ID:        s.ID,
		Parent:    s.Parent,
		PC:        uint64(s.PC),
		StartedAt: s.StartedAt,
	})
}

func (w *ConvertWriter) SpanFinished(f tlog.SpanFinish) error {
	return w.w.SpanFinish(SpanFinish{
		ID:      f.ID,
		Elapsed: f.Elapsed.Nanoseconds(),
	})
}

func (w *ConvertWriter) location(l tlog.PC) error {
	if _, ok := w.ls[l]; ok {
		return nil
	}

	name, file, line := l.NameFileLine()

	err := w.w.Frame(Frame{
		PC:   uint64(l),
		Name: name,
		File: file,
		Line: line,
	})
	if err != nil {
		return err
	}

	w.ls[l] = struct{}{}

	return nil
}

func (w DiscardWriter) Frame(l Frame) error           { return nil }
func (w DiscardWriter) Labels(ls Labels) error        { return nil }
func (w DiscardWriter) Meta(m Meta) error             { return nil }
func (w DiscardWriter) Message(m Message) error       { return nil }
func (w DiscardWriter) Metric(m Metric) error         { return nil }
func (w DiscardWriter) SpanStart(s SpanStart) error   { return nil }
func (w DiscardWriter) SpanFinish(f SpanFinish) error { return nil }