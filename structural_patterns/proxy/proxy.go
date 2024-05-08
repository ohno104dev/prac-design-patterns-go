package proxy

import "fmt"

type CarProxy struct {
	vehicle Vehicle
	driver  *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{&Car{}, driver}
}

// 新功能
func (c *CarProxy) DriveIfAudit() {
	if c.driver.Age >= 18 {
		c.vehicle.Drive()
	} else {
		fmt.Println("driver too young")
	}
}
