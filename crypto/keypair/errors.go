package keypair

type EmptyPublicKeyError struct {
}

func (e EmptyPublicKeyError) Error() string { _ = "STUB: not implemented"; return "" }

type InvalidPublicKeyError struct {
	Err error
}

func (e InvalidPublicKeyError) Error() string { _ = "STUB: not implemented"; return "" }

type EmptyPrivateKeyError struct {
}

func (e EmptyPrivateKeyError) Error() string { _ = "STUB: not implemented"; return "" }

type InvalidPrivateKeyError struct {
	Err error
}

func (e InvalidPrivateKeyError) Error() string { _ = "STUB: not implemented"; return "" }

type EmptyFormatError struct {
}

func (e EmptyFormatError) Error() string { _ = "STUB: not implemented"; return "" }

type UnsupportedKeyFormatError struct {
}

func (e UnsupportedKeyFormatError) Error() string { _ = "STUB: not implemented"; return "" }

type EmptyPaddingError struct {
}

func (e EmptyPaddingError) Error() string { _ = "STUB: not implemented"; return "" }

type UnsupportedPaddingSchemeError struct {
	Padding string
}

func (e UnsupportedPaddingSchemeError) Error() string { _ = "STUB: not implemented"; return "" }

type EmptySignatureError struct {
}

func (e EmptySignatureError) Error() string { _ = "STUB: not implemented"; return "" }
