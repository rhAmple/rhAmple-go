// Package rhAmple enables creation of rhAmple instances using different
// strategys for off-chain simulation.
package rhAmple

import (
	"github.com/rhAmple/rhAmple-go/strategies"
)

type RhAmple struct {
	strategy strategies.Strategy
}

// GetRhAmple returns an rhAmple instance simulating the current rhAmple
// version on-chain.
func GetRhAmple() RhAmple {
	strategy := strategies.GetMainnetStrategyV1()

	return RhAmple{
		strategy: strategy,
	}
}

// StrategySignal returns the rhAmple's current hedging signal deciding whether
// an upcoming rebase should be hedged or not.
func (rh *RhAmple) StrategySignal() (strategies.Signal, error) {
	return rh.strategy.Signal()
}
