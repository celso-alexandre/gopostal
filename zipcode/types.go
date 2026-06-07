package zipcode

type ZipCodeInfo struct {
	ProviderName string `json:"provider_name"`
	ZipCode      string `json:"zip_code"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	// IBGE is the 7-digit IBGE municipality code. Populated by providers that expose
	// it (ViaCEP); empty for providers that don't (BrasilAPI v1).
	IBGE string `json:"ibge"`
}
