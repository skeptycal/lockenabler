package lockenabler

import (
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"testing"
)

var writerTestList = []struct {
	name string
	w    io.Writer
}{
	{"mutexWrapWriter", mutexWrapWriter{&MutexWrap{}, NopWriter{}}},
	{"mutexEnableWriter", mutexEnableWriter{&MutexEnable{}, NopWriter{}}},
	{"nopWriter", NopWriter{}},
	{"lenWriter", LenWriter{}},
	// {"os.Stderr", os.Stderr},
}

func BenchmarkWriters(b *testing.B) {
	for _, bb := range writerTestList {
		for i := 0; i < 8; i++ {
			name := fmt.Sprintf("%v (size: %d)", bb.name, i)
			crazyWriterLoop(b, name, bb.w, 1<<i)
		}
	}
}

func flip() bool { return mathrand.Intn(10000)&1 == 1 }

func crazyWriterLoop(b *testing.B, name string, w WriteEnableLocker, size int) {

	var loopsize = 4
	r := rand.Reader
	buf := make([]byte, 0, size)

	// do a lot of time wasting reading and writing ...
	b.Run(name, func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			for j := 0; j < loopsize; j++ {
				// enable and disable randomly and often ...
				if flip() {
					w.Enable()
				} else {
					w.Disable()
				}

				// lock and unlock if available
				b.Log("lock writer")
				w.Lock()
				defer w.Unlock()

				r.Read(buf)
				n, err := w.Write(buf)

				if err != nil {
					b.Logf("write failed (%v bytes): %v", n, err)
				}
			}
		}
	})
}
