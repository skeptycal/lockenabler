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
	w    LockEnableWriter
}{
	{"mutexWrapWriter", NewWrapWriter(io.Discard)},
	{"mutexEnablerWriter", NewLockEnableWriter(io.Discard)},

	{"nopWriter", nopWriter},
	{"lenWriter", lenWriter},
	// {"os.Stderr", os.Stderr},
}

func BenchmarkWriters(b *testing.B) {
	for _, bb := range writerTestList {
		for i := 2; i < 8; i++ {
			name := fmt.Sprintf("%v (size: %d)", bb.name, 1<<i*8)
			crazyWriterLoop(b, name, bb.w, 1<<i*8)
		}
	}
}

func flip() bool { return mathrand.Intn(10000)&1 == 1 }

func crazyWriterLoop(b *testing.B, name string, w LockEnableWriter, size int) {

	var loopsize = 4
	r := rand.Reader
	buf := make([]byte, 0, size)

	// do a lot of time wasting reading and writing ...
	b.Run(name, func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			for j := 0; j < loopsize; j++ {
				// // enable and disable randomly and often ...
				// if flip() {
				// 	// b.Log("enable writer")
				// 	w.Enable()
				// } else {
				// 	// b.Log("disable writer")
				// 	w.Disable()
				// }

				// lock and unlock if available
				// b.Log("lock writer")
				// w.Lock()
				// defer w.Unlock()

				r.Read(buf)
				n, err := w.Write(buf)

				if err != nil {
					b.Logf("write failed (%v bytes): %v", n, err)
				}
			}
		}
	})
}

/// Benchmark Results:
/*
mutexWrapWriter_(size:_0)-8         	 9644938	       112.9 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_1)-8         	14368303	       106.4 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_2)-8         	10600101	       113.0 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_3)-8         	10650193	       113.0 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_4)-8         	10630678	       112.8 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_5)-8         	16610049	       109.3 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_6)-8         	16565071	       109.5 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_7)-8         	10596236	       112.9 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_0)-8       	10591209	       113.5 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_1)-8       	10592059	       113.2 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_2)-8       	10625377	       114.1 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_3)-8       	10543274	       113.4 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_4)-8       	10578745	       113.3 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_5)-8       	10628400	       113.5 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_6)-8       	10706581	       113.3 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_7)-8       	10597702	       113.5 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_0)-8               	 8328190	       144.5 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_1)-8               	 8358835	       144.0 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_2)-8               	 8343956	       144.1 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_3)-8               	 8327648	       144.0 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_4)-8               	 8307934	       144.5 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_5)-8               	 8300568	       143.9 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_6)-8               	 8347351	       144.1 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_7)-8               	 8358788	       143.9 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_0)-8               	 8326962	       143.8 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_1)-8               	17072320	       114.0 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_2)-8               	 8342193	       143.6 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_3)-8               	17011412	       136.7 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_4)-8               	 8359046	       143.6 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_5)-8               	 8352055	       143.6 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_6)-8               	16914342	       111.9 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_7)-8               	17054498	       127.7 ns/op	       0 B/op	       0 allocs/op

* after adding nilcheck to fnEnable, fnLock, etc...
mutexWrapWriter_(size:_2)-8         	 8103253	       138.4 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_3)-8         	16580845	       105.9 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_4)-8         	 9954454	       116.5 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_5)-8         	 8425255	       141.9 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_6)-8         	 8445932	       141.8 ns/op	       0 B/op	       0 allocs/op
mutexWrapWriter_(size:_7)-8         	 8463759	       142.1 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_2)-8       	 8436583	       142.1 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_3)-8       	 8432134	       142.5 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_4)-8       	 8411516	       142.2 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_5)-8       	 8410075	       142.0 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_6)-8       	 8428946	       142.3 ns/op	       0 B/op	       0 allocs/op
mutexEnableWriter_(size:_7)-8       	 8431159	       142.2 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_2)-8               	 8464684	       142.1 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_3)-8               	16414876	       135.5 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_4)-8               	 8443533	       142.2 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_5)-8               	 8424486	       145.1 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_6)-8               	 8426448	       142.4 ns/op	       0 B/op	       0 allocs/op
nopWriter_(size:_7)-8               	 8433747	       142.5 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_2)-8               	 8437268	       142.4 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_3)-8               	 8422509	       142.4 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_4)-8               	 8429636	       142.4 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_5)-8               	 8420869	       142.4 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_6)-8               	 8452020	       142.4 ns/op	       0 B/op	       0 allocs/op
lenWriter_(size:_7)-8               	 8415997	       142.3 ns/op	       0 B/op	       0 allocs/op
*/

/// n > 100000 seems to keep the variance within 1% regularly
func Test_flip(t *testing.T) {

	trials := 10 // number of trials
	for j := 0; j < trials; j++ {
		t.Run("coin flip test", func(t *testing.T) {
			var sum float64 = 0

			// n is number of coin flips
			var n float64 = 100000
			for i := 0; i < int(n); i++ {
				if flip() {
					sum++
				}
			}
			avg := sum / n
			// if avg is + or - 1% from .50 ... trigger an error
			if avg < .49 || avg > .51 {
				t.Errorf("flip() avg (n=%v): %v", n, avg)
			}
		})

	}
}
