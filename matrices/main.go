package main

import (
	"errors"
	"fmt"
)

// transpose matrix for easy of multiplicaiton. cost O(n*m) for m, n dimensional matrix
func transpose(m [][]int) [][]int {

	result := make([][]int, len(m[0]))

	for i := range m[0] {
		result[i] = make([]int, len(m))
	}
	for i, v := range m {
		for j, n := range v {
			result[j][i] = n

		}
	}
	return result
}

// dotProduct returns dot product of two slices with equal length : cost O(n)
func dotProduct(m []int, n []int) int {

	sum := 0
	for i := range m {
		sum = sum + m[i]*n[i]
	}
	return sum
}

// for multiplication m1 must have m X n dimension, and m2 must have n X p dimension to return m X p product
// cost O(m *n)
func multiplyWithTranspose(m1 [][]int, m2 [][]int) ([][]int, error) {

	if m1 == nil || m2 == nil {
		return nil, errors.New("invalid matrix provided")
	}

	m := len(m1)
	p := len(m2[0])
	if len(m2) != len(m1[0]) {
		return nil, errors.New("matrices dimensions don't allow multiplication")
	}

	var product = make([][]int, m)
	// raw access is faster
	m2 = transpose(m2)

	for i, value1 := range m1 {
		product[i] = make([]int, p)
		for j, value2 := range m2 {

			product[i][j] = dotProduct(value1, value2)

		}
	}

	return product, nil

}

func Multiply(m1 [][]int, m2 [][]int) ([][]int, error) {

	if m1 == nil || m2 == nil {
		return nil, errors.New("invalid matrix provided")
	}

	m, n, p := len(m1), len(m2), len(m2[0])

	if len(m2) != len(m1[0]) {
		return nil, errors.New("matrices dimensions don't allow multiplication")
	}

	var product = make([][]int, m)

	for i := range product {
		product[i] = make([]int, p)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum = sum + m1[i][k]*m2[k][j]
				product[i][j] = sum
			}
		}
	}

	return product, nil
}

func main() {
	x := [][]int{
		{2, 0}, {1, 1}, {3, 2},
	}

	y := [][]int{
		{1, 3}, {4, 2},
	}

	p, err := Multiply(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
}
