package lockenabler

type LockEnabler interface {
	Locker
	Enabler
}

func NewLockEnabler(en *fakeEnabler, lk *fakeLocker) LockEnabler { return &lockEnabler{en, lk} }

type lockEnabler struct {
	*fakeEnabler
	*fakeLocker
}

// func (f *lockEnabler) Lock()   { f.fnLock() }
// func (f *lockEnabler) Unlock() { f.fnUnlock() }

// func (*lockEnabler) noEnable()  {}
// func (*lockEnabler) noDisable() {}
// func (*lockEnabler) noLock()    {}
// func (*lockEnabler) noUnlock()  {}
