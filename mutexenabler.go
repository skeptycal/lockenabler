package lockenabler

import "sync"

/// modified logrus.MutexWrap for testing purposes
// removed boolean flag and replaced with different functions

type MutexEnable struct {
	lock     *sync.Mutex
	fnLock   func()
	fnUnlock func()
}

func (mw *MutexEnable) Lock()      { mw.fnLock() }
func (mw *MutexEnable) Unlock()    { mw.fnUnlock() }
func (mw *MutexEnable) yesLock()   { mw.lock.Lock() }
func (mw *MutexEnable) yesUnlock() { mw.lock.Unlock() }
func (mw *MutexEnable) noLock()    {}
func (mw *MutexEnable) noUnlock()  {}

func (mw *MutexEnable) Disable() {
	mw.fnLock = mw.noLock
	mw.fnUnlock = mw.noUnlock
}

func (mw *MutexEnable) Enable() {
	mw.fnLock = mw.yesLock
	mw.fnUnlock = mw.yesUnlock
}
