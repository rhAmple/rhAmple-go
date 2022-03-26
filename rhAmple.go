package rhAmple

import (
	"github.com/rhAmple/rhAmple-go/ampleforth"
	"github.com/rhAmple/rhAmple-go/strategies"
)

type rhAmple struct {
	strategy strategies.Strategy
}

func GetRhAmple() *rhAmple {
	strategy := strategies.GetMainnetStrategyV1()

	return &rhAmple{
		strategy: strategy,
	}
}

func (rh *rhAmple) StrategySignal() (strategies.Signal, error) {
	return rh.strategy.Signal()
}

func (rh *rhAmple) NextRebase() (int64, error) {
	tokenInfo, err := ampleforth.FetchTokenInfo()
	if err != nil {
		return 0, err
	}

	return tokenInfo.NextRebaseTimeSec, nil
}
