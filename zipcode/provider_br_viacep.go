package zipcode

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Provider_BR_ViaCep_ZipCodeInfo struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v *Provider_BR_ViaCep_ZipCodeInfo) ToZipCodeInfo() *ZipCodeInfo {
	return &ZipCodeInfo{
		ProviderName: "Provider_BR_ViaCep",
		ZipCode:      v.CEP,
		Street:       v.Logradouro,
		Neighborhood: v.Bairro,
		City:         v.Localidade,
		State:        v.UF,
	}
}

type ViaCEPProvider struct{}

func (v ViaCEPProvider) GetZipCodeDetails(zipCode string) *ZipCodeProviderResponse {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)
	resp, err := http.Get(url)
	if err != nil {
		return &ZipCodeProviderResponse{
			Err:         ErrZipCodeTimeout,
			ZipCodeInfo: nil,
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return &ZipCodeProviderResponse{
				Err:         ErrZipCodeNotFound,
				ZipCodeInfo: nil,
			}
		}
		if resp.StatusCode == http.StatusBadRequest {
			return &ZipCodeProviderResponse{
				Err:         ErrZipCodeInvalid,
				ZipCodeInfo: nil,
			}
		}
		return &ZipCodeProviderResponse{
			Err:         ErrZipCodeInternal,
			ZipCodeInfo: nil,
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ZipCodeProviderResponse{
			Err:         ErrZipCodeInternal,
			ZipCodeInfo: nil,
		}
	}

	var zipInfo Provider_BR_ViaCep_ZipCodeInfo
	err = json.Unmarshal(body, &zipInfo)
	if err != nil {
		return &ZipCodeProviderResponse{
			Err:         ErrZipCodeInternal,
			ZipCodeInfo: nil,
		}
	}

	if zipInfo.CEP == "" {
		return &ZipCodeProviderResponse{
			Err:         ErrZipCodeNotFound,
			ZipCodeInfo: nil,
		}
	}

	return &ZipCodeProviderResponse{
		Err:         ErrZipCodeNoError,
		ZipCodeInfo: zipInfo.ToZipCodeInfo(),
	}
}
