package main

import (
	"testing"
)

var num = 1000

func BenchmarkInsertBun(b *testing.B) {
	db = postgresSetUp()
	for i := 0; i < b.N; i++ {
		//primeNumbers(num)
		insertDep()
	}
}

func BenchmarkGetBun(b *testing.B) {
	db = postgresSetUp()
	for i := 0; i < b.N; i++ {
		//primeNumbers(num)
		getSpecificData()
	}
}
