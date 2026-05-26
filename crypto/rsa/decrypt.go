package rsa

import (
	"io"

	"github.com/dromara/dongle/crypto/keypair"
)

type StdDecrypter struct {
	keypair keypair.RsaKeyPair // The key pair containing private key and format
	cache   cache              // Cached keys and hash for better performance
	Error   error              // Error field for storing decryption errors
}

func NewStdDecrypter(kp *keypair.RsaKeyPair) *StdDecrypter { _ = "STUB: not implemented"; return nil }

func (d *StdDecrypter) Decrypt(src []byte) (dst []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type StreamDecrypter struct {
	keypair  keypair.RsaKeyPair // Key pair containing padding and hash configuration
	cache    cache              // Cached keys and hash for better performance
	reader   io.Reader          // Underlying reader for encrypted input
	buffer   []byte             // Buffer for decrypted data
	position int                // Current position in buffer
	Error    error              // Error field for storing decryption errors
}

func NewStreamDecrypter(r io.Reader, kp *keypair.RsaKeyPair) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (d *StreamDecrypter) decrypt(data []byte) (dst []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *StreamDecrypter) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If we have decrypted data available, return it

// If we've exhausted all decrypted data, try to read more

// Determine block size based on key type

// Read one encrypted block from the underlying reader

// Note: io.ReadFull guarantees bytesRead == blockSize when readErr == nil

// Store decrypted data and reset position

// Return decrypted data
