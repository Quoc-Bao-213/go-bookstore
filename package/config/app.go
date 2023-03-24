package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB // Biến global để chứa instance của kết nối db
)

// Sử dụng `gorm.Open để kết nối db`.
// Database kết nối tới trong package là "bookstore" với user: "root", password rỗng, utf8 charset, parseTime là true và location được đặt ở local.
func Connect() {
	d, err := gorm.Open("mysql", "root:@/bookstore?charset=utf8&parseTime=True&loc=Local")

	if err != nil { // nil == null
		panic(err) // panic: ngừng thực hiện chương trình ngay lập tức và thông báo lỗi
	}

	db = d
}

// Trả về trực tiếp giá trị lưu trong biến db, đc sử dụng trong các hàm khác để lấy giá trị db (Biến chưa kết nối tới DB).
func GetDB() *gorm.DB {
	return db
}
