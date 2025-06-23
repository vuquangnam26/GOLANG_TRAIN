## 📄 Xử lý Form và Upload File trong Go (net/http)

Đây là ví dụ và giải thích chi tiết từng dòng lệnh để xử lý dữ liệu form và file upload trong Go.

---

### 📦 Toàn bộ ví dụ:

```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Lỗi parse form", http.StatusBadRequest)
        return
    }

    username := r.FormValue("username")

    file, header, err := r.FormFile("profile")
    if err != nil {
        http.Error(w, "Lỗi lấy file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    fmt.Fprintf(w, "Username: %s\n", username)
    fmt.Fprintf(w, "Uploaded file: %s (%d bytes)", header.Filename, header.Size)
}
```

---

### 🧠 Giải thích từng dòng lệnh:

#### `err := r.ParseMultipartForm(10 << 20)`

- Dùng để phân tích form có kiểu `multipart/form-data` (dạng form upload file).
- `10 << 20` = 10 MB: giới hạn dung lượng đọc từ form.

#### `username := r.FormValue("username")`

- Lấy giá trị đầu tiên từ input form có `name="username"`.
- Tự động gọi `ParseForm()` nếu cần.

#### `file, header, err := r.FormFile("profile")`

- Lấy file đầu tiên được gửi từ input có `name="profile"`.
- Trả về: nội dung file, thông tin file (`header`), và lỗi nếu có.

#### `defer file.Close()`

- Đảm bảo đóng file sau khi xử lý xong.

#### `fmt.Fprintf(w, "Username: %s\n", username)`

- Ghi nội dung username vào phản hồi gửi về client.

#### `fmt.Fprintf(w, "Uploaded file: %s (%d bytes)", header.Filename, header.Size)`

- Hiển thị tên file và kích thước file được upload.

---

### 📌 Tổng hợp chức năng:

| Lệnh                   | Mục đích                                |
| ---------------------- | --------------------------------------- |
| `ParseMultipartForm()` | Phân tích dữ liệu từ form (upload file) |
| `FormValue()`          | Lấy dữ liệu từ form text input          |
| `FormFile()`           | Lấy nội dung và thông tin file upload   |
| `defer file.Close()`   | Giải phóng tài nguyên file              |
| `Fprintf()`            | Ghi kết quả phản hồi về client          |

---

### 💡 Ghi nhớ:

- Phải gọi `ParseMultipartForm()` trước khi dùng `FormFile()`.
- `FormValue()` tiện lợi hơn khi chỉ lấy text, vì tự gọi `ParseForm()`.
- Hạn chế upload file quá lớn, nên đặt giới hạn phù hợp.

---
