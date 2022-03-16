package lockenabler

import "sync"

/// copy of MutexWrap from logrus for testing purposes

type MutexWrap struct {
	lock     *sync.Mutex
	disabled bool
}

func (mw *MutexWrap) Lock() {
	if !mw.disabled {
		mw.lock.Lock()
	}
}

func (mw *MutexWrap) Unlock() {
	if !mw.disabled {
		mw.lock.Unlock()
	}
}

func (mw *MutexWrap) Disable() {
	mw.disabled = true
}

func (mw *MutexWrap) Enable() {
	mw.disabled = false
}
