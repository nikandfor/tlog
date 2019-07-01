package tlog

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/nikandfor/json"
)

type (
	ID int64

	Labels []string

	Logger interface {
		Writer

		Printf(f string, args ...interface{})
		Start() *Span
		Spawn(ID) *Span
	}

	Writer interface {
		Labels(ls Labels)
		SpanStarted(s *Span)
		SpanFinished(s *Span)
		Message(l *Message, s *Span)
	}

	Message struct {
		Location Location
		Time     time.Duration
		Format   string
		Args     []interface{}
	}

	Span struct {
		l Logger

		ID     ID
		Parent ID

		Location Location

		Started time.Time
		Elapsed time.Duration

		Flags int
	}

	SimpleLogger struct {
		Writer
	}

	initLogger struct{}

	ConsoleWriter struct {
		w  io.Writer
		tf string
	}

	JSONWriter struct {
		mu sync.Mutex
		w  *json.Writer
		ls map[Location]struct{}
	}

	FilterWriter struct {
		w    Writer
		args []string

		mu sync.RWMutex
		c  map[Location]bool
	}

	TeeWriter struct {
		mu      sync.Mutex
		Writers []Writer
	}

	Rand interface {
		Int63() int64
	}

	ConcurrentRand struct {
		mu  sync.Mutex
		rnd Rand
	}
)

const (
	FlagError = 1 << iota

	FlagNone = 0
)

var ( // time, rand
	now      = time.Now
	rnd Rand = &ConcurrentRand{rnd: rand.New(rand.NewSource(now().UnixNano()))}
)

var ( // defaults
	DefaultLabels   Labels
	DefaultLogger   Logger = initLogger{}
	DumpDefaultInfo        = true
)

func init() {
	h, err := os.Hostname()
	if h == "" && err != nil {
		h = err.Error()
	}

	DefaultLabels = Labels{
		"_hostname=" + h,
		fmt.Sprintf("_pid=%d", os.Getpid()),
	}
}

func NewLogger(w Writer) Logger {
	l := &SimpleLogger{Writer: w}
	if DumpDefaultInfo {
		l.Labels(DefaultLabels)
		l.Printf("!os.Args: %q", os.Args)
	}

	return l
}

func Printf(f string, args ...interface{}) {
	newmessage(DefaultLogger, nil, f, args)
}

func newspan(l Logger, par ID) *Span {
	s := &Span{
		l:        l,
		ID:       ID(rnd.Int63()),
		Parent:   par,
		Location: funcentry(2),
		Started:  now(),
	}
	l.SpanStarted(s)
	return s
}

func newmessage(l Logger, s *Span, f string, args []interface{}) {
	if l == nil {
		return
	}

	var t time.Duration
	if s == nil {
		t = time.Duration(now().UnixNano())
	} else {
		t = now().Sub(s.Started)
	}

	l.Message(
		&Message{
			Location: location(2),
			Time:     t,
			Format:   f,
			Args:     args,
		},
		s,
	)
}

func Start() *Span {
	if DefaultLogger == nil {
		return nil
	}

	return newspan(DefaultLogger, 0)
}

func Spawn(id ID) *Span {
	if DefaultLogger == nil || id == 0 {
		return nil
	}

	return newspan(DefaultLogger, id)
}

func (l *SimpleLogger) Printf(f string, args ...interface{}) {
	newmessage(l, nil, f, args)
}

func (l *SimpleLogger) Start() *Span {
	if l == nil {
		return nil
	}

	return newspan(l, 0)
}

func (l *SimpleLogger) Spawn(id ID) *Span {
	if l == nil || id == 0 {
		return nil
	}

	return newspan(l, id)
}

func (l initLogger) Labels(ls Labels)            { l.init(); DefaultLogger.Labels(ls) }
func (l initLogger) Message(m *Message, s *Span) { l.init(); DefaultLogger.Message(m, s) }
func (l initLogger) SpanStarted(s *Span)         { l.init(); DefaultLogger.SpanStarted(s) }
func (l initLogger) SpanFinished(s *Span)        { l.init(); DefaultLogger.SpanFinished(s) }

func (l initLogger) Printf(f string, args ...interface{}) {
	l.init()
	newmessage(DefaultLogger, nil, f, args)
}

func (l initLogger) Start() *Span {
	l.init()

	return newspan(DefaultLogger, 0)
}

func (l initLogger) Spawn(id ID) *Span {
	l.init()

	return newspan(DefaultLogger, id)
}

func (l initLogger) init() {
	DefaultLogger = NewLogger(NewConsoleWriter(os.Stderr))
}

func (s *Span) Printf(f string, args ...interface{}) {
	if s == nil {
		return
	}

	newmessage(s.l, s, f, args)
}

func (s *Span) Finish() {
	if s == nil {
		return
	}

	s.Elapsed = now().Sub(s.Started)
	s.l.SpanFinished(s)
}

func NewConsoleWriter(w io.Writer) *ConsoleWriter {
	return &ConsoleWriter{
		w:  w,
		tf: "2006-01-02_15:04:05.000000",
	}
}

func (w *ConsoleWriter) Message(m *Message, s *Span) {
	t := time.Unix(0, m.Time.Nanoseconds())
	endl := ""
	if l := len(m.Format); l == 0 || m.Format[l-1] != '\n' {
		endl = "\n"
	}
	fmt.Fprintf(w.w, "%v %-20v "+m.Format+endl, append([]interface{}{t.Format(w.tf), m.Location.String()}, m.Args...)...)
}

func (w *ConsoleWriter) SpanStarted(s *Span) {
	fmt.Fprintf(w.w, "%v %-20v !Span started  %v\n", s.Started.Format(w.tf), s.Location.String(), s.ID)
}

func (w *ConsoleWriter) SpanFinished(s *Span) {
	fmt.Fprintf(w.w, "%v %-20v !Span finished %v - elapsed %v\n", s.Started.Format(w.tf), s.Location.String(), s.ID, s.Elapsed)
}

func (w *ConsoleWriter) Labels(ls Labels) {
	w.Message(
		&Message{
			Location: location(1),
			Time:     time.Duration(now().UnixNano()),
			Format:   "!Labels: %q",
			Args:     []interface{}{ls},
		},
		nil,
	)
}

func (w *FilterWriter) Message(m *Message, s *Span) {
	if !w.should(m.Location, true) {
		return
	}
	w.w.Message(m, s)
}

func (w *FilterWriter) SpanStarted(s *Span) {
	if !w.should(s.Location, false) {
		return
	}
	w.w.SpanStarted(s)
}

func (w *FilterWriter) SpanFinished(s *Span) {
	if !w.should(s.Location, false) {
		return
	}
	w.w.SpanFinished(s)
}

func (w *FilterWriter) should(l Location, msg bool) bool {
	w.mu.RLock()
	r, ok := w.c[l]
	w.mu.RUnlock()
	if ok {
		return r
	}

	defer w.mu.Unlock()
	w.mu.Lock()

	r = w.compile(l, msg)

	w.c[l] = r

	return r
}

func (w *FilterWriter) compile(l Location, msg bool) (r bool) {
	file, _ := l.FileLine()
	//	fnc := l.Func()

	for _, a := range w.args {
		r = strings.Contains(file, a)
	}

	return r
}

func NewJSONWriter(w *json.Writer) *JSONWriter {
	return &JSONWriter{
		w:  w,
		ls: make(map[Location]struct{}),
	}
}

func (w *JSONWriter) Labels(ls Labels) {
	defer w.mu.Unlock()
	w.mu.Lock()

	w.w.ObjStart()

	w.w.ObjKey([]byte("L"))

	w.w.ArrayStart()

	for _, l := range ls {
		w.w.SafeStringString(l)
	}

	w.w.ArrayEnd()

	w.w.ObjEnd()

	w.w.NewLine()
}

func (w *JSONWriter) Message(m *Message, s *Span) {
	msg := fmt.Sprintf(m.Format, m.Args...)

	defer w.mu.Unlock()
	w.mu.Lock()

	if _, ok := w.ls[m.Location]; !ok {
		w.location(m.Location)
	}

	w.w.ObjStart()

	w.w.ObjKey([]byte("m"))

	w.w.ObjStart()

	w.w.ObjKey([]byte("l"))
	fmt.Fprintf(w.w, "%d", m.Location)

	w.w.ObjKey([]byte("t"))
	fmt.Fprintf(w.w, "%d", m.Time.Nanoseconds()/1000)

	w.w.ObjKey([]byte("m"))
	w.w.SafeStringString(msg)

	if s != nil {
		w.w.ObjKey([]byte("s"))
		fmt.Fprintf(w.w, "%d", s.ID)
	}

	w.w.ObjEnd()

	w.w.ObjEnd()

	w.w.NewLine()
}

func (w *JSONWriter) SpanStarted(s *Span) {
	defer w.mu.Unlock()
	w.mu.Lock()

	if _, ok := w.ls[s.Location]; !ok {
		w.location(s.Location)
	}

	w.w.ObjStart()

	w.w.ObjKey([]byte("s"))

	w.w.ObjStart()

	w.w.ObjKey([]byte("id"))
	fmt.Fprintf(w.w, "%d", s.ID)

	if s.Parent != 0 {
		w.w.ObjKey([]byte("p"))
		fmt.Fprintf(w.w, "%d", s.Parent)
	}

	w.w.ObjKey([]byte("l"))
	fmt.Fprintf(w.w, "%d", s.Location)

	w.w.ObjKey([]byte("s"))
	fmt.Fprintf(w.w, "%d", s.Started.UnixNano()/1000)

	w.w.ObjEnd()

	w.w.ObjEnd()

	w.w.NewLine()
}

func (w *JSONWriter) SpanFinished(s *Span) {
	defer w.mu.Unlock()
	w.mu.Lock()

	w.w.ObjStart()

	w.w.ObjKey([]byte("f"))

	w.w.ObjStart()

	w.w.ObjKey([]byte("id"))
	fmt.Fprintf(w.w, "%d", s.ID)

	w.w.ObjKey([]byte("e"))
	fmt.Fprintf(w.w, "%d", s.Elapsed.Nanoseconds()/1000)

	if s.Flags != 0 {
		w.w.ObjKey([]byte("F"))
		fmt.Fprintf(w.w, "%x", s.Flags)
	}

	w.w.ObjEnd()

	w.w.ObjEnd()

	w.w.NewLine()
}

func (w *JSONWriter) location(l Location) {
	file, line := l.FileLine()
	fnc := l.Func()

	w.w.ObjStart()

	w.w.ObjKey([]byte("l"))

	w.w.ObjStart()

	w.w.ObjKey([]byte("pc"))
	fmt.Fprintf(w.w, "%d", l)

	w.w.ObjKey([]byte("f"))
	w.w.SafeStringString(file)

	w.w.ObjKey([]byte("l"))
	fmt.Fprintf(w.w, "%d", line)

	w.w.ObjKey([]byte("n"))
	w.w.StringString(fnc)

	w.w.ObjEnd()

	w.w.ObjEnd()

	w.w.NewLine()

	w.ls[l] = struct{}{}
}

func NewTeeWriter(w ...Writer) *TeeWriter {
	return &TeeWriter{Writers: w}
}

func (w *TeeWriter) Labels(ls Labels) {
	defer w.mu.Unlock()
	w.mu.Lock()

	for _, w := range w.Writers {
		w.Labels(ls)
	}
}

func (w *TeeWriter) Message(m *Message, s *Span) {
	defer w.mu.Unlock()
	w.mu.Lock()

	for _, w := range w.Writers {
		w.Message(m, s)
	}
}

func (w *TeeWriter) SpanStarted(s *Span) {
	defer w.mu.Unlock()
	w.mu.Lock()

	for _, w := range w.Writers {
		w.SpanStarted(s)
	}
}

func (w *TeeWriter) SpanFinished(s *Span) {
	defer w.mu.Unlock()
	w.mu.Lock()

	for _, w := range w.Writers {
		w.SpanFinished(s)
	}
}

func (ls *Labels) Set(k, v string) {
	val := k
	if v != "" {
		val += "=" + v
	}

	for i, l := range *ls {
		if l == k || strings.HasPrefix(l, k+"=") {
			(*ls)[i] = val
			return
		}
	}
	*ls = append(*ls, val)
}

func (ls *Labels) Get(k string) (string, bool) {
	for _, l := range *ls {
		if l == k {
			return "", true
		} else if strings.HasPrefix(l, k+"=") {
			return l[len(k)+1:], true
		}
	}
	return "", false
}

func (ls *Labels) Del(k string) bool {
	for i, l := range *ls {
		if l == k || strings.HasPrefix(l, k+"=") {
			ll := len(*ls) - 1
			(*ls)[i] = (*ls)[ll]
			(*ls) = (*ls)[:ll]
			return true
		}
	}
	return false
}

func (ls *Labels) Merge(b Labels) {
out:
	for _, add := range b {
		for _, have := range *ls {
			if add == have {
				continue out
			}
		}

		*ls = append(*ls, add)
	}
}

func (i ID) String() string {
	if i == 0 {
		return "________________"
	}
	return fmt.Sprintf("%016x", uint64(i))
}

func (m *Message) AbsTime() time.Time {
	return time.Unix(0, int64(m.Time))
}

func (m *Message) SpanID() ID {
	if m == nil || m.Args == nil {
		return 0
	}
	return m.Args[0].(ID)
}

func (r *ConcurrentRand) Int63() int64 {
	defer r.mu.Unlock()
	r.mu.Lock()

	return r.rnd.Int63()
}
