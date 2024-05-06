package facade

// 外觀模式
// 提供一個簡單的接口, 把子系統的複雜性隱藏起來, 容易使用

import (
	"fmt"

	"github.com/ohno104dev/prac-design-pattern-go/structural_patterns/facade/account"
	"github.com/ohno104dev/prac-design-pattern-go/structural_patterns/facade/risk"
	"github.com/ohno104dev/prac-design-pattern-go/structural_patterns/facade/sales"
)

type Facade struct {
	sales   sales.SalesSystem
	risk    risk.RiskSystem
	account account.AccountSystem
}

func NewFacade() *Facade {
	return &Facade{
		*sales.NewSalesSystem(),
		*risk.NewRiskSystem(),
		*account.NewAccountingSystem(),
	}
}

func (f *Facade) Loan() {
	fmt.Println("loan process")
	f.sales.Apply()
	f.risk.Assess()
	f.account.Trans()
}
