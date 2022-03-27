package strategies

import "github.com/rhAmple/rhAmple-go/ampleforth"

type mainnetStrategyV1 struct {
}

func GetMainnetStrategyV1() *mainnetStrategyV1 {
	return &mainnetStrategyV1{}
}

func (s *mainnetStrategyV1) Signal() (Signal, error) {
	ampleInfo, err := ampleforth.Info()
	if err != nil {
		return Invalid, err
	}

	// Signal is hedge if price is lower than target rate.
	if ampleInfo.Price < ampleInfo.TargetRate {
		return Hedge, nil
	} else {
		return Dehedge, nil
	}
}
