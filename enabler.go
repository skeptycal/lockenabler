package lockenabler

// An Enabler represents an object that can be enabled
// or disabled. It is copied here from the types package
// to prevent a circular dependency.
//
// Reference: http://github.com/skeptycal/types
type Enabler interface {
	Enable()
	Disable()
}

type fakeEnabler struct {
	fnEnable  func()
	fnDisable func()
}

// AddEnabler implements types.Enabler by using
// the given Enable and Disa ble methods. If
// either of these is nil, then the default
// implementation, a Nop, is used.
//
// This may be used to add Enabler functionality
// to structures that do not implement the
// interface natively.
func AddEnabler(fnEnable, fnDisable func()) Enabler {
	return newFakeEnabler(fnEnable, fnDisable)
}

// newFakeEnabler implements types.Enabler
// by using the given Enable and Disable
// methods. If either of these is nil, then
// the default implementation, a Nop, is used.
func newFakeEnabler(fnEnable, fnDisable func()) *fakeEnabler {
	f := fakeEnabler{fnEnable, fnDisable}
	if fnEnable == nil {
		f.fnEnable = f.noEnable
	}
	if fnDisable == nil {
		f.fnDisable = f.noDisable
	}
	return &f
}

// Enable enables the underlying feature.
func (f *fakeEnabler) Enable() { f.fnEnable() }

// Disable disables the underlying feature and
// may be used with a defer statement
// immediately after a call to Enable()
func (f *fakeEnabler) Disable() { f.fnDisable() }

// noEnable is a default Nop method used when
// either Enable or Disabler is unavailable
// in the original implementation.
func (*fakeEnabler) noEnable() {}

// noDisable is a default Nop method used when
// either Enable or Disabler is unavailable
// in the original implementation.
func (*fakeEnabler) noDisable() {}
