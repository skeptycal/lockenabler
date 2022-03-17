package lockenabler

import (
	"io"
	"sync"
)

/// copy of MutexWrap from logrus for testing purposes

func NewWrapWriter(w io.Writer) LockEnableWriter {

	if w == nil {
		w = io.Discard
	}

	mw := mutexWrapWriter{}
	mw.mu = &sync.Mutex{}
	mw.disabled = true
	mw.Writer = w
	return &mw
}

type mutexWrapWriter struct {
	MutexWrap
	io.Writer
}

// TODO: for testing purposes
func (mw *mutexWrapWriter) Write(b []byte) (n int, err error) {
	if mw.Writer == nil {
		return 0, nil
	}
	return mw.Writer.Write(b)
}

type MutexWrap struct {
	mu       *sync.Mutex
	disabled bool
}

func (mw *MutexWrap) Disable() { mw.disabled = true }
func (mw *MutexWrap) Enable()  { mw.disabled = false }

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
