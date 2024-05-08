package template

import "fmt"

type BaseWorkflow struct{}

func (b *BaseWorkflow) start() {
	fmt.Println("工作流程開始")
}

func (b *BaseWorkflow) end() {
	fmt.Println("工作流程結束")
}

func (b *BaseWorkflow) execute() {
	// 可以自訂的部分
}

type Workflow interface {
	start()
	execute()
	end()
}

func RunWorkflow(w Workflow) {
	w.start()
	w.execute()
	w.end()
}
