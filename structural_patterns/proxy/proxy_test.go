package proxy

import (
	"testing"
)

// 提供一種代理在不影響原本object的功能的情況下, 額外增加其他功能

func TestProxy(t *testing.T) {
	teenage := NewCarProxy(&Driver{Age: 12})
	teenage.DriveIfAudit()

	audit := NewCarProxy(&Driver{Age: 20})
	audit.DriveIfAudit()
}
