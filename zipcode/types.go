package zipcode

type ZipCodeInfo struct {
	ProviderName string `json:"provider_name"`
	ZipCode      string `json:"zip_code"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	// Brazil holds Brazil-specific attributes, populated for BR lookups by providers
	// that expose them (e.g. ViaCEP). Nil when unavailable.
	Brazil *BrazilZipCodeInfo `json:"brazil,omitempty"`
}

// BrazilZipCodeInfo carries Brazil-only postal attributes.
type BrazilZipCodeInfo struct {
	// IBGECode is the 7-digit IBGE municipality code.
	IBGECode string `json:"ibge_code"`
}
