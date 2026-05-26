package sm2

import (
	"io"

	"github.com/dromara/dongle/crypto/keypair"
)

// StdSigner signs data using an SM2 private key.
type StdSigner struct {
	keypair keypair.Sm2KeyPair
	cache   cache
	Error   error
}

// NewStdSigner creates a new SM2 signer bound to the given key pair.
func NewStdSigner(kp *keypair.Sm2KeyPair) *StdSigner { _ = "STUB: not implemented"; return nil }

// Sign generates an SM2 signature for the given data.
func (s *StdSigner) Sign(src []byte) (sign []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamSigner buffers data and writes SM2 signature on Close.
type StreamSigner struct {
	writer  io.Writer
	keypair keypair.Sm2KeyPair
	cache   cache
	buffer  []byte
	Error   error
}

// NewStreamSigner returns a WriteCloser that signs all written data
// with the provided key pair and writes the signature on Close.
func NewStreamSigner(w io.Writer, kp *keypair.Sm2KeyPair) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

// sign generates a signature for the given data.
func (s *StreamSigner) sign(data []byte) (sign []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Write buffers data to be signed.
func (s *StreamSigner) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close signs the buffered data and writes the signature to the
// underlying writer. If the writer implements io.Closer, it is closed.
func (s *StreamSigner) Close() error { _ = "STUB: not implemented"; return nil }

// Sign the buffered data

// Write signature to the underlying writer
