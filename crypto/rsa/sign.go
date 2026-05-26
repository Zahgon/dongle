package rsa

import (
	"io"

	"github.com/dromara/dongle/crypto/keypair"
)

type StdSigner struct {
	keypair keypair.RsaKeyPair // The key pair containing private key and format
	cache   cache              // Cached keys and hash for better performance
	Error   error              // Error field for storing signature errors
}

func NewStdSigner(kp *keypair.RsaKeyPair) *StdSigner { _ = "STUB: not implemented"; return nil }

func (s *StdSigner) Sign(src []byte) (sign []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type StreamSigner struct {
	keypair keypair.RsaKeyPair // Key pair containing padding and hash configuration
	cache   cache              // Cached keys and hash for better performance
	writer  io.Writer          // Underlying writer for signature output
	Error   error              // Error field for storing signature errors
}

func NewStreamSigner(w io.Writer, kp *keypair.RsaKeyPair) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func (s *StreamSigner) sign(data []byte) (dst []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StreamSigner) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *StreamSigner) Close() error { _ = "STUB: not implemented"; return nil }

// Get the final hash sum from the hash

// Generate signature for the hashed data

// Write signature to the underlying writer

// Close the underlying writer if it implements io.Closer
