package large

type Person struct {
	id      string
	name    string
	email   string
	pass    string
	roles   []string
	age     int
	enabled bool
}

func PassPersonAsValue(p Person) {
	p.id = "100"
	p.email = "cuong@techmaster.vn"
	p.pass = "2221432sdv"
	p.name = "John"
	p.roles = []string{"admin", "trainer"}
	p.age = 50
	p.enabled = true
}

func PassPersonAsPointer(p *Person) {
	p.id = "100"
	p.email = "cuong@techmaster.vn"
	p.pass = "2221432sdv"
	p.name = "John"
	p.roles = []string{"admin", "trainer"}
	p.age = 50
	p.enabled = true
}
