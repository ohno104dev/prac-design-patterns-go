package lazy

import "sync"

type singleton struct{}

var ins *singleton
var mu sync.Mutex

func GetIns() *singleton {
	// 第一層非nil直接回傳,不用每次進來都lock
	// 確定nil後第二層lock避免concurrency問題
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			ins = &singleton{}
		}
	}

	return ins
}
