package adapter

import "fmt"

type UsbCAdapter struct {
	usbc *UsbC
}

func (u *UsbCAdapter) ConnVGA() {
	fmt.Println("adapter converts USB-C to VGA")
	u.usbc.ConnUsbc()
}

type HdmiAdapter struct {
	hdmi *Hdmi
}

func (h *HdmiAdapter) ConnVGA() {
	fmt.Println("adapter converts HDMI to VGA")
	h.hdmi.ConnHdmi()
}
