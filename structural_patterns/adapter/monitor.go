package adapter

import "fmt"

type Monitor struct{}

func (m *Monitor) InputVGA(conn Conn) {
	fmt.Println("waiting for VGA source")
	conn.ConnVGA()
}

type Conn interface {
	ConnVGA()
}
