package main

import (
	"encoding/json"
	"syscall/js"
	"github.com/nguyentuio1/smart-layout-engine/core"
)

// Wrapper cho hàm Solve để JavaScript có thể gọi
func solveWrapper(this js.Value, args []js.Value) interface{} {
	// 1. Nhận dữ liệu JSON từ JS (Danh sách đồ đạc)
	inputJSON := args[0].String()
	var items []core.Entity
	json.Unmarshal([]byte(inputJSON), &items)

	// 2. Khởi tạo bộ giải (Solver)
	solver := core.SimpleSolver{
		Step: 10,
		Room: core.Dimension{W: 500, H: 500}, // Giả sử phòng 500x500
	}

	// 3. Giải toán
	result, err := solver.Solve(items)
	if err != nil {
		return js.ValueOf(map[string]interface{}{
			"error": err.Error(),
		})
	}

	// 4. Trả kết quả về dạng JSON string cho JS xử lý tiếp
	outputJSON, _ := json.Marshal(result)
	return js.ValueOf(string(outputJSON))
}

func main() {
	// Giữ cho chương trình Go chạy ngầm trên trình duyệt
	c := make(chan struct{}, 0)

	// Đăng ký hàm "solveInGo" với trình duyệt
	js.Global().Set("solveInGo", js.FuncOf(solveWrapper))

	println("Go WebAssembly đã sẵn sàng!")
	<-c
}
