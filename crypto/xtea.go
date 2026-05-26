package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// ByXtea encrypts by xtea.
func (e Encrypter) ByXtea(c *cipher.XteaCipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByXtea decrypts by xtea.
func (d Decrypter) ByXtea(c *cipher.XteaCipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
