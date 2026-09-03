package ports

import "errors"

var ErrStatisticsUnavailable = errors.New("statistics service unavailable")

type MatrixResult struct {
	RotatedMatrix [][]float64    `json:"rotatedMatrix"`
	Q             [][]float64    `json:"q"`
	R             [][]float64    `json:"r"`
	Statistics    map[string]any `json:"statistics"`
}

type StatisticsGateway interface {
	Calculate(rotated, q, r [][]float64) (map[string]any, error)
}
