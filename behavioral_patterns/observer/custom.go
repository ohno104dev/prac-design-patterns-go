package observer

import "fmt"

type Customer struct {
	id string
}

func (c *Customer) getId() string {
	return c.id
}

func (c *Customer) update(itemName string) {
	fmt.Printf("sending email to customer [%s] for item [%s] \n", c.id, itemName)
}
