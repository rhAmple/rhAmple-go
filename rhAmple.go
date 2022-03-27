// Package rhAmple enables creation of rhAmple instances using different
// strategys for off-chain simulation.
package rhAmple

import (
	"github.com/rhAmple/rhAmple-go/strategies"
)

type rhAmple struct {
	strategy strategies.Strategy
}

// GetRhAmple returns an rhAmple instance simulating the current rhAmple
// version on-chain.
func GetRhAmple() rhAmple {
	strategy := strategies.GetMainnetStrategyV1()

	return rhAmple{
		strategy: strategy,
	}
}

// StrategySignal returns the rhAmple's current hedging signal deciding whether
// an upcoming rebase should be hedged or not.
func (rh *rhAmple) StrategySignal() (strategies.Signal, error) {
	return rh.strategy.Signal()
}
