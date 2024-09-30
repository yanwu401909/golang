package utils_test

import (
	"shorturl/utils"
	"testing"
)

func TestRandStr(t *testing.T) {
	v := utils.RandStr(8)
	if len(v) != 8 {
		t.Fatal("size error~")
	}
}

func BenchmarkRandStr(b *testing.B) {
	size := 8
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := utils.RandStr(size)
		// b.Log(">>", v)
		if len(v) != size {
			b.Fatal("size error!")
		}
	}
}
