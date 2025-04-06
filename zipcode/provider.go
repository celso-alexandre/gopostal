package zipcode

type ZipCodeProvider interface {
	GetZipCodeDetails(zipCode string) *ZipCodeProviderResponse
}

type ZipCodeProviderResponse struct {
	Err         ZipCodeError
	ZipCodeInfo *ZipCodeInfo
}
