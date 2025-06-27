## 📘 Ghi chú: Kiểm Tra Tràn Số với `reflect.Value` trong Go

Trong Go, khi sử dụng reflection (`reflect.Value`) để thao tác với dữ liệu, chúng ta cần kiểm tra **giá trị có bị tràn (overflow)** trước khi gán. Go cung cấp các phương thức sau:

---

### 🔢 1. `OverflowFloat(val float64) bool`

- Trả về `true` nếu `val` vượt quá phạm vi lưu trữ của kiểu float (`Float32` hoặc `Float64`).
- Gây **panic** nếu `Value.Kind()` không phải là `reflect.Float32` hoặc `reflect.Float64`.

#### 📌 Ví dụ:

```go
val := reflect.ValueOf(float32(0))
fmt.Println(val.OverflowFloat(1e40)) // true vì 1e40 vượt giới hạn float32
```

---

### 🔢 2. `OverflowInt(val int64) bool`

- Trả về `true` nếu `val` vượt quá phạm vi lưu trữ của kiểu số nguyên có dấu (signed int).
- Gây **panic** nếu không phải kiểu `Int`, `Int8`, `Int16`, `Int32`, `Int64`.

#### 📌 Ví dụ:

```go
val := reflect.ValueOf(int8(0))
fmt.Println(val.OverflowInt(200)) // true vì int8 chỉ chứa tối đa 127
```

---

### 🔢 3. `OverflowUint(val uint64) bool`

- Trả về `true` nếu `val` vượt quá phạm vi của kiểu số nguyên không dấu.
- Gây **panic** nếu không phải kiểu `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`.

#### 📌 Ví dụ:

```go
val := reflect.ValueOf(uint8(0))
fmt.Println(val.OverflowUint(300)) // true vì uint8 tối đa là 255
```

---

### 📋 Bảng Tổng Hợp

| Hàm             | Áp dụng cho         | Giá trị truyền vào | Gây panic khi             | Trả về `true` khi                   |
| --------------- | ------------------- | ------------------ | ------------------------- | ----------------------------------- |
| `OverflowFloat` | Float32, Float64    | `float64`          | Không phải float          | Giá trị vượt phạm vi kiểu float     |
| `OverflowInt`   | Int, Int8..Int64    | `int64`            | Không phải kiểu có dấu    | Giá trị vượt phạm vi kiểu số nguyên |
| `OverflowUint`  | Uint, Uint8..Uint64 | `uint64`           | Không phải kiểu không dấu | Giá trị vượt phạm vi kiểu không dấu |

---

### ✅ Ghi nhớ:

- Luôn kiểm tra `OverflowXxx()` trước khi dùng `SetXxx()` với `reflect.Value`.
- Tránh panic runtime không mong muốn khi thao tác dữ liệu động.

---

### 📚 Tài liệu:

- [reflect.Value OverflowFloat](https://pkg.go.dev/reflect#Value.OverflowFloat)
- [reflect.Value OverflowInt](https://pkg.go.dev/reflect#Value.OverflowInt)
- [reflect.Value OverflowUint](https://pkg.go.dev/reflect#Value.OverflowUint)
