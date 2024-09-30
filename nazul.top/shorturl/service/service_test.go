package service_test

import (
	"fmt"
	"shorturl/service"
	"testing"
)

func BenchmarkGenerate(b *testing.B) {
	url := "http://www.google.com"
	impl := service.BuildMapService()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := impl.Generate(url)
		if err != nil {
			b.Error(err)
		} else {
			b.Log(m)
		}
	}
}

func BenchmarkSpringf(b *testing.B) {
	number := 8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}
