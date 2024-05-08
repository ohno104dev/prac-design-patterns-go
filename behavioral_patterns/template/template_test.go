package template

import "testing"

// 在base類中定義一個模板框架, 允許sub類可以重寫特定的步驟

func TestTemplate(t *testing.T) {
	RunWorkflow(&A{})
	RunWorkflow(&B{})
}
