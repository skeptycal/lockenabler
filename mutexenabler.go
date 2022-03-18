package lockenabler

import (
	"errors"
	"io"
)

/// modified logrus.MutexWrap for testing purposes
// removed boolean flag and replaced with different functions

type (
	LockEnablerWriter interface {
		LockEnabler
		io.Writer
	}

	lockEnableWriter struct {
		LockEnabler
		io.Writer
	}
)

func NewLockEnableWriter(w io.Writer) LockEnableWriter {
	return &lockEnableWriter{NewLockEnabler(), DefaultNopWriter(nil)}
}

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
