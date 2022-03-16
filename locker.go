package lockenabler

// A Locker represents an object that can be locked
// or unlocked. It is copied here from the sync package.
//
// Reference: go standard library sync.Locker
type Locker interface {
	Lock()
	Unlock()
}

type fakeLocker struct {
	fnLock   func()
	fnUnlock func()
}

// AddLocker implements sync.Locker by using
// the given Lock and Unlock methods. If
// either of these is nil, then the default
// implementation, a Nop, is used.
// This may be used to add Locker functionality
// to structures that do not implement the
// interface natively.
func AddLocker(fnLock, fnUnlock func()) Locker {
	return newFakeLocker(fnLock, fnUnlock)
}

// newFakeLocker implements sync.Locker by using
// the given Lock and Unlock methods. If either
// of these is nil, then the default
// implementation, a Nop, is used.
func newFakeLocker(fnLock, fnUnlock func()) *fakeLocker {
	f := fakeLocker{fnLock, fnUnlock}
	if fnLock == nil {
		f.fnLock = f.noLock
	}
	if fnUnlock == nil {
		f.fnUnlock = f.noUnlock
	}
	return &f
}

// Lock locks the underlying mutex.
func (f *fakeLocker) Lock() { f.fnLock() }

// Unlock unlocks the underlying mutex and
// is best used with a defer statement
// immediately after calling Lock()
func (f *fakeLocker) Unlock() { f.fnUnlock() }

// noLock is a default Nop method used when
// Lock is unavailable in the original
// implementation.
func (*fakeLocker) noLock() {}

// noLock is a default Nop method used when
// Unlock is unavailable in the original
// implementation.
func (*fakeLocker) noUnlock() {}
