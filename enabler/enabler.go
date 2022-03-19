package enabler

// AddEnabler implements types.Enabler by using
// the given Enable and Disa ble methods. If
// either of these is nil, then the default
// implementation, a Nop, is used.
//
// This may be used to add Enabler functionality
// to structures that do not implement the
// interface natively. For example:
/*
	type MyMutex interface {
		sync.Locker
		Enabler
	}

	type myMutex struct {
		mu *sync.Mutex
		Enabler
	}

	func (m *MyMutex) fnEnable() {
		m.mu.Lock()
		defer m.mu.Unlock()
		m.fnLock = m.lockFunc
		m.fnUnlock = m.unlockFunc
	}

	func (m *MyMutex) fnDisable() {
		m.mu.Lock()
		defer m.mu.Unlock()
		m.fnLock = m.noLock
		m.fnUnlock = m.noUnlock
	}

	var MyMutex myMutex = myMutex{
			mu new(sync.Mutex)
			lockenabler.AddEnabler(fnEnable,fnDisable)
	}
*/
func AddEnabler(fnEnable, fnDisable func()) Enabler {
	return newenabler(fnEnable, fnDisable)
}

type (

	// An Enabler represents an object that can be enabled
	// or disabled. It is copied here from the types package
	// to prevent a circular dependency.
	//
	// Reference: http://github.com/skeptycal/types
	Enabler interface {
		Enable()
		Disable()
		SetEnableFuncs(enableFunc, disableFunc func())
	}

	// enabler implements Enabler so that major a
	// functionality can be turned on and off. The
	// default implementation is an "enabled"
	// Enabler with no-op Enable() and Disable()
	// methods stored in fnEnable and fnDisable,
	// respectively.
	//
	// Custom enable and disable functions may be
	// registered with SetEnableFuncs().
	enabler struct {
		fnEnable  func() // current Enable()
		fnDisable func() // current Disable()

		// customEnable is a ... surprise ... custom
		// Enable() function. It is set with
		// SetEnableFuncs(). If set to nil, the
		// default no-op enable function is set.
		customEnable func()

		// customDisable is a ... surprise ... custom
		// Disable() function. It is set with
		// SetEnableFuncs(). If set to nil, the
		// default no-op disable function is set.
		customDisable func() // custom Disable()
	}
)

// newFakeEnabler returns a *fakeEnabler that
// implements Enabler by using the given Enable
// and Disable methods.
//
// If either of these is nil, then
// the default implementation, a Nop, is used.
func newenabler(fnEnable, fnDisable func()) *enabler {
	f := new(enabler)
	f.SetEnableFuncs(fnEnable, fnDisable)
	f.Enable()
	return f
}

// Enable enables the underlying feature and
// may be used with a defer if practical
// immediately before a call to Disable(), e.g.
/*

	func main() {
		var MyThing = AddEnabler(nil,nil)
		fmt.Println(MyThing)
	}

	func DoStuff(my Enabler) error {
		my.Enable()
		defer my.Disable()

		err = doTheThing()
		if err != nil {
			return err
		}
		// process stuff ...

	 	return nil
	}
*/
func (f *enabler) Enable() {
	// TODO: check this ...
	// f.fnEnable should never be nil due to
	// the constructor and SetEnableFuncs
	// method disllowing it?
	// if f.fnEnable == nil {
	// 	f.fnEnable = f.noEnable
	// }
	f.fnEnable()
}

// Disable disables the underlying feature and
// may be used with a defer if practical
// immediately after a call to Enable(), e.g.
/*

	func main() {
		var MyThing = AddEnabler(nil,nil)
		fmt.Println(MyThing)
	}

	func DoStuff(my Enabler) error {
		my.Enable()
		defer my.Disable()

		err = doTheThing()
		if err != nil {
			return err
		}
		// process stuff ...

	 	return nil
	}
*/
func (f *enabler) Disable() {
	// TODO: check this ...
	// f.fnEnable should never be nil due to
	// the constructor and SetEnableFuncs
	// method disllowing it?
	// if f.fnDisable == nil {
	// 	f.fnDisable = f.noDisable
	// }
	f.fnDisable()
}

// noEnable is a default Nop method used when
// either Enable or Disabler is unavailable
// in the original implementation.
func (*enabler) noEnable() {}

// noDisable is a default Nop method used when
// either Enable or Disabler is unavailable
// in the original implementation.
func (*enabler) noDisable() {}

// SetEnableFuncs allows replacement of the default
// Enable() and Disable() functions with enableFunc
// and disableFunc. If either of these functions
// is nil, the default implementation is used.
func (f *enabler) SetEnableFuncs(enableFunc, disableFunc func()) {
	if enableFunc != nil {
		f.customEnable = enableFunc
	} else {
		f.customEnable = f.noEnable
	}

	if disableFunc != nil {
		f.customDisable = disableFunc
	} else {
		f.customDisable = f.noDisable
	}
}
