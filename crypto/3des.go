package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// By3Des encrypts by triple des.
func (e Encrypter) By3Des(c *cipher.TripleDesCipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// By3Des decrypts by triple des.
func (d Decrypter) By3Des(c *cipher.TripleDesCipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
