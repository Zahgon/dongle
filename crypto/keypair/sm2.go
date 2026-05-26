package keypair

import (
	"crypto/ecdsa"

	"github.com/dromara/dongle/crypto/internal/sm2"
)

// Sm2CipherMode specifies the concatenation mode of SM2 ciphertext
// components. It controls how the library assembles (encrypt) and
// interprets (decrypt) the C1, C2, C3 parts.
//
// C1: EC point (x1||y1) in uncompressed form; C2: XORed plaintext;
// C3: SM3 digest over x2 || M || y2.
//
// NOTE: For performance and boundary checks, would it be better to set the type to uint8?
type Sm2CipherMode string

// Supported SM2 ciphertext orders.
const (
	// C1C2C3 means ciphertext bytes are C1 || C2 || C3 in bytes.
	C1C2C3 Sm2CipherMode = "c1c2c3"
	// C1C3C2 means ciphertext bytes are C1 || C3 || C2 in bytes.
	C1C3C2 Sm2CipherMode = "c1c3c2"
	// ASN1C1C2C3 means ciphertext bytes are C1 || C2 || C3 in ASN1.
	ASN1C1C2C3 Sm2CipherMode = "asn1_c1c2c3"
	// ASN1C1C3C2 means ciphertext bytes are C1 || C3 || C2 in ASN1.
	ASN1C1C3C2 Sm2CipherMode = "asn1_c1c3c2"
)

type Sm2SingMode uint8

const (
	// Digital signature in ASN1 format
	ASN1 Sm2SingMode = iota
	// Digital signature in bytes format
	Bytes
)

var (
	bitStringPublicKeyParser  = sm2.ParseBitStringPublicKey
	bitStringPrivateKeyParser = sm2.ParseBitStringPrivateKey
)

// Sm2KeyPair represents an SM2 key pair with public and private keys.
// Keys are handled in PKCS8 (for private) and PKIX (for public) PEM formats.
type Sm2KeyPair struct {
	// PublicKey contains the PEM-encoded public key
	PublicKey []byte

	// PrivateKey contains the PEM-encoded private key
	PrivateKey []byte

	// Order specifies the mode of SM2 ciphertext components.
	// It controls how Encrypt assembles and Decrypt interprets ciphertext.
	// NOTE: Perhaps renaming this to CipherMode would be more appropriate?
	Mode Sm2CipherMode

	// SingMode controls the logic of signing and verification.
	// There are two common ways to handle SM2 signature data:
	// one is to encode R and S in ASN1 format, and the other is to concatenate R and S.
	//
	// Default is ASN1 format.
	SingMode Sm2SingMode

	// Window controls internal SM2 fixed-base/wNAF window size (2..6).
	// 4 means use library default.
	Window int

	// UID is the user identifier for SM2 signature operations.
	// If empty, the default UID "1234567812345678" will be used (per GM/T 0009-2012).
	UID []byte
}

// NewSm2KeyPair returns a new Sm2KeyPair with defaults
// (Order=C1C3C2, Window=4).
func NewSm2KeyPair() *Sm2KeyPair { _ = "STUB: not implemented"; return nil }

// GenKeyPair generates a new SM2 key pair and fills PublicKey/PrivateKey.
// Private key is PKCS#8 (PEM "PRIVATE KEY"), public key is SPKI/PKIX (PEM "PUBLIC KEY").
func (k *Sm2KeyPair) GenKeyPair() error {
	_ = "STUB: not implemented"

	// Generate unbiased scalar d in range [1, n-1]
	return nil
}

// Marshal PKCS8 private key

// Marshal SPKI public key

// SetOrder sets ciphertext order to C1C3C2 or C1C2C3.
// Deprecated: `SetOrder` will be removed in the future, use `SetMode` instead.
func (k *Sm2KeyPair) SetOrder(order Sm2CipherMode) {
	_ = "STUB: not implemented"

	// SetMode sets ciphertext mode to C1C3C2 or C1C2C3.
	// It affects how Encrypt assembles and Decrypt interprets ciphertext.
	return
}

func (k *Sm2KeyPair) SetMode(mode Sm2CipherMode) {
	_ = "STUB: not implemented"

	// SetSingMode sets the mode for SM2 Sign and Verify
	return
}

func (k *Sm2KeyPair) SetSingMode(mode Sm2SingMode) {
	_ = "STUB: not implemented"

	// SetWindow sets scalar-multiplication window (2..6).
	// Values outside the range are clamped.
	return
}

func (k *Sm2KeyPair) SetWindow(window int) { _ = "STUB: not implemented"; return }

// SetUID sets the user identifier for SM2 signature operations.
// If uid is nil or empty, the default UID "1234567812345678" will be used.
func (k *Sm2KeyPair) SetUID(uid []byte) {
	_ = "STUB: not implemented"

	// SetPublicKey sets the public key after formatting to PEM.
	// Accepts base64-encoded DER of SubjectPublicKeyInfo.
	return
}

func (k *Sm2KeyPair) SetPublicKey(publicKey []byte) error { _ = "STUB: not implemented"; return nil }

// SetPrivateKey sets the private key after formatting to PEM.
// Accepts base64-encoded DER of PKCS#8 PrivateKeyInfo.
func (k *Sm2KeyPair) SetPrivateKey(privateKey []byte) error { _ = "STUB: not implemented"; return nil }

// ParsePublicKey parses the PEM-encoded public key and returns *sm2.PublicKey.
func (k *Sm2KeyPair) ParsePublicKey() (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParsePrivateKey parses the PEM-encoded private key and returns *sm2.PrivateKey.
func (k *Sm2KeyPair) ParsePrivateKey() (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FormatPublicKey formats base64-encoded der public key into the specified PEM format.
func (k *Sm2KeyPair) FormatPublicKey(publicKey []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FormatPrivateKey formats base64-encoded der private key into the specified PEM format.
func (k *Sm2KeyPair) FormatPrivateKey(privateKey []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CompressPublicKey strips headers/footers and whitespace from the PEM public key.
func (k *Sm2KeyPair) CompressPublicKey(publicKey []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// CompressPrivateKey strips headers/footers and whitespace from the PEM private key.
func (k *Sm2KeyPair) CompressPrivateKey(privateKey []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}
