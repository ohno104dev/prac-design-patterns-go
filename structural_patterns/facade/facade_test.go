package facade

import (
	"testing"
)

// 提供一個簡單的接口, 把子系統的複雜性隱藏起來

func TestFacade(t *testing.T) {
	c := NewFacade()
	c.Loan()
}
