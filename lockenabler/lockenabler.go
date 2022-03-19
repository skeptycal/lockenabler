package lockenabler

type (

// // LockEnabler implements sync.Locker and
// // types.Enabler, basically a Locker that
// // can be enabled or disabled on demand
// // to increase performance when mutex
// // locking is not required.
// //
// // LockEnabler implements SetLockFuncs and
// // SetEnableFuncs as a way to set custom lock,
// // unlock, enable, and disable functions.
// LockEnabler interface {
// 	locker.Locker
// 	enabler.Enabler
// 	SetLockFuncs(lockFunc, unlockFunc func())
// 	SetEnableFuncs(enableFunc, disableFunc func())
// }

// lockEnabler struct {
// 	*locker
// }

)

///------> LockerEnabler constructor

// // NewLockerEnabler implements sync.Locker by using
// // the given Lock and Unlock methods. If
// // either of these is nil, then the default
// // implementation, a Nop, is used.
// //
// // This may be used to add Locker functionality
// // to structures that do not implement the
// // interface natively.
// //
// // The default implementation is an unlocked,
// // enabled sync.Mutex with Lock() and Unlock()
// // pointing to its underlying methods of the
// // same names.
// func NewLockEnabler() LockEnabler {
// 	f := &locker.locker{
// 		mu: new(sync.Mutex),
// 	}
// 	f.SetLockFuncs(nil, nil)
// 	f.SetEnableFuncs(nil, nil)
// 	f.Enable()
// 	return f
// }
