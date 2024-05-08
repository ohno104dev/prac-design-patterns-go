package strategy

type sub struct{}

func (*sub) do(a, b int) int {
	return a - b
}
