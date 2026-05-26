package sm2

import (
	"io"

	"github.com/dromara/dongle/crypto/keypair"
)

// StdEncrypter encrypts data using an SM2 public key.
// The ciphertext component order is derived from Sm2KeyPair.Order.
type StdEncrypter struct {
	keypair keypair.Sm2KeyPair
	cache   cache
	Error   error
}

// NewStdEncrypter creates a new SM2 encrypter bound to the given key pair.
func NewStdEncrypter(kp *keypair.Sm2KeyPair) *StdEncrypter { _ = "STUB: not implemented"; return nil }

// Encrypt encrypts data with SM2 public key.
func (e *StdEncrypter) Encrypt(src []byte) (dst []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamEncrypter buffers plaintext and writes SM2 ciphertext on Close.
type StreamEncrypter struct {
	writer  io.Writer
	keypair keypair.Sm2KeyPair
	cache   cache
	buffer  []byte
	Error   error
}

// NewStreamEncrypter returns a WriteCloser that encrypts all written data
// with the provided key pair and writes the ciphertext on Close.
func NewStreamEncrypter(w io.Writer, kp *keypair.Sm2KeyPair) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

// encrypt encrypts plaintext with SM2 public key.
func (e *StreamEncrypter) encrypt(src []byte) (dst []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Write buffers plaintext to be encrypted.
func (e *StreamEncrypter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close encrypts the buffered plaintext and writes the ciphertext to the
// underlying writer. If the writer implements io.Closer, it is closed.
func (e *StreamEncrypter) Close() error { _ = "STUB: not implemented"; return nil }
