# gopostal

**`gopostal`** is a Go library for fast and fault-tolerant ZIP/postal code lookups using multiple providers concurrently. It returns as soon as one provider succeeds.

## 🚀 Features

- Supports concurrent ZIP code lookups
- Built-in support for Brazilian postal codes (CEP)
- Extensible with custom providers
- Optional timeout for responsiveness
- Graceful handling of provider failures

## 🧩 Extending Providers
- Currently missing support for other countries than Brazil

To create a new ZIP code provider, implement the ZipCodeProvider interface:
```go
// copy this file: zipcode/provider_br_viacep.go -> zipcode/provider_us_providername.go
type ZipCodeProvider interface {
	GetZipCodeDetails(zip string) ProviderResult
}

type ProviderResult struct {
	ZipCodeInfo *ZipCodeInfo
	Err         string
}
```

Adding a new provider group (country, probably):
```go
// modify file: zipcode/group.go
ProviderGroupUnitedStates ProviderGroupName = "US"

var ProviderGroups = map[ProviderGroupName][]ZipCodeProvider{
	...
	ProviderGroupUnitedStates: {Provider_US_ProviderName{}},
}
```

## 🤝 Contributing
PRs welcome! Add support for other countries or improve the provider fallback logic.

## 🔍 TODO
- Better organize the project (split in more packages)
- tests
- Add more providers?

## 📦 Installation

```bash
go get github.com/celso-alexandre/gopostal
```

## Usage
```go
package main

import (
	"fmt"
	"time"

	"github.com/celso-alexandre/gopostal"
	"github.com/celso-alexandre/gopostal/zipcode"
)

func main() {
	zip := "01001-000" // Example CEP (Brazil)
	info, err := gopostal.FetchZipCode(zip, zipcode.ProviderGroupBrazil, 3*time.Second)
	if err != nil {
		fmt.Println("Failed:", err)
		return
	}
	fmt.Printf("Found: %+v\n", info)
}
```
