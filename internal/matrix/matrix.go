package matrix

import (
	"fmt"
	"math"
)

func Validate(input [][]float64) error {
	if len(input) == 0 || len(input[0]) == 0 {
		return fmt.Errorf("matrix must not be empty")
	}
	columns := len(input[0])
	for rowIndex, row := range input {
		if len(row) != columns {
			return fmt.Errorf("matrix must be rectangular: row %d has %d columns, expected %d", rowIndex, len(row), columns)
		}
	}
	return nil
}

func RotateClockwise(input [][]float64) [][]float64 {
	rows, columns := len(input), len(input[0])
	rotated := make([][]float64, columns)
	for column := 0; column < columns; column++ {
		rotated[column] = make([]float64, rows)
		for row := 0; row < rows; row++ {
			rotated[column][rows-1-row] = input[row][column]
		}
	}
	return rotated
}

func QR(input [][]float64) ([][]float64, [][]float64, error) {
	if err := Validate(input); err != nil {
		return nil, nil, err
	}
	rows, columns := len(input), len(input[0])
	rank := rows
	if columns < rank {
		rank = columns
	}
	working := clone(input)
	qFull := identity(rows)

	for step := 0; step < rank; step++ {
		norm := 0.0
		for row := step; row < rows; row++ {
			norm += working[row][step] * working[row][step]
		}
		if norm == 0 {
			continue
		}
		norm = math.Sqrt(norm)
		sign := 1.0
		if working[step][step] < 0 {
			sign = -1
		}
		vector := make([]float64, rows-step)
		for row := step; row < rows; row++ {
			vector[row-step] = working[row][step]
		}
		vector[0] += sign * norm
		vectorNorm := 0.0
		for _, value := range vector {
			vectorNorm += value * value
		}
		if vectorNorm == 0 {
			continue
		}
		vectorNorm = math.Sqrt(vectorNorm)
		for index := range vector {
			vector[index] /= vectorNorm
		}

		applyLeftReflector(working, vector, step, step)
		applyRightReflector(qFull, vector, step, 0)
	}

	q := make([][]float64, rows)
	for row := range q {
		q[row] = make([]float64, rank)
		copy(q[row], qFull[row][:rank])
	}
	r := make([][]float64, rank)
	for row := range r {
		r[row] = make([]float64, columns)
		copy(r[row], working[row])
	}
	return q, r, nil
}

func clone(input [][]float64) [][]float64 {
	output := make([][]float64, len(input))
	for row := range input {
		output[row] = append([]float64(nil), input[row]...)
	}
	return output
}

func identity(size int) [][]float64 {
	output := make([][]float64, size)
	for index := range output {
		output[index] = make([]float64, size)
		output[index][index] = 1
	}
	return output
}

func applyLeftReflector(target [][]float64, vector []float64, startRow, startColumn int) {
	for column := startColumn; column < len(target[0]); column++ {
		dot := 0.0
		for row := startRow; row < len(target); row++ {
			dot += vector[row-startRow] * target[row][column]
		}
		for row := startRow; row < len(target); row++ {
			target[row][column] -= 2 * vector[row-startRow] * dot
		}
	}
}

func applyRightReflector(target [][]float64, vector []float64, startColumn, startRow int) {
	for row := startRow; row < len(target); row++ {
		dot := 0.0
		for column := startColumn; column < len(target[0]); column++ {
			dot += target[row][column] * vector[column-startColumn]
		}
		for column := startColumn; column < len(target[0]); column++ {
			target[row][column] -= 2 * dot * vector[column-startColumn]
		}
	}
}

