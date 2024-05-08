package strategy

type IStrategy interface {
	do(int, int) int
}

type Operator struct {
	strategy IStrategy
}

func (o *Operator) setStrategy(strategy IStrategy) {
	o.strategy = strategy
}

func (o *Operator) Calculate(a, b int) int {
	return o.strategy.do(a, b)
}
