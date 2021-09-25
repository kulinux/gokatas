package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m *Matrix) Rows() [][]int {
	res := make([][]int, len(*m))
	for i, row := range *m {
		nrow := make([]int, len(row))
		copy(nrow, row)
		res[i] = nrow
	}
	return res
}

func (m *Matrix) Cols() [][]int {
	if len(*m) == 0 || len((*m)[0]) == 0 {
		return make([][]int, 0)
	}
	res := make([][]int, len((*m)[0]))
	for i := range (*m)[0] {
		res[i] = make([]int, len(*m))
	}
	for i, row := range *m {
		for j, cell := range row {
			res[j][i] = cell
		}
	}
	return res
}

func (m *Matrix) Set(r int, c int, val int) bool {
	if r < 0 || r >= len(*m) {
		return false
	}
	if c < 0 || c >= len((*m)[0]) {
		return false
	}
	(*m)[r][c] = val
	return true
}

//New export
func New(str string) (Matrix, error) {
	lines := strings.FieldsFunc(str, func(r rune) bool {
		return r == '\n'
	})

	if strings.Count(str, "\n")+1 != len(lines) {
		return nil, errors.New("empty rows")
	}

	res := make(Matrix, len(lines))

	for i, line := range lines {
		cols := strings.Fields(line)
		res[i] = make([]int, 0)
		for _, c := range cols {
			cInt, err := strconv.Atoi(c)
			if err != nil {
				return nil, errors.New("Int Overflow")
			}
			res[i] = append(res[i], cInt)
		}
	}

	currLength := 0
	for i, rr := range res {
		if i == 0 {
			currLength = len(rr)
			continue
		}
		if currLength != len(rr) {
			return nil, errors.New("uneven rows")
		}
	}

	return res, nil
}
