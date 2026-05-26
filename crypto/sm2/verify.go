package sm2

import (
	"io"

	"github.com/dromara/dongle/crypto/keypair"
)

// StdVerifier verifies data using an SM2 public key.
type StdVerifier struct {
	keypair keypair.Sm2KeyPair
	cache   cache
	Error   error
}

// NewStdVerifier creates a new SM2 verifier bound to the given key pair.
func NewStdVerifier(kp *keypair.Sm2KeyPair) *StdVerifier { _ = "STUB: not implemented"; return nil }

// Verify verifies an SM2 signature for the given data.
func (v *StdVerifier) Verify(src, sign []byte) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Verify the signature (Verify internally calculates ZA and digest)

// StreamVerifier reads signature from an io.Reader and verifies data written to it.
type StreamVerifier struct {
	reader    io.Reader
	keypair   keypair.Sm2KeyPair
	cache     cache
	buffer    []byte
	signature []byte
	verified  bool
	Error     error
}

// NewStreamVerifier creates a WriteCloser that verifies data written to it
// using the signature read from the provided reader.
func NewStreamVerifier(r io.Reader, kp *keypair.Sm2KeyPair) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

// Parse and cache the public key for reuse

// verify verifies the signature for the given data.
func (v *StreamVerifier) verify(data, signature []byte) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Verify the signature (Verify internally calculates ZA and digest)

// Write buffers data for verification.
func (v *StreamVerifier) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close reads the signature from the underlying reader and performs verification.
func (v *StreamVerifier) Close() error { _ = "STUB: not implemented"; return nil }

// Read signature data from the underlying reader

// Verify the signature using the buffered data

// Close the underlying reader if it implements io.Closer
