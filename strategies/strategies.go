package strategies

type Signal int

const (
	Hedge Signal = iota
	Dehedge
	Invalid
)

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
