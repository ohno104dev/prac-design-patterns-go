package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

type House struct {
	windowType string
	doorType   string
	floor      int
}
