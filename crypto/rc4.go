package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// ByRc4 encrypts by rc4.
func (e Encrypter) ByRc4(c *cipher.Rc4Cipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByRc4 decrypts by rc4.
func (d Decrypter) ByRc4(c *cipher.Rc4Cipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
