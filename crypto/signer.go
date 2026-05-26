package crypto

import (
	"io"
	"io/fs"
)

// Signer defines a Signer struct.
type Signer struct {
	data   []byte
	sign   []byte
	reader io.Reader
	Error  error
}

// NewSigner returns a new Signer instance.
func NewSigner() Signer {
	_ = "STUB: not implemented"

	// FromString signs from string.
	return *new(Signer)
}

func (s Signer) FromString(str string) Signer { _ = "STUB: not implemented"; return *new(Signer) }

// FromBytes signs from byte slice.
func (s Signer) FromBytes(b []byte) Signer {
	_ = "STUB: not implemented"
	return *

	// FromFile signs from file.
	new(Signer)
}

func (s Signer) FromFile(f fs.File) Signer {
	_ = "STUB: not implemented"
	return *

	// ToRawString outputs as raw string.
	new(Signer)
}

func (s Signer) ToRawString() string { _ = "STUB: not implemented"; return "" }

// ToRawBytes outputs as raw byte slice.
func (s Signer) ToRawBytes() []byte { _ = "STUB: not implemented"; return nil }

// ToBase64String outputs as base64 string.
func (s Signer) ToBase64String() string { _ = "STUB: not implemented"; return "" }

// ToBase64Bytes outputs as base64 byte slice.
func (s Signer) ToBase64Bytes() []byte { _ = "STUB: not implemented"; return nil }

// ToHexString outputs as hex string.
func (s Signer) ToHexString() string { _ = "STUB: not implemented"; return "" }

// ToHexBytes outputs as hex byte slice.
func (s Signer) ToHexBytes() []byte { _ = "STUB: not implemented"; return nil }

func (s Signer) stream(fn func(io.Writer) io.WriteCloser) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to reset the reader position if it's a seeker
