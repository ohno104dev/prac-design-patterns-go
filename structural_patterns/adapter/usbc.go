package adapter

import "fmt"

type UsbC struct{}

func (u *UsbC) ConnUsbc() {
	fmt.Println("USB-C port is connected")
}
