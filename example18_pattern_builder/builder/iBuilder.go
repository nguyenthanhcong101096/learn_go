package builder

type IBuilder interface {
	setWindownType()
	setFloor()
	getHouse() House
}

func GetBuilder(buildType string) IBuilder {
	switch buildType {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &normalBuilder{}
	}

	return nil
}