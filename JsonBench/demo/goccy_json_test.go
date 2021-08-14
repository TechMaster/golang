package demo

import (
	"testing"
)

func Benchmark_Goccy_JSON2Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Goccy_JSON2Struct(StringJSON)
	}
}

func Benchmark_Goccy_JSON2Map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Goccy_JSON2Map(StringJSON)
	}
}

func Benchmark_Default_JSON2Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Default_JSON2Struct(StringJSON)
	}
}

func Benchmark_Default_JSON2Map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Default_JSON2Map(StringJSON)
	}
}

//----------------------------------------------------------------
func Benchmark_Jsoniter_JSON2Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Jsoniter_JSON2Struct(StringJSON)
	}
}

func Benchmark_Jsoniter_JSON2Map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Jsoniter_JSON2Map(StringJSON)
	}
}
