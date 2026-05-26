package ed25519

import (
	"io"

	"github.com/dromara/dongle/crypto/keypair"
)

// StdVerifier represents a standard ED25519 verifier.
type StdVerifier struct {
	keypair keypair.Ed25519KeyPair
	cache   cache // Cached keys for better performance
	Error   error // Error field for storing verification errors
}

// NewStdVerifier creates a new standard ED25519 verifier.
func NewStdVerifier(kp *keypair.Ed25519KeyPair) *StdVerifier { _ = "STUB: not implemented"; return nil }

// Verify verifies the signature for the given data using the ED25519 public key.
func (v *StdVerifier) Verify(src, sign []byte) (valid bool, err error) {
	_ = "STUB: not implemented"
	// Check for existing errors from initialization
	return false, nil
}

// ED25519 verification does not require hashing as it handles hashing internally

// StreamVerifier represents a streaming ED25519 verifier that processes data in chunks.
type StreamVerifier struct {
	keypair   keypair.Ed25519KeyPair // Key pair containing public key
	cache     cache                  // Cached keys for better performance
	reader    io.Reader              // Underlying reader for signature input
	buffer    []byte                 // Buffer to accumulate data for verification
	signature []byte                 // Signature to verify
	verified  bool                   // Whether verification has been performed
	Error     error                  // Error field for storing verification errors
}

// NewStreamVerifier creates a new streaming ED25519 verifier.
func NewStreamVerifier(r io.Reader, kp *keypair.Ed25519KeyPair) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

// verify verifies the signature for the given data.
func (v *StreamVerifier) verify(data, sign []byte) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Write accumulates data for verification.
func (v *StreamVerifier) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Append data to buffer

// Close performs the final verification.
func (v *StreamVerifier) Close() error { _ = "STUB: not implemented"; return nil }

// Read signature data from the underlying reader

// Verify the signature using the accumulated data

// Close the underlying reader if it implements io.Closer
