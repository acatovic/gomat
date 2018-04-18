package gomat

import (
	"fmt"
	"testing"
)

//////////////
// Examples //
//////////////

func ExampleAdd() {
	ma := New([][]float64{{1,2},
		                  {3,4}})
	mb := New([][]float64{{5,6},
		                  {7,8}})
	mc := Add(ma, mb)
	fmt.Printf("Rows: %d\nCols: %d\nData: %v", mc.m, mc.n, mc.data)
	// Output:
	// Rows: 2
	// Cols: 2
	// Data: [6 8 10 12]
}

func ExampleDot() {
	ma := New([][]float64{{1,2},
						  {3,4},
						  {5,6}})
	mb := New([][]float64{{1,2,3},
						  {4,5,6}})
	mc := Dot(ma, mb)
	fmt.Printf("Rows: %d\nCols: %d\nData: %v", mc.m, mc.n, mc.data)
	// Output:
	// Rows: 3
	// Cols: 3
	// Data: [9 12 15 19 26 33 29 40 51]
}

func ExampleSub() {
	ma := New([][]float64{{1,2},
		                  {3,4}})
	mb := New([][]float64{{5,6},
		                  {7,8}})
	mc := Sub(ma, mb)
	fmt.Printf("Rows: %d\nCols: %d\nData: %v", mc.m, mc.n, mc.data)
	// Output:
	// Rows: 2
	// Cols: 2
	// Data: [-4 -4 -4 -4]
}

func ExampleTranspose() {
	ma := New([][]float64{{1,2,3},
						  {4,5,6}})
	mb := Transpose(ma)
	fmt.Printf("Rows: %d\nCols: %d\nData: %v", mb.m, mb.n, mb.data)
	// Output:
	// Rows: 3
	// Cols: 2
	// Data: [1 4 2 5 3 6]
}

func ExampleZeros() {
	ma := Zeros(3,2)
	fmt.Printf("Rows: %d\nCols: %d\nData: %v", ma.m, ma.n, ma.data)
	// Output:
	// Rows: 3
	// Cols: 2
	// Data: [0 0 0 0 0 0]
}

////////////////
// Benchmarks //
////////////////

func BenchmarkAdd(b *testing.B) {
	ma := Randn(1024, 1024)
	mb := Randn(1024, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Add(ma, mb)
	}
}

func BenchmarkDot(b *testing.B) {
	ma := Randn(1024, 1024)
	mb := Randn(1024, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot(ma, mb)
	}
}

func BenchmarkSub(b *testing.B) {
	ma := Randn(1024, 1024)
	mb := Randn(1024, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sub(ma, mb)
	}
}