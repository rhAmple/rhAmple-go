package rhAmple

import "github.com/rhAmple/rhAmple-go/strategies"

type RhAmple struct {
	strategy strategies.Strategy
}

func GetRhAmple() *RhAmple {
	strategy := strategies.GetMainnetStrategyV1()

	return &RhAmple{
		strategy: strategy,
	}
}
