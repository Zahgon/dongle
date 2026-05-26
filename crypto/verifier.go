package crypto

import (
	"io"
	"io/fs"
)

// Verifier defines a Verifier struct.
type Verifier struct {
	data   []byte
	sign   []byte
	verify bool
	reader io.Reader
	Error  error
}

// NewVerifier returns a new Verifier instance.
func NewVerifier() Verifier {
	_ = "STUB: not implemented"

	// FromString verifies from string.
	return *new(Verifier)
}

func (v Verifier) FromString(s string) Verifier { _ = "STUB: not implemented"; return *new(Verifier) }

// FromBytes verifies from byte slice.
func (v Verifier) FromBytes(b []byte) Verifier {
	_ = "STUB: not implemented"
	return *

	// FromFile verifies from file.
	new(Verifier)
}

func (v Verifier) FromFile(f fs.File) Verifier {
	_ = "STUB: not implemented"
	return *

	// WithHexSign verifies with hex sign.
	new(Verifier)
}

func (v Verifier) WithHexSign(s []byte) Verifier { _ = "STUB: not implemented"; return *new(Verifier) }

// WithBase64Sign verifies with base64 sign.
func (v Verifier) WithBase64Sign(s []byte) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}

// WithRawSign verifies with raw sign.
func (v Verifier) WithRawSign(s []byte) Verifier {
	_ = "STUB: not implemented"
	return *

	// ToBool returns true if verification is successful.
	new(Verifier)
}

func (v Verifier) ToBool() bool { _ = "STUB: not implemented"; return false }

func (v Verifier) stream(fn func(io.Writer) io.WriteCloser) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to reset the reader position if it's a seeker
