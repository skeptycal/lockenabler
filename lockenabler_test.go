package lockenabler

import (
	"sync"
)

// fakeLocker implements sync.Locker by using
// the given Lock and Unlock methods. If either
// of these is nil, then the default
// implementation, a Nop, is used.
func fakeLocker(lockFunc, unlockFunc func()) *locker {
	f := &locker{mu: new(sync.Mutex)}
	f.SetLockFuncs(lockFunc, unlockFunc)
	f.Enable()
	return f
}
