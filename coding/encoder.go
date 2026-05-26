package coding

import (
	"io"
	"io/fs"
)

// Encoder defines a Encoder struct.
type Encoder struct {
	src    []byte
	dst    []byte
	reader io.Reader
	Error  error
}

// NewEncoder returns a new Encoder instance.
func NewEncoder() Encoder {
	_ = "STUB: not implemented"

	// FromString encodes from string.
	return *new(Encoder)
}

func (e Encoder) FromString(s string) Encoder { _ = "STUB: not implemented"; return *new(Encoder) }

// FromBytes encodes from byte slice.
func (e Encoder) FromBytes(b []byte) Encoder {
	_ = "STUB: not implemented"
	return *

	// FromFile encodes from file.
	new(Encoder)
}

func (e Encoder) FromFile(f fs.File) Encoder {
	_ = "STUB: not implemented"
	return *

	// ToString outputs as string.
	new(Encoder)
}

func (e Encoder) ToString() string { _ = "STUB: not implemented"; return "" }

// ToBytes outputs as byte slice.
func (e Encoder) ToBytes() []byte { _ = "STUB: not implemented"; return nil }

func (e Encoder) stream(fn func(io.Writer) io.WriteCloser) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to reset the reader position if it's a seeker
