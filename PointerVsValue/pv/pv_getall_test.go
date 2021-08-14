package pv

import (
	"pointer_value/pointer"
	"pointer_value/value"
	"testing"
)

var pAccRepo = pointer.AccountRepository
var vAccRepo = value.AccountRepository
var accNewP = pointer.AccountNew{
	Email:    "test@example.com",
	FullName: "Test Name",
	Password: "Real Password",
}

var accNewV = value.AccountNew{
	Email:    "test@example.com",
	FullName: "Test Name",
	Password: "Real Password",
}

func Benchmark_GetAllPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = pAccRepo.GetAll()
	}
}

func Benchmark_GetAllValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = vAccRepo.GetAll()
	}
}

//------
func Benchmark_ValidatePointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = accNewP.Validate()
	}
}

func Benchmark_ValidateValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = accNewV.Validate()
	}
}

//------
func Benchmark_SavePointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = pAccRepo.Save(&accNewP)
	}
}

func Benchmark_SaveValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = vAccRepo.Save(accNewV)
	}
}

//----------------------------------------------------------------
func Benchmark_GetIdPointer(b *testing.B) {
	id, _ := pAccRepo.Save(&accNewP)
	for i := 0; i < b.N; i++ {
		_, _ = pAccRepo.GetById(id)
	}
}
func Benchmark_GetIdValue(b *testing.B) {
	id, _ := vAccRepo.Save(accNewV)
	for i := 0; i < b.N; i++ {
		_, _ = vAccRepo.GetById(id)
	}
}
