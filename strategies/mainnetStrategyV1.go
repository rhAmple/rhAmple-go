package strategies

import "github.com/rhAmple/rhAmple-go/ampleforth"

type mainnetStrategyV1 struct {
}

func GetMainnetStrategyV1() *mainnetStrategyV1 {
	return &mainnetStrategyV1{}
}

func (s *mainnetStrategyV1) Signal() (Signal, error) {
	tokenInfo, err := ampleforth.FetchTokenInfo()
	if err != nil {
		return Invalid, err
	}

	targetRate := tokenInfo.LastRebaseInfo.TargetRateAtRebase

	price, err := ampleforth.FetchPrice()
	if err != nil {
		return Invalid, err
	}

	// Signal is hedge if price is lower than target rate.
	if price < targetRate {
		return Hedge, nil
	} else {
		return Dehedge, nil
	}
}
