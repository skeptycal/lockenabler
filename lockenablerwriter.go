package lockenabler

type (
	// Writer implements io.Writer
	Writer interface {
		Write(p []byte) (n int, err error)
	}

	// StringWriter implements io.StringWriter
	StringWriter interface {
		WriteString(string) (n int, err error)
	}

	LockEnableWriter interface {
		LockEnabler
		Writer
	}
)
