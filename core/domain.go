package core

// Point đại diện cho tọa độ (x, y)
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Dimension đại diện cho kích thước (rộng, cao)
type Dimension struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

// Entity là một vật thể cần sắp xếp (ví dụ: cái bàn, cái giường)
type Entity struct {
	ID       string    `json:"id"`
	Size     Dimension `json:"size"`
	Position Point     `json:"position"` // Máy sẽ tìm cái này
	IsFixed  bool      `json:"is_fixed"` // Nếu true, máy không được tự ý dời đi
	
}
// Constraint định nghĩa một luật lệ (ví dụ: A phải cách B một khoảng d)
type Constraint interface {
	IsValid(entities []Entity) bool
}

// Solver là interface cho bộ máy giải toán
type Solver interface {
	Solve(entities []Entity, constraints []Constraint) ([]Entity, error)
}
