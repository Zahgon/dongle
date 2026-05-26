package sm2

import (
	"io"

	"github.com/dromara/dongle/crypto/keypair"
)

// StdDecrypter decrypts data using an SM2 private key.
type StdDecrypter struct {
	keypair keypair.Sm2KeyPair
	cache   cache
	Error   error
}

// NewStdDecrypter creates a new SM2 decrypter bound to the given key pair.
func NewStdDecrypter(kp *keypair.Sm2KeyPair) *StdDecrypter { _ = "STUB: not implemented"; return nil }

// Decrypt decrypts data with SM2 private key.
func (d *StdDecrypter) Decrypt(src []byte) (dst []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamDecrypter reads all ciphertext from an io.Reader and exposes the
// decrypted plaintext via Read.
type StreamDecrypter struct {
	reader   io.Reader
	keypair  keypair.Sm2KeyPair
	cache    cache
	buffer   []byte
	position int
	Error    error
}

// NewStreamDecrypter creates a Reader that decrypts the entire input from r
// using the provided key pair, serving plaintext on subsequent Read calls.
func NewStreamDecrypter(r io.Reader, kp *keypair.Sm2KeyPair) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

// decrypt decrypts ciphertext with SM2 private key.
func (d *StreamDecrypter) decrypt(src []byte) (dst []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read serves decrypted plaintext from the internal buffer.
func (d *StreamDecrypter) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Serve from buffer if available

// Otherwise, read all ciphertext and decrypt once

// Return plaintext
