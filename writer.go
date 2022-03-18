package lockenabler

import "io"

var (
	discardLEWriter LockEnableWriter = NewLockEnableWriter(io.Discard)
	nopLEWriter     LockEnableWriter = NewLockEnableWriter(nopWriter{})
	lenLEWriter     LockEnableWriter = NewLockEnableWriter(lenWriter{})
)

func defaultNopWriter(w ioWriter) ioWriter {
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
