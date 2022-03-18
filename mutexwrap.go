package lockenabler

import (
	"errors"
	"io"
	"sync"
)

/// copy of MutexWrap from logrus for testing purposes

type (
	mutexWrapperWriter = LockEnableWriter

	// Reference: mutexWrap = logrus.MutexWrap

	MutexWrap struct {
		mu       *sync.Mutex
		disabled bool
	}

	mutexWrapWriter struct {
		MutexWrap
		lockEnableWriter
	}
)

func NewWrapWriter(w ioWriter) mutexWrapperWriter {

	if w == nil {
		w = io.Discard
	}

	mw := mutexWrapWriter{}
	mw.mu = &sync.Mutex{}
	mw.disabled = true
	mw.ioWriter = w
	return &mw
}

func newMutexWrap() LockEnabler {
	return &MutexWrap{}
}

// Write writes b to the underlying writer,
// returning the number of bytes written and
// any error encountered.
//
// If the underlying writer is nil, no further
// processing is done, len(b) and an error are
// returned to maintain consistency with io.Writer.
//
// If the mutexWriter is disabled, 0 and an
// error are returned.
func (mw *mutexWrapWriter) Write(b []byte) (n int, err error) {
	if mw.disabled {
		return 0, errors.New("mutex writer disabled")
	}

	if mw.ioWriter == nil {
		return len(b), errors.New("mutex writer is nil")
	}
	return mw.ioWriter.Write(b)
}

/// implement LockEnabler
//  Lock, Unlock, Enable, Disable, SetLockFuncs, SetEnableFuncs

func (mw *MutexWrap) Lock() {
	if !mw.disabled {
		mw.mu.Lock()
	}
}

func (mw *MutexWrap) Unlock() {
	if !mw.disabled {
		mw.mu.Unlock()
	}
}

func (mw *MutexWrap) Enable()  { mw.disabled = false }
func (mw *MutexWrap) Disable() { mw.disabled = true }

// SetLockFuncs is a no-op function required to
// implement LockEnabler.
func (*MutexWrap) SetLockFuncs(lockFunc, unlockFunc func()) {}

// SetEnableFuncs is a no-op function required to
// implement LockEnabler.
func (*MutexWrap) SetEnableFuncs(enableFunc, disableFunc func()) {}
