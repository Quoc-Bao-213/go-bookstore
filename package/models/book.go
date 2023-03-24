package models

import (
	"github.com/akhil/go-bookstore/package/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB // Biến db chứa một instance giữ kết nối đến CSDL

// `type Book` định nghĩa kiểu dữ liệu Book để đại diện cho một object book.
// Biến `Book` là một struct bao gồm các trường lưu thông tin của một quyển sách.
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Hàm được gọi đầu tiên khi chương trình được thực thi.
// Hàm được sử dụng để khởi tạo và kết nối đến CSDL thông qua module `config`.
// Khởi tạo một instance của struct `Book`.
func init() {
	config.Connect()
	db = config.GetDB()
	// Tự động tạo bảng trong CSDL tương ứng với model `Book`
	// Khi cấu trúc của `Book` thay đổi, có thể chạy lại phương thức này để cập nhật bảng.
	db.AutoMigrate(&Book{})
}

// Một phương thức của `Book` dùng để tạo mới một quyển sách.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // Tạo một record mới trong CSDL và trả về một input book mới
	db.Create(&b) // tạo một record mới cho đối tượng `Book`
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
