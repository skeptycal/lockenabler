package lockenabler

import "io"

var (
	DiscardWriter LockEnableWriter = NewLockEnableWriter(io.Discard)
	NopWriter     LockEnableWriter = NewLockEnableWriter(nopWriter{})
	LenWriter     LockEnableWriter = NewLockEnableWriter(lenWriter{})
)

func DefaultNopWriter(w io.Writer) io.Writer {
	if w == nil {
		return io.Discard
	}
	return w
}

type nopWriter struct{}

// Write returns 0, nil no matter what the
// input is and does not other processing.
func (nopWriter) Write(b []byte) (n int, err error) {
	return 0, nil
}

type lenWriter struct{}

// Write returns len(b), nil no matter what the
// input is and does not other processing.
func (lenWriter) Write(b []byte) (n int, err error) {
	return len(b), nil
}
