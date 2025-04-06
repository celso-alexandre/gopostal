package zipcode

type ProviderGroupName = string

const (
	ProviderGroupBrazil ProviderGroupName = "BR"
)

var ProviderGroups = map[ProviderGroupName][]ZipCodeProvider{
	ProviderGroupBrazil: {Provider_BR_ViaCEP{}, Provider_BR_BrasilApi{}},
}
