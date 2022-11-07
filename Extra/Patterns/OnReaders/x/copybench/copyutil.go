package copybench

import (
	"crypto/rand"
	"io"
)

type randFile struct {
	Size int
	I    int
	r    io.Reader
}

func RandFile(size int) *randFile {
	return &randFile{
		Size: size,
		I:    0,
		r:    rand.Reader,
	}
}

// Read reads random bytes.
func (f *randFile) Read(p []byte) (int, error) {
	if f.I >= f.Size {
		return 0, io.EOF
	}
	gap := f.Size - f.I
	if gap < len(p) {
		f.I += gap
		return f.r.Read(p[:gap])
	}
	f.I += len(p)
	return f.r.Read(p)
}

// emptyFile returns zero file of a given size.
type emptyFile struct {
	Size  int
	I     int
	Zeros []byte
}

func EmptyFile(size int) *emptyFile {
	return &emptyFile{
		Size:  size,
		I:     0,
		Zeros: make([]byte, 16777216),
	}
}

// Read fills p with zeros.
func (f *emptyFile) Read(p []byte) (int, error) {
	if f.I >= f.Size {
		return 0, io.EOF
	}
	f.I += len(p)
	copy(p, f.Zeros)
	return len(p), nil
}
