package factory

type animal interface {
	Sound() string
}

// Cat
type cat struct {
}

func newCat() cat {
	return cat{}
}

func (c cat) Sound() string {
	return "meow"
}

// Dog
type dog struct {
	name string
}

func newDog() *dog {
	return &dog{name: "Alfredo"}
}

func (d dog) Sound() string {
	return "woof"
}

func Farm(x int) string {
	var a animal
	if x > 40 {
		a = newCat()
	} else {
		a = newDog()
	}
	return a.Sound()
}
