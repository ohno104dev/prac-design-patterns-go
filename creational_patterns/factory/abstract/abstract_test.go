package abstract

import "testing"

// 返回interface, 在不公開內部實現的情況下讓調用者使用提供的func

func TestAbstract(t *testing.T) {
	p := NewPerson("jack", 18)
	p.Greet()
}
