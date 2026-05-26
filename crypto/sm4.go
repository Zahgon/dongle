package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// BySm4 encrypts by sm4.
func (e Encrypter) BySm4(c *cipher.Sm4Cipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// BySm4 decrypts by sm4.
func (d Decrypter) BySm4(c *cipher.Sm4Cipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
