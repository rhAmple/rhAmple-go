package ampleforth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Coingecko API constants.
// For more info see https://www.coingecko.com/en/api/documentation.
const (
	api               = "https://api.coingecko.com/api/v3"
	tokenId           = "ampleforth"
	priceDenomination = "usd"
)

type priceInfo struct {
	Ampleforth struct {
		Usd float64 `json:"usd"`
	} `json:"ampleforth"`
}

func price() (float64, error) {
	endpoint := api + "/simple/price"
	v := url.Values{}
	v.Set("ids", tokenId)
	v.Set("vs_currencies", priceDenomination)

	c := http.Client{}

	req, err := http.NewRequest(http.MethodGet, endpoint+"?"+v.Encode(), nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", userAgent)

	res, err := c.Do(req)
	if err != nil {
		return 0, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	priceInfo := priceInfo{}
	err = json.Unmarshal(body, &priceInfo)
	if err != nil {
		return 0, err
	}

	return priceInfo.Ampleforth.Usd, nil
}
