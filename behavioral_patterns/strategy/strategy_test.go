package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 把一系列算法封裝起來,使他們能夠相互替換

// 添加新策略不會對上下文造成影響

func TestStrategy(t *testing.T) {
	op := &Operator{}
	op.setStrategy(&add{})
	res := op.Calculate(1, 2)
	assert.Equal(t, res, 3)

	op.setStrategy(&sub{})
	res = op.Calculate(5, 2)
	assert.Equal(t, res, 3)
}
