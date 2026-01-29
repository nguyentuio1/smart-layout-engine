package core

import "testing"

func TestSimpleSolver_Solve(t *testing.T) {
	solver := SimpleSolver{
		Step: 10,
		Room: Dimension{W: 100, H: 100},
	}

	items := []Entity{
		{ID: "A", Size: Dimension{W: 40, H: 40}},
		{ID: "B", Size: Dimension{W: 40, H: 40}},
	}

	result, err := solver.Solve(items)
	if err != nil {
		t.Errorf("Thuật toán thất bại: %v", err)
	}

	if !IsValidLayout(result) {
		t.Error("Kết quả không hợp lệ: Các vật thể bị đè lên nhau!")
	}
}
