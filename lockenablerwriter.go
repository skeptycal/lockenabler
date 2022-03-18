package lockenabler

type (
	LockEnableWriter interface {
		LockEnabler
		ioWriter
	}

	// ioWriter implements io.ioWriter
	ioWriter interface {
		Write(p []byte) (n int, err error)
	}

	// ioStringWriter implements io.ioStringWriter
	ioStringWriter interface {
		WriteString(string) (n int, err error)
	}
)
