package lockenabler

import (
	"io"
	"sync"
)

/// modified logrus.MutexWrap for testing purposes
// removed boolean flag and replaced with different functions

func NewEnablerWriter(w io.Writer) LockEnableWriter {
	if w == nil {
		w = io.Discard
	}

	mw := mutexEnableWriter{}
	mw.mu = &sync.Mutex{}
	mw.Writer = w
	mw.Enable()
	return &mw
}

type mutexEnableWriter struct {
	MutexEnable
	io.Writer
}

// TODO: for testing purposes
func (mw *mutexEnableWriter) Write(b []byte) (n int, err error) {
	if mw.Writer == nil {
		return 0, nil
	}
	return mw.Writer.Write(b)
}

type MutexEnabler interface {
	Locker
	Enabler
}

type MutexEnable struct {
	mu       *sync.Mutex
	fnLock   func()
	fnUnlock func()
}

func (mw *MutexEnable) Lock()   { mw.fnLock() }
func (mw *MutexEnable) Unlock() { mw.fnUnlock() }

func (mw *MutexEnable) yesLock()   { mw.mu.Lock() }
func (mw *MutexEnable) yesUnlock() { mw.mu.Unlock() }
func (mw *MutexEnable) noLock()    {}
func (mw *MutexEnable) noUnlock()  {}

func (mw *MutexEnable) Disable() {
	mw.mu.Lock()
	defer mw.mu.Unlock()
	mw.fnLock = mw.noLock
	mw.fnUnlock = mw.noUnlock
}

func (mw *MutexEnable) Enable() {
	mw.mu.Lock()
	defer mw.mu.Unlock()
	mw.fnLock = mw.yesLock
	mw.fnUnlock = mw.yesUnlock
}
