package rsa

import (
	"io"

	"github.com/dromara/dongle/crypto/keypair"
)

type StdVerifier struct {
	keypair keypair.RsaKeyPair // The key pair containing public key and format
	cache   cache              // Cached keys and hash for better performance
	Error   error              // Error field for storing verification errors
}

func NewStdVerifier(kp *keypair.RsaKeyPair) *StdVerifier { _ = "STUB: not implemented"; return nil }

func (v *StdVerifier) Verify(src, sign []byte) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

type StreamVerifier struct {
	keypair   keypair.RsaKeyPair // Key pair containing padding and hash configuration
	cache     cache              // Cached keys and hash for better performance
	reader    io.Reader          // Underlying reader for data input
	signature []byte             // Signature to verify
	verified  bool               // Whether verification has been performed
	Error     error              // Error field for storing verification errors
}

func NewStreamVerifier(r io.Reader, kp *keypair.RsaKeyPair) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func (v *StreamVerifier) verify(hashed, signature []byte) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Write processes data through the hash function for streaming verification
func (v *StreamVerifier) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Process data through the hash function for streaming

// Close performs the final verification and closes the verifier
func (v *StreamVerifier) Close() error { _ = "STUB: not implemented"; return nil }

// Read signature data from the underlying reader

// Get the final hash sum from the hash

// Verify the signature using the hashed data

// Mark verification as completed

// Close the underlying reader if it implements io.Closer
