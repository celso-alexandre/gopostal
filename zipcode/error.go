package zipcode

type ZipCodeError = string

const (
	ErrZipCodeNoError  ZipCodeError = ""
	ErrZipCodeNotFound ZipCodeError = "not_found"
	ErrZipCodeInvalid  ZipCodeError = "invalid"
	ErrZipCodeTimeout  ZipCodeError = "timeout"
	ErrZipCodeInternal ZipCodeError = "internal"
)
