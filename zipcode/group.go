package zipcode

type ProviderGroupName = string

const (
	ProviderGroupBrazil ProviderGroupName = "BR"

	// ProviderGroupBrazilIBGE only contains providers that populate
	// ZipCodeInfo.Brazil.IBGECode (the IBGE municipality code). Use it when the
	// IBGE code is required: racing the full BR group returns the first answer,
	// which may come from a provider that does not carry it.
	ProviderGroupBrazilIBGE ProviderGroupName = "BR_IBGE"
)

var ProviderGroups = map[ProviderGroupName][]ZipCodeProvider{
	ProviderGroupBrazil:     {Provider_BR_ViaCEP{}, Provider_BR_BrasilApi{}},
	ProviderGroupBrazilIBGE: {Provider_BR_ViaCEP{}},
}
