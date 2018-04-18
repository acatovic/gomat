// Copyright 2018 Armin Catovic. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package gomat is a simple matrix implemenation.
// At the core is the Matrix struct, which simply
// wraps a slice of floats, i.e. []float64.
// We make a light attempt at optimising cache locality;
// however it is by no means the most optimal implementation.
//
// Example - Adding two matrices:
//  ma := gomat.Randn(16, 4)
//  mb := gomat.Randn(16, 4)
//  mc := gomat.Add(ma, mb)
//  fmt.Println(mc)
//
// Example - Dot product
//  ma := gomat.New([][]float64{{1,2},{3,4},{5,6}})
//  mb := gomat.New([][]float64{{1,2},{3,4}})
//  mc := gomat.Dot(ma, mb)
//  fmt.Println(mc)
package gomat

import (
	"math"
	"math/rand"
)

type Matrix struct {
	m int
	n int
	data []float64
}

// Private Matrix methods

func (mat *Matrix) index_at(i, j int) int {
	return i * mat.n + j
}

func (mat *Matrix) value_at(i, j int) float64 {
	return mat.data[i * mat.n + j]
}

// Public Matrix methods

func (mat *Matrix) Cols() int {
	return mat.n
}

func (mat *Matrix) Rows() int {
	return mat.m
}

// Private helper functions used by matrix operations

func add_vec(va, vb, vc []float64) {
	for i := range va {
		vc[i] = va[i] + vb[i]
	}
}

func dot_vec(va, vb []float64) float64 {
	res := 0.0
	for i := range va {
		res += va[i] * vb[i]
	}
	return res
}

func sigmoid(val float64) float64 {
	return 1.0 / (1.0 + math.Exp((-1.0 * val)))
}

func sigmoid_prime(val float64) float64 {
	return sigmoid(val) * (1 - sigmoid(val))
}

func sub_vec(va, vb, vc []float64) {
	for i := range va {
		vc[i] = va[i] - vb[i]
	}
}

// Public functions that implement matrix operations

// Add performs addition of two matrices 'ma' and 'mb' and
// returns a ptr to the resulting matrix
func Add(ma, mb *Matrix) *Matrix {
	if ma.m != mb.m && ma.n != mb.n {
		panic("Dimensions of matrix A and matrix B must be equal")
	}
	mc := &Matrix{ma.m, ma.n, make([]float64, ma.m * ma.n)}
	if ma.n > 15 {
		r, c := ma.m, ma.n
		for i := 0; i < r; i++ {
			add_vec(ma.data[i*c:i*c+c],
				mb.data[i*c:i*c+c], mc.data[i*c:i*c+c])
		}
	} else {
		for i := 0; i < ma.m * ma.n; i++ {
			mc.data[i] = ma.data[i] + mb.data[i]
		}
	}
	return mc
}

// Dot performs dot product of matrices 'ma' and 'mb' and
// returns a ptr to the resulting matrix. Each row in 'ma' is
// multiplied by column in 'mb'
func Dot(ma, mb *Matrix) *Matrix {
	if ma.n != mb.m {
		panic("Num cols in matrix A must be equal to num rows in matrix B")
	}
	mc := &Matrix{ma.m, mb.n, make([]float64, ma.m * mb.n)}
	tr_mb := Transpose(mb)
	for i := 0; i < mc.m; i++ {
		for j := 0; j < tr_mb.m; j++ {
			mc.data[mc.index_at(i, j)] = dot_vec(ma.data[i*ma.n:i*ma.n+ma.n],
				tr_mb.data[j*tr_mb.n:j*tr_mb.n+tr_mb.n])
		}
	}
	return mc
}

// FromVec takes a slice of M elements and
// returns a ptr to a Mx1 matrix
func FromVec(vec []float64) *Matrix {
	if len(vec) == 0 {
		panic("Empty input vector")
	}
	return &Matrix{len(vec), 1, vec}
}

// Mul applies Hadamard product between matrix 'ma' and 'mb'
// and returns a ptr to the resulting matrix
func Mul(ma, mb *Matrix) *Matrix {
	if ma.m != mb.m && ma.n != mb.n {
		panic("Dimensions of matrix A and matrix B must be equal")
	}
	mc := &Matrix{ma.m, ma.n, make([]float64, ma.m * ma.n)}
	for i := 0; i < ma.m * ma.n; i++ {
		mc.data[i] = ma.data[i] * mb.data[i]
	}
	return mc
}

// New returns a ptr to a MxN matrix using manually-inputted floats
// specified by 'd'
func New(d [][]float64) *Matrix {
	m := len(d)
	if m == 0 {
		panic("No rows defined")
	}
	n := len(d[0])
	if n == 0 {
		panic("No columns defined")
	}
	mat := &Matrix{m, n, make([]float64, m*n)}
	for i := 0; i < m; i++ {
		if len(d[i]) != n {
			panic("All rows must be equal length")
		}
		for j := 0; j < n; j++ {
			mat.data[mat.index_at(i, j)] = d[i][j]
		}
	}
	return mat
}

// Randn returns a ptr to a m x n matrix with random
// normally distributed values, i.e. mean=0, stdev=1
func Randn(m, n int) *Matrix {
	mat := &Matrix{m, n, make([]float64, m * n)}
	for i := 0; i < m * n; i++ {
		mat.data[i] = rand.NormFloat64()
	}
	return mat
}

// Sigmoid applies the sigmoid function element-wise
// on matrix 'ma' and returns a ptr to the resulting matrix
func Sigmoid(ma *Matrix) *Matrix {
	mb := &Matrix{ma.m, ma.n, make([]float64, ma.m * ma.n)}
	for i, val := range ma.data {
		mb.data[i] = sigmoid(val)
	}
	return mb
}

// Sigmoidpr applies the sigmoid derivative function element-wise
// on matrix 'ma' and returns a ptr to the resulting matrix
func Sigmoidpr(ma *Matrix) *Matrix {
	mb := &Matrix{ma.m, ma.n, make([]float64, ma.m * ma.n)}
	for i, val := range ma.data {
		mb.data[i] = sigmoid_prime(val)
	}
	return mb
}

// Scale multiplies scalar 'v' by matrix 'ma' and returns
// a ptr to the resulting matrix
func Scale(v float64, ma *Matrix) *Matrix {
	mb := &Matrix{ma.m, ma.n, make([]float64, ma.m * ma.n)}
	for i, val := range ma.data {
		mb.data[i] = val * v
	}
	return mb
}

// Sub subtracts 'mb' from matrix 'ma' and returns
// a ptr to the resulting matrix
func Sub(ma, mb *Matrix) *Matrix {
	if ma.m != mb.m && ma.n != mb.n {
		panic("Dimensions of matrix A and matrix B must be equal")
	}
	mc := &Matrix{ma.m, ma.n, make([]float64, ma.m * ma.n)}
	if ma.n > 15 {
		r, c := ma.m, ma.n
		for i := 0; i < r; i++ {
			sub_vec(ma.data[i*c:i*c+c],
				mb.data[i*c:i*c+c], mc.data[i*c:i*c+c])
		}
	} else {
		for i := 0; i < ma.m * ma.n; i++ {
			mc.data[i] = ma.data[i] - mb.data[i]
		}
	}
	return mc
}

// Transpose performs a transpose on matrix 'ma' and returns
// a ptr to the resulting matrix
func Transpose(ma *Matrix) *Matrix {
	mb := &Matrix{ma.n, ma.m, make([]float64, ma.m * ma.n)}
	for j := 0; j < ma.n; j++ {
		for i := 0; i < ma.m; i++ {
			mb.data[j * mb.n + i] = ma.value_at(i, j)
		}
	}
	return mb
}

// Zeros returns a ptr to a m x n matrix with zero-initialised elements
func Zeros(m, n int) *Matrix {
	return &Matrix{m, n, make([]float64, m * n)}
}