package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Khởi tạo Gin với cấu hình mặc định
	r := gin.Default()

	// 2. Phục vụ các file tĩnh trong thư mục frontend (HTML, JS, WASM)
	// Khi bạn truy cập http://localhost:8080, nó sẽ mở index.html
	r.StaticFS("/app", http.Dir("frontend"))

	// 3. Tạo một API endpoint đơn giản để kiểm tra
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Backend của bạn đã sẵn sàng!",
		})
	})

	// 4. Endpoint để lưu kết quả layout (sau này sẽ kết nối Database)
	r.POST("/save-layout", func(c *gin.Context) {
		// Ở đây chúng ta sẽ nhận dữ liệu từ Frontend gửi lên
		c.JSON(http.StatusOK, gin.H{
			"status": "Đã nhận dữ liệu và sẵn sàng lưu trữ!",
		})
	})

	// Chạy server tại cổng 8080
	r.Run(":8080")
}
