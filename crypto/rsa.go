package crypto

import (
	"github.com/dromara/dongle/crypto/keypair"
)

// ByRsa encrypts by rsa.
func (e Encrypter) ByRsa(kp *keypair.RsaKeyPair) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByRsa decrypts by rsa.
func (d Decrypter) ByRsa(kp *keypair.RsaKeyPair) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode

// ByRsa signs by rsa.
func (s Signer) ByRsa(kp *keypair.RsaKeyPair) Signer {
	_ = "STUB: not implemented"
	return *new(Signer)
}

// Streaming signing mode

// Standard signing mode

// ByRsa verifies by rsa.
func (v Verifier) ByRsa(kp *keypair.RsaKeyPair) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}

// Streaming verification mode

// Write the data to be verified

// Close the verifier to perform verification

// Standard verification mode
