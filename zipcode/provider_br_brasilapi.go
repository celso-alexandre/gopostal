package zipcode

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Provider_BR_BrasilApi_ZipCodeInfo struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func (v *Provider_BR_BrasilApi_ZipCodeInfo) ToZipCodeInfo() *ZipCodeInfo {
	return &ZipCodeInfo{
		ProviderName: "br_brasilapi",
		ZipCode:      NormalizeBrazilZipCode(v.CEP),
		State:        v.State,
		City:         v.City,
		Neighborhood: v.Neighborhood,
		Street:       v.Street,
	}
}

type Provider_BR_BrasilApi struct{}

func (v Provider_BR_BrasilApi) GetZipCodeDetails(zipCode string) *ZipCodeProviderResponse {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", zipCode)
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

	var zipInfo Provider_BR_BrasilApi_ZipCodeInfo
	err = json.Unmarshal(body, &zipInfo)
	if err != nil {
		return &ZipCodeProviderResponse{
			Err:         ErrZipCodeInternal,
			ZipCodeInfo: nil,
		}
	}

	return &ZipCodeProviderResponse{
		Err:         ErrZipCodeNoError,
		ZipCodeInfo: zipInfo.ToZipCodeInfo(),
	}
}
