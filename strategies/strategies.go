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
	if s == Hedge {
		return "hedge"
	}
	if s == Dehedge {
		return "dehedge"
	}
	return "invalid"
}
