package zipcode

type ProviderGroupName = string

const (
	ProviderGroupBrazil ProviderGroupName = "BR"
)

var ProviderGroups = map[ProviderGroupName][]ZipCodeProvider{
	ProviderGroupBrazil: {ViaCEPProvider{}, Provider_BR_BrasilApi{}},
}
