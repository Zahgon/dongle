package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// BySalsa20 encrypts by salsa20.
func (e Encrypter) BySalsa20(c *cipher.Salsa20Cipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// BySalsa20 decrypts by salsa20.
func (d Decrypter) BySalsa20(c *cipher.Salsa20Cipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
