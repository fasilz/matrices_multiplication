package matrix

import (
	"testing"
)

var product [][]int

func BenchmarkMultiplyWithTranspose(b *testing.B) {

	var p [][]int
	x := [][]int{
		{2, 0, 3, 9, 5, 5, 5, 5, 5, 5}, {1, 1, 1, 5, 1, 1, 1, 1, 5, 1}, {3, 2, 2, 2, 2, 3, 2, 2, 2, 2},
	}

	y := [][]int{
		{1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9},
	}

	for i := 0; i < b.N; i++ {
		p, _ = multiplyWithTranspose(x, y)
	}

	product = p

}

func BenchmarkMultiply(b *testing.B) {

	var p [][]int
	x := [][]int{
		{2, 0, 3, 9, 5, 5, 5, 5, 5, 5}, {1, 1, 1, 5, 1, 1, 1, 1, 5, 1}, {3, 2, 2, 2, 2, 3, 2, 2, 2, 2},
	}

	y := [][]int{
		{1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9}, {1, 3, 3, 4, 5, 6, 7, 8, 9},
	}

	for i := 0; i < b.N; i++ {
		p, _ = Multiply(x, y)
	}

	product = p

}

func Test_getCoordinates(t *testing.T) {
	type args struct {
		index int
		row   int
		col   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{"(0,0)", args{0, 3, 2}, 0, 0, false},
		{"(1,0)", args{2, 3, 2}, 1, 0, false},
		{"(3,0)", args{4, 3, 2}, 2, 0, false},
		{"(3,1)", args{5, 3, 2}, 2, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getCoordinates(tt.args.index, tt.args.row, tt.args.col)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCoordinates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCoordinates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getCoordinates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
