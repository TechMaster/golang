package large

import (
	"testing"
)

var p = Person{
	id:      "OX13",
	name:    "Valya",
	email:   "valya@gmail.com",
	pass:    "2334234j,m as",
	roles:   []string{"admin", "manager"},
	age:     18,
	enabled: false,
}

func BenchmarkPassPersonAsValue(b *testing.B) {

	for i := 0; i < b.N; i++ {
		PassPersonAsValue(p)
	}
}

func BenchmarkPassPersonAsPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassPersonAsPointer(&p)
	}
}
