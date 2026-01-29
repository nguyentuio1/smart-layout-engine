package core

// CheckOverlap kiểm tra xem hai vật thể có đè lên nhau không (thuật toán AABB)
func CheckOverlap(a, b Entity) bool {
	return a.Position.X < b.Position.X+b.Size.W &&
		a.Position.X+a.Size.W > b.Position.X &&
		a.Position.Y < b.Position.Y+b.Size.H &&
		a.Position.Y+a.Size.H > b.Position.Y
}

// IsValidLayout kiểm tra xem toàn bộ danh sách đồ đạc có cái nào đè nhau không
func IsValidLayout(entities []Entity) bool {
	for i := 0; i < len(entities); i++ {
		for j := i + 1; j < len(entities); j++ {
			if CheckOverlap(entities[i], entities[j]) {
				return false
			}
		}
	}
	return true
}
