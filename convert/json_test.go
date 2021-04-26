package convert

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/nikandfor/tlog"
	"github.com/nikandfor/tlog/low"
	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	tm := time.Date(2020, time.December, 25, 22, 8, 13, 0, time.FixedZone("Europe/Moscow", int(3*time.Hour/time.Second)))
	tlog.TestSetTime(func() time.Time { return tm }, tm.UnixNano)

	var b low.Buf

	j := NewJSONWriter(&b)
	j.AttachLabels = true
	j.TimeZone = time.FixedZone("MSK", int(3*time.Hour/time.Second))
	j.TimeFormat = time.RFC3339Nano

	l := tlog.New(j)

	l.SetLabels(tlog.Labels{"a=b", "c"})

	l.Printw("message", "str", "arg", "int", 5, "struct", struct {
		A string `json:"a"`
		B int    `tlog:"bb" yaml:"b"`
		C *int   `tlog:"c,omitempty"`
	}{
		A: "A field",
		B: 9,
	})

	exp := `{"t":"2020-12-25T22:08:13\+03:00","L":\["a=b","c"\]}
{"t":"2020-12-25T22:08:13\+03:00","l":"[\w./-]*json_test.go:29","m":"message","str":"arg","int":5,"struct":{"a":"A field","bb":9},"L":\["a=b","c"\]}
`

	exps := strings.Split(exp, "\n")
	ls := strings.Split(string(b), "\n")
	for i := 0; i < len(exps); i++ {
		re := regexp.MustCompile(exps[i])

		var have string
		if i < len(ls) {
			have = ls[i]
		}

		assert.True(t, re.MatchString(have), "expected\n%s\ngot\n%s", exps[i], have)
	}

	for i := len(exps); i < len(ls); i++ {
		assert.True(t, false, "expected\n%s\ngot\n%s", "", ls[i])
	}
}
