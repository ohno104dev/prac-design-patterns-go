package funcs

import "testing"

// 解耦合, 變成由一群子類負責對具體類的實例化

func TestFuncs(t *testing.T) {
	babyFactory := NewPersonFactory(1)
	teenagerFactory := NewPersonFactory(18)

	baby := babyFactory("baby")
	teenager := teenagerFactory("teenager")

	baby.Greet()
	teenager.Greet()
}
