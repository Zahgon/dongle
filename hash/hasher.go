// Package hash provides cryptographic hash and hmac functions.
// It supports multiple hash algorithms including MD2, MD4, MD5, SHA1, SHA2, SHA3,
// BLAKE2b, BLAKE2s, RIPEMD160, SM3 and so on, with both standard and streaming modes.
package hash

import (
	"hash"
	"io"
	"io/fs"
)

// BufferSize buffer size for streaming (64KB is a good balance)
var BufferSize = 64 * 1024

// Hasher defines a Hasher struct.
type Hasher struct {
	src    []byte
	dst    []byte
	key    []byte
	reader io.Reader
	Error  error
}

// NewHasher returns a new Hasher instance.
func NewHasher() Hasher {
	_ = "STUB: not implemented"

	// FromString encrypts from string.
	return *new(Hasher)
}

func (h Hasher) FromString(s string) Hasher { _ = "STUB: not implemented"; return *new(Hasher) }

// FromBytes encrypts from byte slice.
func (h Hasher) FromBytes(b []byte) Hasher {
	_ = "STUB: not implemented"
	return *

	// FromFile encrypts from file.
	new(Hasher)
}

func (h Hasher) FromFile(f fs.File) Hasher {
	_ = "STUB: not implemented"
	return *

	// WithKey sets the key for HMAC calculation from byte slice.
	new(Hasher)
}

func (h Hasher) WithKey(key []byte) Hasher { _ = "STUB: not implemented"; return *new(Hasher) }

// ToRawString outputs as raw string without encoding.
func (h Hasher) ToRawString() string { _ = "STUB: not implemented"; return "" }

// ToRawBytes outputs as raw byte slice without encoding.
func (h Hasher) ToRawBytes() []byte { _ = "STUB: not implemented"; return nil }

// ToBase64String outputs as base64 string.
func (h Hasher) ToBase64String() string { _ = "STUB: not implemented"; return "" }

// ToBase64Bytes outputs as base64 byte slice.
func (h Hasher) ToBase64Bytes() []byte { _ = "STUB: not implemented"; return nil }

// ToHexString outputs as hex string.
func (h Hasher) ToHexString() string { _ = "STUB: not implemented"; return "" }

// ToHexBytes outputs as hex byte slice.
func (h Hasher) ToHexBytes() []byte { _ = "STUB: not implemented"; return nil }

func (h Hasher) stream(fn func() hash.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to reset the reader position if it's a seeker

func (h Hasher) hmac(fn func() hash.Hash) Hasher { _ = "STUB: not implemented"; return *new(Hasher) }

// Streaming mode

// Try to reset the reader position if it's a seeker

// Standard mode
