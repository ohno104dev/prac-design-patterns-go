package lazy

import (
	"fmt"
	"testing"
)

// 使用時才初始化
// 會有concurrency問題, 需要加鎖

func TestLazy(t *testing.T) {
	ins := GetIns()
	fmt.Println(ins)
}
