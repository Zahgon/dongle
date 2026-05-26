package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// ByChaCha20 encrypts by chacha20.
func (e Encrypter) ByChaCha20(c *cipher.ChaCha20Cipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByChaCha20 decrypts by chacha20.
func (d Decrypter) ByChaCha20(c *cipher.ChaCha20Cipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
