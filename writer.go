package lockenabler

import "io"

func NewWEL(w io.Writer) LockEnableWriter {
	return nil
}

type LockEnableWriter interface {
	LockEnabler
	io.Writer
}

type wel struct {
	mu LockEnabler
	io.Writer
}

func (w *wel) Enable()  { w.mu.Enable() }
func (w *wel) Disable() { w.mu.Disable() }
func (w *wel) Lock()    { w.mu.Lock() }
func (w *wel) Unlock()  { w.mu.Unlock() }

type mutexWrapWriter struct {
	mu *MutexWrap
	io.Writer
}

func (w *mutexWrapWriter) Enable()  { w.mu.Enable() }
func (w *mutexWrapWriter) Disable() { w.mu.Disable() }
func (w *mutexWrapWriter) Lock()    { w.mu.Lock() }
func (w *mutexWrapWriter) Unlock()  { w.mu.Unlock() }

// func (w mutexWrapWriter) Write(p []byte) (n int, err error) { return w.w.Write(p) }

type mutexEnableWriter struct {
	mu *MutexEnable
	io.Writer
}

func (w *mutexEnableWriter) Enable()  { w.mu.Enable() }
func (w *mutexEnableWriter) Disable() { w.mu.Disable() }
func (w *mutexEnableWriter) Lock()    { w.mu.Lock() }
func (w *mutexEnableWriter) Unlock()  { w.mu.Unlock() }

// func (w mutexEnableWriter) Write(p []byte) (n int, err error) { return w.w.Write(p) }
