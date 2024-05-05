package eager

import (
	"fmt"
	"testing"
)

// package導入時初始化
// 導致加載時間比較長

func TestEager(t *testing.T) {
	ins := GetIns()
	fmt.Println(ins)
}
