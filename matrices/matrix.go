package matrix

import (
	"errors"
	"fmt"
	"log"

	"github.com/willf/bitset"
)

func TransposeSqMatrix(m [][]int) error {

	l := len(m)
	w := len(m[0])

	size := l*w - 1
	b := bitset.New(uint(size))
	b.Set(0).Set(uint(size))

	for index := 1; index < size; index++ {

		next := (index * l) % size

		if b.Test(uint(next)) {
			continue
		}

		x, y, err := getCoordinates(index, l, w)
		if err != nil {
			return err
		}
		value := m[x][y]

		p, q, err := getCoordinates(next, l, w)
		if err != nil {
			return err
		}

		m[x][y] = m[p][q]
		m[p][q] = value

		b.Set(uint(index))

	}

	return nil

}

func getCoordinates(index, row, col int) (int, int, error) {

	for i := 0; i < row; i++ {
		if index < ((col*i)+col) && (index >= col*i) {
			return i, index - col*i, nil
		}
	}
	return 0, 0, errors.New("invalid index")
}

func TransposeNoneSqMatrix(m [][]int) error {

	l := len(m)
	w := len(m[0])

	size := l*w - 1

	b := bitset.New(uint(size))
	b.Set(0).Set(uint(size))
	log.Println(b.String())
	var next, start, index, value int

	index = 1

	for index < size && int(b.Count()) != size {
		fmt.Println(b.String())
		start = index
		fmt.Println("num of sets,size", b.Count(), size)
		for ok := true; ok; ok = (start != index) {
			fmt.Println("start,index,next", start, index, next)
			next = (index * l) % size
			if b.Test(uint(next)) {
				break
			}

			x, y, err := getCoordinates(index, l, w)
			if err != nil {
				return err
			}
			value = m[x][y]

			p, q, err := getCoordinates(next, l, w)
			if err != nil {
				return err
			}

			fmt.Println("swap ", m[x][y], m[p][q])

			m[x][y] = m[p][q]
			m[p][q] = value

			fmt.Println(b.String())
			b.Set(uint(index))

			index = next

		}
		fmt.Println(m)
		if n, ok := b.NextClear(uint(start)); ok {
			index = int(n)
			continue
		}
		fmt.Println(b.String())
		break
	}

	return nil
}

// transpose matrix for easy of multiplicaiton. cost O(n*m) for m, n dimensional matrix
func LazyTranspose(m [][]int) [][]int {

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
	m2 = LazyTranspose(m2)

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
