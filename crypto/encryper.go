package crypto

import (
	"io"
	"io/fs"
)

// Encrypter defines a Encrypter struct.
type Encrypter struct {
	src    []byte
	dst    []byte
	reader io.Reader
	Error  error
}

// NewEncrypter returns a new Encrypter instance.
func NewEncrypter() Encrypter {
	_ = "STUB: not implemented"

	// FromString encrypts from string.
	return *new(Encrypter)
}

func (e Encrypter) FromString(s string) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// FromBytes encrypts from byte slice.
func (e Encrypter) FromBytes(b []byte) Encrypter {
	_ = "STUB: not implemented"
	return *

	// FromFile encrypts from file.
	new(Encrypter)
}

func (e Encrypter) FromFile(f fs.File) Encrypter {
	_ = "STUB: not implemented"
	return *

	// ToRawString outputs as raw string.
	new(Encrypter)
}

func (e Encrypter) ToRawString() string { _ = "STUB: not implemented"; return "" }

// ToRawBytes outputs as raw byte slice.
func (e Encrypter) ToRawBytes() []byte { _ = "STUB: not implemented"; return nil }

// ToBase64String outputs as base64 string.
func (e Encrypter) ToBase64String() string { _ = "STUB: not implemented"; return "" }

// ToBase64Bytes outputs as base64 byte slice.
func (e Encrypter) ToBase64Bytes() []byte { _ = "STUB: not implemented"; return nil }

// ToHexString outputs as hex string.
func (e Encrypter) ToHexString() string { _ = "STUB: not implemented"; return "" }

// ToHexBytes outputs as hex byte slice.
func (e Encrypter) ToHexBytes() []byte { _ = "STUB: not implemented"; return nil }

func (e Encrypter) stream(fn func(io.Writer) io.WriteCloser) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to reset the reader position if it's a seeker
