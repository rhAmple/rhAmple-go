package rhAmple

import "github.com/rhAmple/rhAmple-go/strategies"

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
