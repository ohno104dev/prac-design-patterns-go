package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 將複雜類別的struct與實作細節分離, 透過相同的struct呈現不同的實作細節

// 適合創建製造過程相似且僅有細節上的差異的object
// 避免過多的構造函數的出現

func TestBuilder(t *testing.T) {
	normalBuilder := newNormalBuilder()
	iglooBuilder := newIglooBuilder()

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	assert.Equal(t, normalHouse.windowType, "Wooden Window")
	assert.Equal(t, normalHouse.doorType, "Wooden Door")
	assert.Equal(t, normalHouse.floor, 2)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	assert.Equal(t, iglooHouse.windowType, "Snow Window")
	assert.Equal(t, iglooHouse.doorType, "Snow Door")
	assert.Equal(t, iglooHouse.floor, 1)
}
