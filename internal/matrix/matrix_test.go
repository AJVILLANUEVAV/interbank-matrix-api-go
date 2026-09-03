package matrix

import "testing"

func TestRotateClockwise(t *testing.T) {
	input := [][]float64{{1, 2, 3}, {4, 5, 6}}
	want := [][]float64{{4, 1}, {5, 2}, {6, 3}}
	got := RotateClockwise(input)
	assertMatrixEqual(t, got, want)
}

func TestValidateRejectsJaggedMatrix(t *testing.T) {
	if err := Validate([][]float64{{1, 2}, {3}}); err == nil {
		t.Fatal("expected jagged matrix to be rejected")
	}
}

func TestQRReconstructsMatrix(t *testing.T) {
	input := [][]float64{{12, -51, 4}, {6, 167, -68}, {-4, 24, -41}}
	q, r, err := QR(input)
	if err != nil {
		t.Fatalf("QR returned error: %v", err)
	}
	for row := range input {
		for column := range input[0] {
			value := 0.0
			for index := range q[0] {
				value += q[row][index] * r[index][column]
			}
			if difference := value - input[row][column]; difference > 1e-8 || difference < -1e-8 {
				t.Fatalf("Q*R[%d][%d] = %v, want %v", row, column, value, input[row][column])
			}
		}
	}
}

func assertMatrixEqual(t *testing.T, got, want [][]float64) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %d rows, want %d", len(got), len(want))
	}
	for row := range want {
		if len(got[row]) != len(want[row]) {
			t.Fatalf("row %d has %d columns, want %d", row, len(got[row]), len(want[row]))
		}
		for column := range want[row] {
			if got[row][column] != want[row][column] {
				t.Fatalf("got [%d][%d] = %v, want %v", row, column, got[row][column], want[row][column])
			}
		}
	}
}
