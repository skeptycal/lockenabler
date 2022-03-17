package lockenabler

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type LockEnabler interface {
	Locker
	Enabler
}

func AddLockEnabler() LockEnabler {
	return fakeLockEnabler()
}

func fakeLockEnabler() *lockEnabler {
	le := lockEnabler{}
	le.fakeEnabler = newFakeEnabler(nil, nil)
	le.fakeLocker = newFakeLocker(nil, nil)
	return &le
}

func NewLockEnabler(en *fakeEnabler, lk *locker) LockEnabler {
	return &lockEnabler{en, lk}
}

type lockEnabler struct {
	*fakeEnabler
	*locker
}
