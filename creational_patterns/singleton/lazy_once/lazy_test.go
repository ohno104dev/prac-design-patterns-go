package lazyonce

import (
	"fmt"
	"testing"
)

// 使用時才初始化
// 使用sync.Once, 簡潔無需加鎖

func TestLazy(t *testing.T) {
	ins := GetIns()
	fmt.Println(ins)
}
