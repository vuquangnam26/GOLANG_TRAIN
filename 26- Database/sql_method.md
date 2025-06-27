# 📄 Tổng hợp: Làm việc với `Rows.Columns()` và `Rows.ColumnTypes()` trong Go (database/sql)

Dưới đây là tài liệu tổng hợp các khái niệm, giải thích và ví dụ thực tế về cách lấy metadata (thông tin cấu trúc) của kết quả truy vấn trong Go bằng `database/sql`.

---

## ✅ `Columns()`

### Mô tả:

- Trả về danh sách tên các cột có trong kết quả truy vấn.
- Kiểu trả về: `[]string`

### Ví dụ:

```go
rows, _ := db.Query("SELECT id, name, price FROM products")
columns, _ := rows.Columns()
fmt.Println("Tên các cột:", columns)
```

---

## ✅ `ColumnTypes()`

### Mô tả:

- Trả về danh sách các con trỏ `*sql.ColumnType`
- Cho phép lấy được thông tin như:

  - Tên cột
  - Kiểu dữ liệu trong DB (VARCHAR, INT, etc.)
  - Có nullable hay không
  - Chiều dài, độ chính xác
  - Kiểu dữ liệu tương ứng trong Go

### Ví dụ:

```go
rows, _ := db.Query("SELECT id, name, price FROM products")
colTypes, _ := rows.ColumnTypes()
for _, col := range colTypes {
    fmt.Println("Tên cột:", col.Name())
    fmt.Println("Kiểu trong DB:", col.DatabaseTypeName())
    nullable, _ := col.Nullable()
    fmt.Println("Có thể NULL:", nullable)
    length, ok := col.Length()
    if ok {
        fmt.Println("Chiều dài:", length)
    }
    fmt.Println("Kiểu Go:", col.ScanType())
    fmt.Println("---")
}
```

---

## 🧠 Tóm tắt các phương thức của `*sql.ColumnType`

| Phương thức          | Trả về               | Mô tả                 |
| -------------------- | -------------------- | --------------------- |
| `Name()`             | `string`             | Tên cột               |
| `DatabaseTypeName()` | `string`             | Kiểu dữ liệu trong DB |
| `Nullable()`         | `bool, bool`         | Có cho NULL không     |
| `DecimalSize()`      | `int64, int64, bool` | precision, scale      |
| `Length()`           | `int64, bool`        | độ dài                |
| `ScanType()`         | `reflect.Type`       | Kiểu Go tương ứng     |

---

## 📌 Ghi chú thêm:

- Dùng `Columns()` khi bạn chỉ cần tên cột để hiển thị hoặc xử lý linh hoạt.
- Dùng `ColumnTypes()` khi bạn cần thông tin chi tiết để làm việc với dữ liệu một cách động (ví dụ dùng `reflect`).
- Nếu cần scan dữ liệu tự động vào struct: nên dùng thư viện như [sqlx](https://github.com/jmoiron/sqlx) để giảm thiểu lỗi và tối ưu hiệu suất.

---

## ✅ Những gì đã trao đổi:

- Phân biệt rõ giữa `Columns()` và `ColumnTypes()`
- Ý nghĩa của các phương thức như `ScanType`, `Nullable`, `Length`, `DecimalSize`
- Giải thích khái niệm reflection
- Cảnh báo khi dùng reflection không đúng cách
- Gợi ý dùng thư viện `sqlx` nếu cần mapping tự động

---

📌 **Bạn có thể dùng tài liệu này để ôn tập trước phỏng vấn hoặc làm tài liệu nội bộ cho team phát triển backend với Go.**
