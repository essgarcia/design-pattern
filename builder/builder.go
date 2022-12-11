package builder

type human struct {
	name   string
	age    int
	height float64
}

// NewHuman Empty constructor
func NewHuman() human {
	return human{}
}

// CompleteNewHuman Full constructors allow you to pass all properties once
func CompleteNewHuman(
	name string,
	age int,
	height float64) human {
	return human{
		name:   name,
		age:    age,
		height: height,
	}
}

// Builders
func (h human) Named(name string) human {
	h.name = name
	return h
}

func (h human) WithAge(age int) human {
	h.age = age
	return h
}

func (h human) WithHeight(height float64) human {
	h.height = height
	return h
}

// Giant Pre-defined builders: when you know you're gonna build frequently the same set of properties
func Giant() human {
	return NewHuman().Named("Onias").WithHeight(2.5)
}
