package strategy

type add struct{}

func (*add) do(a, b int) int {
	return a + b
}
