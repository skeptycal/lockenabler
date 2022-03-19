package writers

import (
	"errors"
	"io"
)

var (
	discardWriter LockEnableWriter = NewLockEnableWriter(io.Discard)
	nopWriter     LockEnableWriter = NewLockEnableWriter(nopWrite{})
	lenWriter     LockEnableWriter = NewLockEnableWriter(lenWrite{})
)

// Write writes b to the underlying writer,
// returning the number of bytes written and
// any error encountered.
//
// If the underlying writer is nil, no further
// processing is done, len(b) and an error are
// returned to maintain consistency with io.Writer.
//
// If the lockEnableWriter is disabled, 0
// and an error are returned. (this feature
// is not yet implemented.)
func (lew *lockEnableWriter) Write(b []byte) (n int, err error) {

	// TODO: check for disabled writer??
	// if v, ok := lew.LockEnabler.(*locker); ok {
	// 	if v.fnLock == v.noLock {
	// 		return 0, errors.New("LockEnableWriter is disabled")
	// 	}
	// }

	if lew.Writer == nil {
		return len(b), errors.New("writer is nil")
	}
	return lew.Writer.Write(b)
}

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
