// Package ampleforth is a stateless library for fetching data about
// Ampleforth.
package ampleforth

// Info fetches data about Ampleforth and returns it encapsulated
// in an AmpleInfo struct.
func Info() (AmpleInfo, error) {
	price, err := price()
	if err != nil {
		return AmpleInfo{}, err
	}

	tokenInfo, err := tokenInfo()
	if err != nil {
		return AmpleInfo{}, err
	}

	return AmpleInfo{
		TargetRate:    tokenInfo.LastRebaseInfo.TargetRateAtRebase,
		Price:         price,
		NextRebaseSec: tokenInfo.NextRebaseTimeSec,
	}, nil
}

// AmpleInfo struct encapsulates Ampleforth data.
type AmpleInfo struct {
	TargetRate    float64
	Price         float64
	NextRebaseSec int64
}
