package data

type Employee struct {
	Id   int
	Age  int
	City string
	Namn string
}

func (emp Employee) IsCool() bool {
	if emp.Namn == "Stefan" {
		return true
	}
	return false
}

func IsCool(emp Employee) bool {
	if emp.Namn == "Stefan" {
		return true
	}
	return false
}

type Team struct {
	Id          int
	FoundedYear int
	City        string
	Name        string
}
type Player struct {
	Id           int
	Name         string
	TeamId       int
	JerseyNumber int
}
type CityName struct {
	City string
	Namn string
}
