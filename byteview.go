package GeeCache

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return v.clonebytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func (v ByteView) clonebytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
