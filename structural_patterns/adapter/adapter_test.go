package adapter

import "testing"

// 使接口不兼容的object能夠相互合作

func TestAdapter(t *testing.T) {
	monitor := &Monitor{}

	usbcAdapter := &UsbCAdapter{
		&UsbC{},
	}

	monitor.InputVGA(usbcAdapter)

	hdmiAdapter := &HdmiAdapter{
		&Hdmi{},
	}

	monitor.InputVGA(hdmiAdapter)
}
