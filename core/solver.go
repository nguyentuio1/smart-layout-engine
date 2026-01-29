package core

import "errors"

type SimpleSolver struct {
	Step float64 // Độ chi tiết của mỗi lần thử (ví dụ: nhảy mỗi lần 10px)
	Room Dimension
}

func (s *SimpleSolver) Solve(entities []Entity) ([]Entity, error) {
	if s.backtrack(&entities, 0) {
		return entities, nil
	}
	return nil, errors.New("không tìm thấy phương án sắp xếp phù hợp")
}

func (s *SimpleSolver) backtrack(entities *[]Entity, index int) bool {
	// Nếu đã đặt xong tất cả đồ đạc
	if index == len((*entities)) {
		return IsValidLayout(*entities)
	}

	// Nếu vật thể này được cố định (IsFixed), bỏ qua và sang vật thể tiếp theo
	if (*entities)[index].IsFixed {
		return s.backtrack(entities, index+1)
	}

	// Thử từng tọa độ X, Y trong căn phòng
	for x := 0.0; x <= s.Room.W-(*entities)[index].Size.W; x += s.Step {
		for y := 0.0; y <= s.Room.H-(*entities)[index].Size.H; y += s.Step {

			// Đặt thử
			(*entities)[index].Position = Point{X: x, Y: y}

			// Kiểm tra nhanh: nếu đặt cái này vào mà đã đè lên các cái trước đó thì lướt qua luôn
			if s.isCurrentValid(*entities, index) {
				// Đệ quy để đặt vật thể tiếp theo
				if s.backtrack(entities, index+1) {
					return true
				}
			}
		}
	}

	return false
}

// Hàm bổ trợ kiểm tra vật thể hiện tại với các vật thể đã đặt trước đó
func (s *SimpleSolver) isCurrentValid(entities []Entity, index int) bool {
	for i := 0; i < index; i++ {
		if CheckOverlap(entities[i], entities[index]) {
			return false
		}
	}
	return true
}
