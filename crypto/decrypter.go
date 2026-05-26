package crypto

import (
	"io"
	"io/fs"
)

// Decrypter defines a Decrypter struct.
type Decrypter struct {
	src    []byte
	dst    []byte
	reader io.Reader
	Error  error
}

// NewDecrypter returns a new Decrypter instance.
func NewDecrypter() Decrypter {
	_ = "STUB: not implemented"

	// FromRawString decrypts from raw string.
	return *new(Decrypter)
}

func (d Decrypter) FromRawString(s string) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// FromRawBytes decrypts from raw bytes.
func (d Decrypter) FromRawBytes(b []byte) Decrypter {
	_ = "STUB: not implemented"
	return *

	// FromRawFile decrypts from raw file.
	new(Decrypter)
}

func (d Decrypter) FromRawFile(f fs.File) Decrypter {
	_ = "STUB: not implemented"
	return *

	// FromBase64String decrypts from base64 string.
	new(Decrypter)
}

func (d Decrypter) FromBase64String(s string) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// FromBase64Bytes decrypts from base64 bytes.
func (d Decrypter) FromBase64Bytes(b []byte) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// FromBase64File decrypts from base64 file.
func (d Decrypter) FromBase64File(f fs.File) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// FromHexString decrypts from hex string.
func (d Decrypter) FromHexString(s string) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// FromHexBytes decrypts from hex bytes.
func (d Decrypter) FromHexBytes(b []byte) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// FromHexFile decrypts from hex file.
func (d Decrypter) FromHexFile(f fs.File) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// ToString outputs as string.
func (d Decrypter) ToString() string { _ = "STUB: not implemented"; return "" }

// ToBytes outputs as byte slice.
func (d Decrypter) ToBytes() []byte { _ = "STUB: not implemented"; return nil }

func (d Decrypter) stream(fn func(io.Reader) io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to reset the reader position if it's a seeker
