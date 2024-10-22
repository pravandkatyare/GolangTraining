package employee

type Employer interface {
	GetName() string
	SetName(string)
}

type Employee struct {
	name string
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) SetName(name string) {
	e.name = name
}
