package option

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 方便創建一個帶有default值的物件, 並可以選擇性的修改一些物件的參數

func TestOption(t *testing.T) {
	s := NewServer("127.0.0.1", 8000, WithTimeout(30*time.Second), WithLogLevel("error"))
	assert.Equal(t, &Server{"127.0.0.1", 8000, 30 * time.Second, "error"}, s)
}
