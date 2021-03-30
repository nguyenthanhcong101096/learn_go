package builder

type iglooBuilder struct {
	windownType string
	floor int
}

func NewIglooBuidler() *iglooBuilder {
	return &iglooBuilder{}
}

func (n *iglooBuilder) setWindownType() {
	n.windownType = "igloo windown type"
}

func (n *iglooBuilder) setFloor() {
	n.floor = 5
}

func (n *iglooBuilder) getHouse() House {
	return House {
		windownType: n.windownType,
		floor: n.floor,
	}
}