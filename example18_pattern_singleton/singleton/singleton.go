package singleton

type Singleton interface {
	AddItem() int
}

type singleton struct {
	count int
}

var instance *singleton

func (s *singleton) AddItem() int {
	s.count++
	return s.count
}

func GetInstance() Singleton {
	if instance == nil {
		instance = &singleton{count: 100}
	}

	return instance
}