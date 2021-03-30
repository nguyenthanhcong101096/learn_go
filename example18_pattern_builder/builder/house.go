package builder

type House struct {
	windownType string
	floor int
}

func (h House) GetWindownType() string {
	return h.windownType
}

func (h House) GetFloor() int {
	return h.floor
}
