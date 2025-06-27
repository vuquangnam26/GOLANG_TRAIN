## 📘 Ghi chú về xử lý kết quả truy vấn trong Go (database/sql)

Dưới đây là hướng dẫn các hàm thường dùng để xử lý kết quả truy vấn trả về từ cơ sở dữ liệu trong Go.

---

### 🔄 Next()

- **Mục đích**: Di chuyển sang dòng dữ liệu kế tiếp trong kết quả trả về.
- **Kiểu trả về**: `bool`

  - `true`: nếu còn dòng dữ liệu.
  - `false`: nếu đã đến cuối dữ liệu. Khi đó `Close()` được tự động gọi.

```go
for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
    fmt.Println(id, name)
}
```

---

### 📦 NextResultSet()

- **Mục đích**: Duyệt qua các "result sets" khi truy vấn trả về nhiều kết quả (thường gặp trong stored procedure).
- **Kiểu trả về**: `bool`

```go
for rows.NextResultSet() {
    for rows.Next() {
        // Xử lý từng dòng trong từng result set
    }
}
```

---

### 🧪 Scan(...targets)

- **Mục đích**: Gán dữ liệu từ dòng hiện tại vào các biến Go.
- **Chú ý**: Số lượng và kiểu dữ liệu phải khớp với các cột trả về.

```go
rows.Scan(&id, &name)
```

---

### ❌ Close()

- **Mục đích**: Giải phóng tài nguyên liên quan đến truy vấn.
- **Ghi nhớ**: Nếu dùng `Next()` đến hết thì không cần gọi thủ công, nhưng vẫn nên dùng `defer rows.Close()` để an toàn.

```go
rows, _ := db.Query("SELECT id FROM users")
defer rows.Close()
```

---

### ✅ Tổng kết:

| Hàm               | Mục đích                           |
| ----------------- | ---------------------------------- |
| `Next()`          | Di chuyển sang dòng tiếp theo      |
| `NextResultSet()` | Duyệt qua các bộ kết quả tiếp theo |
| `Scan(...)`       | Đọc dữ liệu dòng hiện tại vào biến |
| `Close()`         | Đóng và giải phóng tài nguyên      |

---

> 💡 **Tip**: Dùng `defer rows.Close()` ngay sau khi gọi `Query` để tránh quên đóng tài nguyên.

## 📘 Giải thích hai hàm `Scan(...targets)` và `Err()` trong Go (package database/sql)

---

### 🔹 `Scan(...targets)`

- **Chức năng:** Gán giá trị từ các cột trong SQL row vào các biến truyền vào (phải là con trỏ).
- **Cách dùng:**

```go
var id int
var name string
row := db.QueryRow("SELECT id, name FROM users WHERE id = ?", 1)
err := row.Scan(&id, &name)
```

- **Lưu ý:**

  - Thứ tự biến truyền vào phải khớp với thứ tự cột trong câu SQL.
  - Kiểu dữ liệu phải tương đương (VD: column trong DB là VARCHAR thì biến phải là string).
  - Nếu không khớp số lượng hoặc kiểu, sẽ sinh ra lỗi.
  - Dùng cho `QueryRow()` hoặc dòng lần `for rows.Next()`

---

### 🔹 `Err()`

- **Chức năng:** Trả về lỗi phát sinh khi lặp qua tập kết quả `rows`.
- **Cách dùng:**

```go
rows, err := db.Query("SELECT * FROM products")
if err != nil {
  log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
  // Scan vào biến
}

if err := rows.Err(); err != nil {
  log.Fatal(err)
}
```

- **Lưu ý:** Phải gọi sau khi duyệt hết dữ liệu từ `rows.Next()` để kiểm tra lỗi tiềm ẩn trong quá trình duyệt (VD: DB connection ngắt quá trình).

---

### ✅ Tóm tắt:

| Tên hàm  | Dùng để                            | Lưu ý                                   |
| -------- | ---------------------------------- | --------------------------------------- |
| `Scan()` | Gán giá trị từ SQL row → biến      | Truyền vào con trỏ; khớp thứ tự và kiểu |
| `Err()`  | Kiểm tra lỗi sau khi dòng `Next()` | Nên gọi sau khi duyệt hết rows          |

---

Nếu bạn cần ví dụ nâng cao hơn (nhiều bảng, nested struct, mapping thủ công sang struct), mình có thể hỗ trợ mở rộng thêm.
