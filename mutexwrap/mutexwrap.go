package mutexwrap

import (
	"io"
	"sync"

	"github.com/sirupsen/logrus"
	. "github.com/skeptycal/lockenabler/locker"
	. "github.com/skeptycal/lockenabler/writers"
)

/// copy of MutexWrap from logrus for testing purposes

type (

	// Reference: mutexWrap = logrus.MutexWrap

	MutexWrap struct {
		mu       *sync.Mutex
		disabled bool
	}

	logrusMutexWrap struct {
		*logrus.MutexWrap
		disabled bool
	}

	mutexWrapWriter struct {
		LockEnabler
		io.Writer
	}
)

func NewMutexWrapWriter(w io.Writer) LockEnableWriter {
	if w == nil {
		w = NewLockEnableWriter(io.Discard)
	}
	return &mutexWrapWriter{NewMutexWrap(), w}
	// mmm := logrus.MutexWrap{}
}

func NewLogrusWriter(w io.Writer) LockEnableWriter {
	if w == nil {
		w = NewLockEnableWriter(io.Discard)
	}
	return &mutexWrapWriter{NewMutexWrap(), w}
	// mmm := logrus.MutexWrap{}
}

// NewMutexWrap returns a mutex wrapper that wraps
// a sync.Mutex and implements LockEnabler.
// This may be used to add LockEnabler functionality
// to structures that do not implement the interface
// natively.
//
// This is an alternative implementation that is
// based on logrus.MutexWrap{} which uses a
// boolean state variable to indicate whether
// the LockEnabler is enabled or not.
//
// The default implementation is an unlocked,
// enabled sync.Mutex with Lock(), Unlock(),
// Enable(), and Disable() methods.
func NewMutexWrap() LockEnabler {
	f := &MutexWrap{new(sync.Mutex), false}
	// f.SetLockFuncs(nil, nil)   // this is a no-op for now ...
	// f.SetEnableFuncs(nil, nil) // this is a no-op for now ...
	// f.Enable()                 // this is a no-op for now ...
	return f
}

// This is an alternative implementation that
// of MutexWrap specifically uses the actual
// logrus.MutexWrap{} structure.
func NewLogrusWrap() LockEnabler {
	f := &logrusMutexWrap{new(logrus.MutexWrap), false}
	// f.SetLockFuncs(nil, nil)   // this is a no-op for now ...
	// f.SetEnableFuncs(nil, nil) // this is a no-op for now ...
	// f.Enable()                 // this is a no-op for now ...
	return f
}

// Write writes b to the underlying writer,
// returning the number of bytes written and
// any error encountered.
func (mw *mutexWrapWriter) Write(b []byte) (n int, err error) {
	return mw.Writer.Write(b)
}

// If the underlying writer is nil, no further
// processing is done, len(b) and an error are
// returned to maintain consistency with io.Writer.
//
// If the mutexWriter is disabled, 0 and an
// error are returned immediately.

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

func (mw *MutexWrap) Enable()                                    { mw.disabled = false }
func (mw *MutexWrap) Disable()                                   { mw.disabled = true }
func (*MutexWrap) SetLockFuncs(lockFunc, unlockFunc func())      {}
func (*MutexWrap) SetEnableFuncs(enableFunc, disableFunc func()) {}

func (lw *logrusMutexWrap) Enable()                                    { lw.disabled = false }
func (lw *logrusMutexWrap) Disable()                                   { lw.disabled = true }
func (*logrusMutexWrap) SetLockFuncs(lockFunc, unlockFunc func())      {}
func (*logrusMutexWrap) SetEnableFuncs(enableFunc, disableFunc func()) {}
