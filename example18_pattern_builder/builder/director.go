package builder

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) BuildHouse() House {
	d.builder.setWindownType()
	d.builder.setFloor()

	return d.builder.getHouse()
}