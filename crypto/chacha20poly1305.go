package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// ByChaCha20Poly1305 encrypts by chacha20-poly1305.
func (e Encrypter) ByChaCha20Poly1305(c *cipher.ChaCha20Poly1305Cipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByChaCha20Poly1305 decrypts by chacha20-poly1305.
func (d Decrypter) ByChaCha20Poly1305(c *cipher.ChaCha20Poly1305Cipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
