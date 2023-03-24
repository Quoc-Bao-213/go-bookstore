package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Nhận tham số là một đối tượng request `*http.Request` và một biến `X` có kiểu interface{}.
// Hàm được sử dụng để phân tích nội dung của request, đọc dữ liệu từ ngoài vào, phân tích nội dung đó từ JSON và gán cho một object X tương ứng.
func ParseBody(r *http.Request, X interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil { // Đọc request body `ioutil.ReadAll(r.Body)`, `r` là một instance của `http.Request`.
		if err := json.Unmarshal([]byte(body), X); err != nil { // Hàm thực hiện phân tích JSON và gán giá trị cho biến `X` đối với dữ liệu vừa được đọc
			return
		}
	}
}
