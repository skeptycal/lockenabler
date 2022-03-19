package lockenabler

import "io"

var (
	discardWriter LockEnableWriter = NewLockEnableWriter(io.Discard)
	nopWriter     LockEnableWriter = NewLockEnableWriter(nopWrite{})
	lenWriter     LockEnableWriter = NewLockEnableWriter(lenWrite{})
)

func defaultNopWriter(w ioWriter) ioWriter {
	if w == nil {
		return io.Discard
	}
	return w
}

// LenWriter returns a LockEnableWriter
// that does not write bytes! It simply
// returns the length of the input []byte
// value and nil.
// This is designed to be used for mocking,
// testing, or for situations where locking
// and enabling features of LockEnableWriter
// are desired but the implementation of
// a writer is not.
func LenWriter(w ioWriter) LockEnableWriter {
	if w == nil {
		return lenWriter
	}
	return lenWriter
}

// LenWriter returns a LockEnableWriter
// that does not write bytes! It returns
// 0, nil immediately.
// This is designed to be used for mocking,
// testing, or for situations where locking
// and enabling features of LockEnableWriter
// are desired but the implementation of
// a writer is not.
func NopWriter(w io.Writer) LockEnableWriter {
	if w == nil {
		return nopWriter
	}
	return nopWriter
}

type nopWrite struct{}

// Write returns 0, nil no matter what the
// input is and does not other processing.
func (nopWrite) Write(b []byte) (n int, err error) {
	return 0, nil
}

type lenWrite struct{}

// Write returns len(b), nil no matter what the
// input is and does not other processing.
func (lenWrite) Write(b []byte) (n int, err error) {
	return len(b), nil
}
