package test

import (
	"testing"
	"time"

	"github.com/celso-alexandre/gopostal"
	"github.com/celso-alexandre/gopostal/zipcode"
	"github.com/stretchr/testify/require"
)

func TestFetchZipCodeBRProviderGroup_Integration(t *testing.T) {
	zip := "01001-000"
	info, err := gopostal.FetchZipCode(zip, zipcode.ProviderGroupBrazil, 3*time.Second)

	require.NoError(t, err)
	require.NotNil(t, info)
	t.Logf("ZipCode Info: %+v\n", info)
	require.Equal(t, zip, info.ZipCode)
}
