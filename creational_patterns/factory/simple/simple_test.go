package simple

import (
	"testing"
)

// 避免直接使用p := &Person{}的方式創建object, 並確保創建的實例具有需要的參數

func TestSimple(t *testing.T) {
	p := NewPerson("abc", 18)
	p.Greet()
}
