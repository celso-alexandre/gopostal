package test

import (
	"strings"
	"testing"
	"time"

	"github.com/celso-alexandre/gopostal"
	"github.com/celso-alexandre/gopostal/zipcode"
	"github.com/stretchr/testify/require"
)

func TestFetchZipCode_Integration(t *testing.T) {
	info, err := gopostal.FetchZipCode("01001-000", zipcode.ProviderGroupBrazil, 3*time.Second)

	require.NoError(t, err)
	require.NotNil(t, info)
	require.Equal(t, "01001000", strings.ReplaceAll(info.ZipCode, "-", ""))
	t.Logf("ZipCode Info: %+v\n", info)
}
