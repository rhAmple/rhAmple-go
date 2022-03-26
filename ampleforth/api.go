package ampleforth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// Ampleforth API constants.
const (
	host  = "https://web-api.ampleforth.org/"
	chain = "eth"
)

// TokenInfo encapsulates Ampleforth's WebApi response for /token-info.
type TokenInfo struct {
	Addresses struct {
		NonCirculatingWallets []string `json:"nonCirculatingWallets"`
		Orchestrator          string   `json:"orchestrator"`
		Policy                string   `json:"policy"`
		TokenContract         string   `json:"tokenContract"`
	} `json:"addresses"`
	Constants struct {
		BaseCpi  float64 `json:"baseCpi"`
		Decimals int64   `json:"decimals"`
		Name     string  `json:"name"`
		Symbol   string  `json:"symbol"`
	} `json:"constants"`
	Epoch          int64 `json:"epoch"`
	LastRebaseInfo struct {
		AppliedSupplyAdjustment int64   `json:"appliedSupplyAdjustment"`
		Epoch                   int64   `json:"epoch"`
		MarketRateAtRebase      float64 `json:"marketRateAtRebase"`
		PrecentageChange        int64   `json:"precentageChange"`
		PreviousSupply          float64 `json:"previousSupply"`
		TargetRateAtRebase      float64 `json:"targetRateAtRebase"`
		TimestampSec            int64   `json:"timestampSec"`
		TotalSupply             float64 `json:"totalSupply"`
	} `json:"lastRebaseInfo"`
	NextRebaseTimeSec int64 `json:"nextRebaseTimeSec"`
	PolicyParams      struct {
		DeviationThresholdPerc float64 `json:"deviationThresholdPerc"`
		MinRebaseTimeSec       int64   `json:"minRebaseTimeSec"`
		RebaseWindowLengthSec  int64   `json:"rebaseWindowLengthSec"`
		RebaseWindowOffsetSec  int64   `json:"rebaseWindowOffsetSec"`
	} `json:"policyParams"`
	SupplyInfo struct {
		CirculatingSupply float64 `json:"circulatingSupply"`
		LockedBalance     float64 `json:"lockedBalance"`
		TotalSupply       float64 `json:"totalSupply"`
	} `json:"supplyInfo"`
}

// FetchTokenInfo returns a TokenInfo struct fetched from the Ampleforth
// WebApi.
func FetchTokenInfo() (*TokenInfo, error) {
	url := filepath.Join(host, chain, "token-info")

	c := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	tokenInfo := TokenInfo{}
	err = json.Unmarshal(body, &tokenInfo)
	if err != nil {
		return nil, err
	}

	return &tokenInfo, nil
}
