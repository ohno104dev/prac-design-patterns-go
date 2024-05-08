package observer

import "testing"

// 可在object(被觀察者/publisher)的事件發生時,通知多個observers(subscribers)

// 一種訂閱機制

func TestObserver(t *testing.T) {

	newItem := NewItem("Nike Shirt")

	user1 := &Customer{id: "abc@gmail.com"}
	newItem.register(user1)
	user2 := &Customer{id: "xyz@gmail.com"}
	newItem.register(user2)

	newItem.updateAvailability()
}
