package tlflag

import (
	"os"
	"testing"

	"github.com/nikandfor/assert"
	"github.com/nikandfor/errors"

	"github.com/nikandfor/tlog"
	"github.com/nikandfor/tlog/convert"
	"github.com/nikandfor/tlog/rotated"
	"github.com/nikandfor/tlog/tlio"
	"github.com/nikandfor/tlog/tlz"
)

type testFile string

func TestingFileOpener(n string, f int, m os.FileMode) (interface{}, error) {
	return testFile(n), nil
}

func TestFileWriter(t *testing.T) {
	OpenFileWriter = TestingFileOpener

	const CompressorBlockSize = 1 * tlz.MiB

	w, err := OpenWriter("stderr")
	assert.NoError(t, err)
	assert.Equal(t, tlio.MultiWriter{
		tlog.NewConsoleWriter(tlog.Stderr, tlog.LstdFlags),
	}, w)

	w, err = OpenWriter("stderr?dm,stderr?dm")
	assert.NoError(t, err)
	assert.Equal(t, tlio.MultiWriter{
		tlog.NewConsoleWriter(tlog.Stderr, tlog.LdetFlags|tlog.Lmilliseconds),
		tlog.NewConsoleWriter(tlog.Stderr, tlog.LdetFlags|tlog.Lmilliseconds),
	}, w)

	w, err = OpenWriter(".tl,-.tl")
	assert.NoError(t, err)
	assert.Equal(t, tlio.MultiWriter{
		tlio.NopCloser{Writer: tlog.Stderr},
		tlio.NopCloser{Writer: tlog.Stdout},
	}, w)

	w, err = OpenWriter(".tlz")
	assert.NoError(t, err)
	assert.Equal(t, tlio.MultiWriter{
		tlz.NewEncoder(tlog.Stderr, CompressorBlockSize),
	}, w)

	w, err = OpenWriter("file.tl")
	assert.NoError(t, err)
	assert.Equal(t, testFile("file.tl"), w)

	w, err = OpenWriter("file.tlz")
	assert.NoError(t, err)
	assert.Equal(t, tlio.WriteCloser{
		Writer: tlz.NewEncoder(testFile("file.tlz"), CompressorBlockSize),
		Closer: testFile("file.tlz"),
	}, w)

	w, err = OpenWriter("file.tl.ez")
	assert.NoError(t, err)
	assert.Equal(t, tlio.WriteCloser{
		Writer: tlz.NewEncoder(testFile("file.tl.ez"), CompressorBlockSize),
		Closer: testFile("file.tl.ez"),
	}, w)

	w, err = OpenWriter("file.ezdump")
	assert.NoError(t, err)
	assert.Equal(t, tlio.WriteCloser{
		Writer: tlz.NewEncoder(tlz.NewDumper(testFile("file.ezdump")), tlz.MiB),
		Closer: testFile("file.ezdump"),
	}, w)

	w, err = OpenWriter("file.json")
	assert.NoError(t, err)
	assert.Equal(t, tlio.WriteCloser{
		Writer: convert.NewJSON(testFile("file.json")),
		Closer: testFile("file.json"),
	}, w)

	w, err = OpenWriter("file.json.ez")
	assert.NoError(t, err)
	assert.Equal(t, tlio.WriteCloser{
		Writer: convert.NewJSON(
			tlz.NewEncoder(
				testFile("file.json.ez"),
				CompressorBlockSize)),
		Closer: testFile("file.json.ez"),
	}, w)
}

func TestURLWriter(t *testing.T) { //nolint:dupl
	OpenFileWriter = func(n string, f int, m os.FileMode) (interface{}, error) {
		return testFile(n), nil
	}

	w, err := OpenWriter("relative/path.tlog")
	assert.NoError(t, err)
	assert.Equal(t, testFile("relative/path.tlog"), w)

	w, err = OpenWriter("file://relative/path.tlog")
	assert.NoError(t, err)
	assert.Equal(t, testFile("relative/path.tlog"), w)

	w, err = OpenWriter("file:///absolute/path.tlog")
	assert.NoError(t, err)
	assert.Equal(t, testFile("/absolute/path.tlog"), w)
}

func TestRotatedWriter(t *testing.T) {
	OpenFileWriter = TestingFileOpener

	const CompressorBlockSize = 1 * tlz.MiB

	with := func(f *rotated.File, wrap func(*rotated.File)) *rotated.File {
		wrap(f)

		return f
	}

	w, err := OpenWriter("file_XXXX.tl")
	assert.NoError(t, err)
	assert.Equal(t, with(rotated.Create("file_XXXX.tl"), func(f *rotated.File) {
		f.OpenFile = openFileWriter
	}), w)

	w, err = OpenWriter("file.tl?rotated=1")
	assert.NoError(t, err)
	assert.Equal(t, with(rotated.Create("file.tl"), func(f *rotated.File) {
		f.OpenFile = openFileWriter
	}), w)

	w, err = OpenWriter("file_XXXX.tlz")
	assert.NoError(t, err)
	assert.Equal(t, with(rotated.Create("file_XXXX.tlz"), func(f *rotated.File) {
		f.OpenFile = RotatedTLZFileOpener(nil)
	}), w)

	w, err = OpenWriter("file_XXXX.tl.ez")
	assert.NoError(t, err)
	assert.Equal(t, with(rotated.Create("file_XXXX.tl.ez"), func(f *rotated.File) {
		f.OpenFile = RotatedTLZFileOpener(nil)
	}), w)

	w, err = OpenWriter("file_XXXX.json.ez")
	assert.NoError(t, err)
	assert.Equal(t, tlio.WriteCloser{
		Writer: convert.NewJSON(with(rotated.Create("file_XXXX.json.ez"), func(f *rotated.File) {
			f.OpenFile = RotatedTLZFileOpener(nil)
		})),
		Closer: with(rotated.Create("file_XXXX.json.ez"), func(f *rotated.File) {
			f.OpenFile = RotatedTLZFileOpener(nil)
		}),
	}, w)
}

func TestFileReader(t *testing.T) {
	OpenFileReader = TestingFileOpener

	r, err := OpenReader("stdin")
	assert.NoError(t, err)
	assert.Equal(t, tlio.NopCloser{
		Reader: os.Stdin,
	}, r)

	r, err = OpenReader("./stdin")
	assert.NoError(t, err)
	assert.Equal(t, testFile("./stdin"), r)

	r, err = OpenReader(".tlog.ez")
	assert.NoError(t, err)
	assert.Equal(t, tlio.NopCloser{Reader: tlz.NewDecoder(os.Stdin)}, r)
}

func TestURLReader(t *testing.T) { //nolint:dupl
	OpenFileReader = func(n string, f int, m os.FileMode) (interface{}, error) {
		return testFile(n), nil
	}

	w, err := OpenReader("relative/path.tlog")
	assert.NoError(t, err)
	assert.Equal(t, testFile("relative/path.tlog"), w)

	w, err = OpenReader("file://relative/path.tlog")
	assert.NoError(t, err)
	assert.Equal(t, testFile("relative/path.tlog"), w)

	w, err = OpenReader("file:///absolute/path.tlog")
	assert.NoError(t, err)
	assert.Equal(t, testFile("/absolute/path.tlog"), w)
}

func (testFile) Write(p []byte) (int, error) { return len(p), nil }

func (testFile) Read(p []byte) (int, error) { return 0, errors.New("test mock") }

func (testFile) Close() error { return nil }
