package builder

type normalBuilder struct{
	windownType string
	floor int
}

func NewNormalBuidler() *normalBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder)setWindownType() {
	n.windownType = "normal windown type"
}

func (n *normalBuilder)setFloor() {
	n.floor = 3
}

func (n *normalBuilder)getHouse() House {
	return House {
		windownType: n.windownType,
		floor: n.floor,
	}
}