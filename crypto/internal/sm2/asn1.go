package sm2

import (
	"crypto/ecdsa"
)

// MarshalSPKIPublicKey encodes a SubjectPublicKeyInfo (SPKI) for the given SM2 public key.
func MarshalSPKIPublicKey(pub *ecdsa.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AlgorithmIdentifier

// subjectPublicKey BIT STRING

// MarshalPKCS8PrivateKey encodes a PKCS#8 PrivateKeyInfo for the given SM2 private key.
func MarshalPKCS8PrivateKey(pri *ecdsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// version

// privateKey OCTET STRING wrapping ECPrivateKey

// ec version

// [0] parameters namedCurve OID (explicit)

// [1] publicKey BIT STRING (explicit)

// ParseSPKIPublicKey parses a SubjectPublicKeyInfo (SPKI) and returns an SM2 public key.
func ParseSPKIPublicKey(der []byte) (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParsePKCS8PrivateKey parses a PKCS#8 PrivateKeyInfo and returns an SM2 private key.
// Simplified: ignores optional parameters/publicKey fields inside ECPrivateKey.
func ParsePKCS8PrivateKey(der []byte) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ECPrivateKey (version, d)

// ParseBitStringPublicKey parses a BIT_STRING PublicKeyInfo and returns an SM2 public key.
//
//go:inline
func ParseBitStringPublicKey(key []byte) (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseBitStringPrivateKey parses a BIT_STRING PrivateKeyInfo and returns an SM2 private key.
//
//go:inline
func ParseBitStringPrivateKey(key []byte) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
