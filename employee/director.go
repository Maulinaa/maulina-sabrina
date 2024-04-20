package employee

// Intern is a struct for intern
// It embeds Employee
// This means that Intern has all the fields and methods of Employee
type Director struct {
	Employee
}

// NewIntern creates a new intern
// It returns a pointer to the intern
// Creational method
func NewDirector() *Director {
	return &Director{
		Employee: Employee {
			Name : "Director",
			Salary : 5000,
		},
	}
	// TODO: Create a new intern
	// Set the name to "Intern"
	// Set the salary to 100
}

func (d *Director) GetBonus() float64 {
	return float64(d.Salary * 3 / 10)
}