package tlog

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/nikandfor/errors"

	"github.com/nikandfor/tlog/low"
)

type (
	ID [16]byte

	// ShortIDError is an ID parsing error.
	ShortIDError struct {
		Bytes int // Bytes successfully parsed
	}

	concurrentRand struct {
		mu sync.Mutex
		r  *rand.Rand
	}
)

var rnd = &concurrentRand{r: rand.New(rand.NewSource(time.Now().UnixNano()))} //nolint:gosec

// String returns short string representation.
//
// It's not supposed to be able to recover it back to the same value as it was.
func (id ID) String() string {
	var b [8]byte
	id.FormatTo(b[:], 'v')
	return string(b[:])
}

// StringFull returns full id represented as string.
func (id ID) StringFull() string {
	var b [32]byte
	id.FormatTo(b[:], 'v')
	return string(b[:])
}

// IDFromBytes decodes ID from bytes slice.
//
// If byte slice is shorter than type length result is returned as is and ShortIDError as error value.
// You may use result if you expected short ID prefix.
func IDFromBytes(b []byte) (id ID, err error) {
	n := copy(id[:], b)

	if n < len(id) {
		err = ShortIDError{Bytes: n}
	}

	return
}

// IDFromString parses ID from string.
//
// If parsed string is shorter than type length result is returned as is and ShortIDError as error value.
// You may use result if you expected short ID prefix (profuced by ID.String, for example).
func IDFromString(s string) (id ID, err error) {
	if "________________________________"[:len(s)] == s {
		return
	}

	var i int
	var c byte
	for ; i < len(s); i++ {
		switch {
		case '0' <= s[i] && s[i] <= '9':
			c = s[i] - '0'
		case 'a' <= s[i] && s[i] <= 'f':
			c = s[i] - 'a' + 10
		default:
			err = hex.InvalidByteError(s[i])
			return
		}

		if i&1 == 0 {
			c <<= 4
		}

		id[i>>1] |= c
	}

	if i < 2*len(id) {
		err = ShortIDError{Bytes: i / 2}
	}

	return
}

// IDFromStringAsBytes is the same as IDFromString. It avoids alloc in IDFromString(string(b)).
func IDFromStringAsBytes(s []byte) (id ID, err error) {
	if bytes.Equal([]byte("________________________________")[:len(s)], s) {
		return
	}

	n, err := hex.Decode(id[:], s)
	if err != nil {
		return
	}

	if n < len(id) {
		return id, ShortIDError{Bytes: n}
	}

	return id, nil
}

// ShouldID wraps IDFrom* call and skips error if any.
func ShouldID(id ID, err error) ID {
	return id
}

// MustID wraps IDFrom* call and panics if error occurred.
func MustID(id ID, err error) ID {
	if err != nil {
		panic(err)
	}

	return id
}

// Error is an error interface implementation.
func (e ShortIDError) Error() string {
	return fmt.Sprintf("too short id: %d bytes, wanted %d", e.Bytes, len(ID{}))
}

// Format is fmt.Formatter interface implementation.
// It supports width. '+' flag sets width to full ID length.
func (id ID) Format(s fmt.State, c rune) {
	var buf0 [32]byte
	buf := low.NoEscapeBuffer(buf0[:])

	w := 8
	if W, ok := s.Width(); ok {
		w = W
	}
	if s.Flag('+') {
		w = 2 * len(id)
	}
	id.FormatTo(buf[:w], c)
	_, _ = s.Write(buf[:w])
}

// FormatTo is alloc free Format alternative.
func (id ID) FormatTo(b []byte, f rune) {
	if id == (ID{}) {
		if f == 'v' || f == 'V' {
			copy(b, "________________________________")
		} else {
			copy(b, "00000000000000000000000000000000")
		}
		return
	}

	const digitsx = "0123456789abcdef"
	const digitsX = "0123456789ABCDEF"

	dg := digitsx
	if f == 'X' || f == 'V' {
		dg = digitsX
	}

	m := len(b)
	if 2*len(id) < m {
		m = 2 * len(id)
	}

	ji := 0
	for j := 0; j+1 < m; j += 2 {
		b[j] = dg[id[ji]>>4]
		b[j+1] = dg[id[ji]&0xf]
		ji++
	}

	if m&1 == 1 {
		b[m-1] = dg[id[m>>1]>>4]
	}
}

func (id ID) MarshalJSON() ([]byte, error) {
	b := make([]byte, len(id)*2+2)
	b[0] = '"'
	b[len(b)-1] = '"'

	id.FormatTo(b[1:], 'x')

	return b, nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	if len(b) < 4 {
		return errors.New("bad id")
	}

	if b[0] != '"' || b[len(b)-1] != '"' {
		return errors.New("bad id encoding")
	}

	x, err := IDFromStringAsBytes(b[1 : len(b)-1])
	if err != nil {
		return err
	}

	*id = x

	return nil
}

func MathRandID() (id ID) {
	rnd.mu.Lock()

	for id == (ID{}) {
		_, _ = rnd.r.Read(id[:])
	}

	rnd.mu.Unlock()

	return
}

func RandIDFromReader(read func(p []byte) (int, error)) func() ID {
	return func() (id ID) {
		n, err := read(id[:])
		if err != nil {
			panic(err)
		}
		if n != len(id) {
			panic(n)
		}

		return id
	}
}

/* will repeat at most after 2 ** (32 - 2) ids
func FastRandID() (id ID) {
	*(*uint32)(unsafe.Pointer(&id[0])) = fastrand()
	*(*uint32)(unsafe.Pointer(&id[4])) = fastrand()
	*(*uint32)(unsafe.Pointer(&id[8])) = fastrand()
	*(*uint32)(unsafe.Pointer(&id[12])) = fastrand()
	return
}
*/

// UUID creates ID generation function.
// read is a random Read method. Function panics on Read error.
// read must be safe for concurrent use.
//
// It's got from github.com/google/uuid.
func UUID(read func(p []byte) (int, error)) func() ID {
	return func() (uuid ID) {
		n, err := read(uuid[:])
		if err != nil {
			panic(err)
		}
		if n != len(uuid) {
			panic(n)
		}

		uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
		uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10

		return uuid
	}
}
