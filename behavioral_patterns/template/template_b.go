package template

import "fmt"

type B struct {
	BaseWorkflow
}

func (*B) execute() {
	fmt.Println("工作流程B執行中")
}
