package lockenabler

type NopWriter struct{}

// nopWriter_Write returns 0, nil no matter what the input is.
func (NopWriter) Write(b []byte) (n int, err error) {
	return 0, nil
}

type LenWriter struct{}

// nopWriter_Write returns 0, nil no matter what the input is.
func (LenWriter) Write(b []byte) (n int, err error) {
	return len(b), nil
}
