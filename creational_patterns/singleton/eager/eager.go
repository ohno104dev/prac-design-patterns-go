package eager

type singleton struct{}

var ins *singleton = &singleton{}

func GetIns() *singleton {
	return ins
}
