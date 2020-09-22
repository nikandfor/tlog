package tlog

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrameFillCallers(t *testing.T) {
	st := make(Frames, 1)

	st = CallersFill(0, st)

	assert.Len(t, st, 1)
	assert.Equal(t, "frames_test.go:14", st[0].String())
}

func testFramesInside() (st Frames) {
	func() {
		func() {
			st = Callers(1, 3)
		}()
	}()
	return
}

func TestFrameFramesString(t *testing.T) {
	var st Frames
	func() {
		func() {
			st = testFramesInside()
		}()
	}()

	assert.Len(t, st, 3)
	assert.Equal(t, "frames_test.go:24", st[0].String())
	assert.Equal(t, "frames_test.go:25", st[1].String())
	assert.Equal(t, "frames_test.go:33", st[2].String())

	re := `tlog.testFramesInside.func1                                   at [\w.-/]*frames_test.go:24
tlog.testFramesInside                                         at [\w.-/]*frames_test.go:25
tlog.TestFrameFramesString.func1.1                            at [\w.-/]*frames_test.go:33
`
	ok, err := regexp.MatchString(re, st.String())
	assert.NoError(t, err)
	assert.True(t, ok, "expected:\n%v\ngot:\n%v\n", re, st.String())
}

func TestFrameFramesFormat(t *testing.T) {
	var st Frames
	func() {
		func() {
			st = testFramesInside()
		}()
	}()

	assert.Equal(t, "frames_test.go:24 at frames_test.go:25 at frames_test.go:55", fmt.Sprintf("%v", st))

	assert.Equal(t, "testFramesInside.func1:24 at testFramesInside:25 at TestFrameFramesFormat.func1.1:55", fmt.Sprintf("%#v", st))

	re := `at [\w.-/]*frames_test.go:24
at [\w.-/]*frames_test.go:25
at [\w.-/]*frames_test.go:55
`
	v := fmt.Sprintf("%+v", st)
	assert.True(t, regexp.MustCompile(re).MatchString(v), "expected:\n%vgot:\n%v", re, v)
}
