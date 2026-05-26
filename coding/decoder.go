package coding

import (
	"io"
	"io/fs"
)

// Decoder defines a Decoder struct.
type Decoder struct {
	src    []byte
	dst    []byte
	reader io.Reader
	Error  error
}

// NewDecoder returns a new Decoder instance.
func NewDecoder() Decoder {
	_ = "STUB: not implemented"

	// FromString decodes from string.
	return *new(Decoder)
}

func (d Decoder) FromString(s string) Decoder { _ = "STUB: not implemented"; return *new(Decoder) }

// FromBytes decodes from byte slice.
func (d Decoder) FromBytes(b []byte) Decoder {
	_ = "STUB: not implemented"
	return *

	// FromFile decodes from file.
	new(Decoder)
}

func (d Decoder) FromFile(f fs.File) Decoder {
	_ = "STUB: not implemented"
	return *

	// ToString outputs as string.
	new(Decoder)
}

func (d Decoder) ToString() string { _ = "STUB: not implemented"; return "" }

// ToBytes outputs as byte slice.
func (d Decoder) ToBytes() []byte { _ = "STUB: not implemented"; return nil }

func (d Decoder) stream(fn func(io.Reader) io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to reset the reader position if it's a seeker
