// Package strategies provides simulations of strategy implementations used
// on-chain.
package strategies

// Signal represents the possible signals a rebase strategy can return.
type Signal int

const (
	Hedge Signal = iota
	Dehedge
	Invalid
)

// A strategy returns a signal whether an upcoming rebase should be hedged or
// not.
type Strategy interface {
	// Signal returns the hedging signal the strategy implementation returns
	// on-chain.
	Signal() (Signal, error)
}

func (s Signal) String() string {
	switch s {
	case Hedge:
		return "hedge"
	case Dehedge:
		return "dehedge"
	case Invalid:
		return "invalid"
	default:
		return "unknown signal type"
	}
}
