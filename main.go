package gopostal

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/celso-alexandre/gopostal/zipcode"
)

// FetchZipCode fetches zip code details using a specific provider group.
func FetchZipCode(
	zipCode string,
	providerGroup zipcode.ProviderGroupName,
	timeout time.Duration,
) (*zipcode.ZipCodeInfo, error) {
	return FetchZipCodeFromProviders(zipCode, zipcode.ProviderGroups[providerGroup], timeout)
}

// FetchZipCodeFromProviders tries multiple providers concurrently and returns the first successful result.
func FetchZipCodeFromProviders(zipCode string, providers []zipcode.ZipCodeProvider, timeout time.Duration) (*zipcode.ZipCodeInfo, error) {
	var wg sync.WaitGroup
	providersCount := len(providers)

	resultChan := make(chan *zipcode.ZipCodeInfo, 1)
	errChan := make(chan error, providersCount)

	var isDoneMu sync.Mutex
	isDone := false

	errorsCount := 0
	var errorCountMu sync.Mutex

	for _, provider := range providers {
		wg.Add(1)
		go func(p zipcode.ZipCodeProvider) {
			defer wg.Done()
			info := p.GetZipCodeDetails(zipCode)
			isDoneMu.Lock()
			if isDone {
				isDoneMu.Unlock()
				return
			}
			isDoneMu.Unlock()

			if info.Err != "" {
				errChan <- errors.New(info.Err)
				return
			}

			select {
			case resultChan <- info.ZipCodeInfo:
			default:
			}
		}(provider)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	timeoutChan := time.After(timeout)

	for {
		select {
		case result := <-resultChan:
			isDoneMu.Lock()
			isDone = true
			isDoneMu.Unlock()
			return result, nil

		case <-timeoutChan:
			isDoneMu.Lock()
			isDone = true
			isDoneMu.Unlock()
			return nil, errors.New(zipcode.ErrZipCodeTimeout)

		case err := <-errChan:
			errorCountMu.Lock()
			errorsCount++
			allFailed := errorsCount == providersCount
			errorCountMu.Unlock()

			if allFailed {
				isDoneMu.Lock()
				isDone = true
				isDoneMu.Unlock()
				return nil, err
			} else {
				fmt.Println("FetchZipCodeDetails (not yet given up) err:", err)
			}
		}
	}
}
