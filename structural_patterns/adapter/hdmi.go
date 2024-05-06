package adapter

import "fmt"

type Hdmi struct{}

func (h *Hdmi) ConnHdmi() {
	fmt.Println("HDMI port is connected")
}
