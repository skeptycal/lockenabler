package lockenabler

import "sync"

type (
	// A Locker represents an object that can be locked
	// or unlocked.
	//
	// Reference: go standard library sync.Locker
	//	 type Locker interface {
	//	 	Lock()
	//	 	Unlock()
	//	 }
	Locker = sync.Locker

	// locker implements sync.Locker and types.Enabler
	// so that the mutex lock can be turned off and on.
	//
	// The default implementation is an unlocked,
	// enabled sync.Mutex with Lock() and Unlock()
	// methods stored in fnLock and fnUnlock,
	// respectively.
	locker struct {
		mu         *sync.Mutex //  mutual exclusion lock.
		fnLock     func()      // Lock()
		fnUnlock   func()      // Unlock()
		lockFunc   func()      // custom Lock()
		unlockFunc func()      // custom Unlock()
		enabler
	}
)

///------> LockerEnabler interface implementation

// Lock locks the underlying mutex by
// calling locker.fnLock()
func (f *locker) Lock() { f.fnLock() }

// Unlock unlocks the underlying mutex by
// calling locker.fnLock()
//
// It is best used with a defer statement
// immediately after calling Lock():
//  lkr.Lock()
// 	defer lkr.Unlock()
func (f *locker) Unlock() { f.fnUnlock() }

///------>  locker default method implementations

// Enable is the default implementation of
// Enabler that enables the lock functionality.
func (f *locker) enable() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.fnLock = f.lockFunc
	f.fnUnlock = f.unlockFunc
}

// Disable disables the lock functionality
// by setting the fnLock and fnUnlock methods
// to point to Nop implementations.
func (f *locker) disable() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.fnLock = f.noLock
	f.fnUnlock = f.noUnlock
}

// yesLock is a default Lock method used when
// Lock is not otherwise implementated.
func (f *locker) yesLock() { f.mu.Lock() }

// yesUnlock is a default Unlock method used when
// Unlock is not otherwise implemented.
func (f *locker) yesUnlock() { f.mu.Unlock() }

// noLock is a default Nop method used when
// Lock is disabled or unavailable in the
// original implementation.
func (*locker) noLock() {}

// noUnlock is a default Nop method used when
// Unlock is disabled or unavailable in the
// original implementation.
func (*locker) noUnlock() {}

///------>  locker optional custom method implementations

// SetLockFuncs allows replacement of the default
// Lock() and Unlock() functions with lockFunc
// and unlockFunc. If either of these functions
// is nil, they are ignored.
func (f *locker) SetLockFuncs(lockFunc, unlockFunc func()) {
	if lockFunc != nil {
		f.lockFunc = lockFunc
	} else {
		f.lockFunc = f.yesLock
	}
	if unlockFunc != nil {
		f.unlockFunc = unlockFunc
	} else {
		f.unlockFunc = f.yesUnlock
	}
}

// func addLockEnabler() LockEnabler {
// 	return AddLockerEnabler(nil, nil)
// }

// Enable enables the lock functionality.
// Any custom enableFunc should begin with:
// 	f.mu.Lock()
// 	defer f.mu.Unlock()
// func (f *locker) Enable() {
// 	if f.enableFunc == nil {
// 		f.enableFunc = f.enable
// 	}
// 	f.enableFunc()
// }

// Disable disables the lock functionality.
// Any custom disableFunc should begin with:
// 	f.mu.Lock()
// 	defer f.mu.Unlock()
// func (f *locker) Disable() {
// 	if f.disableFunc == nil {
// 		f.disableFunc = f.disable
// 	}
// 	f.disableFunc()
// }
