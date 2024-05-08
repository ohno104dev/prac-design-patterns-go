package template

import "fmt"

type A struct {
	BaseWorkflow
}

func (*A) execute() {
	fmt.Println("工作流程A執行中")
}
