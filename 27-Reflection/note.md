## 🧠 reflect.Type trong Golang

Struct `reflect.Type` đại diện cho thông tin về kiểu dữ liệu trong Go tại runtime. Dưới đây là giải thích chi tiết về các phương thức và hằng số `Kind` trong package `reflect`.

---

### 📘 Các phương thức phổ biến của `reflect.Type`

#### 🔹 `Name()`

- **Mô tả:** Trả về tên của kiểu dữ liệu (ví dụ: "int", "MyStruct").
- **Ví dụ:**

```go
reflect.TypeOf(123).Name() // "int"
```

#### 🔹 `PkgPath()`

- **Mô tả:** Trả về đường dẫn gói chứa kiểu dữ liệu.
- **Ví dụ:**

```go
t := reflect.TypeOf(time.Time{})
fmt.Println(t.PkgPath()) // "time"
```

#### 🔹 `Kind()`

- **Mô tả:** Trả về loại của kiểu dữ liệu (reflect.Kind).
- **Ví dụ:**

```go
reflect.TypeOf(123).Kind() // reflect.Int
reflect.TypeOf([]int{}).Kind() // reflect.Slice
```

#### 🔹 `String()`

- **Mô tả:** Trả về tên kiểu dữ liệu có package nếu có.
- **Ví dụ:**

```go
reflect.TypeOf(time.Time{}).String() // "time.Time"
```

#### 🔹 `Comparable()`

- **Mô tả:** Trả về true nếu có thể dùng toán tử so sánh như == hoặc !=.
- **Ví dụ:**

```go
reflect.TypeOf(123).Comparable() // true
reflect.TypeOf([]int{}).Comparable() // false
```

#### 🔹 `AssignableTo(type)`

- **Mô tả:** Trả về true nếu kiểu hiện tại có thể gán cho kiểu được truyền vào.
- **Ví dụ:**

```go
t1 := reflect.TypeOf("hello")
t2 := reflect.TypeOf(interface{}(nil))
t1.AssignableTo(t2) // true
```

---

### 📗 Các giá trị `reflect.Kind`

| Tên          | Mô tả                         |
| ------------ | ----------------------------- |
| `Bool`       | Kiểu bool                     |
| `Int` ...    | Các kiểu int: int8, int16,... |
| `Uint` ...   | Các kiểu uint: uint8,...      |
| `Float32/64` | Kiểu số thực                  |
| `String`     | Chuỗi                         |
| `Struct`     | Cấu trúc                      |
| `Array`      | Mảng cố định                  |
| `Slice`      | Mảng động                     |
| `Map`        | Bản đồ                        |
| `Chan`       | Channel                       |
| `Func`       | Hàm                           |
| `Interface`  | Interface                     |
| `Ptr`        | Con trỏ                       |

---

### 🔍 Ví dụ tổng hợp

```go
package main

import (
	"fmt"
	"reflect"
	"time"
)

type Product struct {
	Name string
	Price float64
}

func main() {
	var p Product
	t := reflect.TypeOf(p)

	fmt.Println("Name:", t.Name())
	fmt.Println("PkgPath:", t.PkgPath())
	fmt.Println("Kind:", t.Kind())
	fmt.Println("String:", t.String())
	fmt.Println("Comparable:", t.Comparable())
}
```

---

### 💡 Lưu ý thêm

- Sử dụng `reflect.Type` khi cần phân tích cấu trúc dữ liệu lúc runtime.
- Thường dùng khi làm việc với `interface{}` hoặc xây dựng thư viện trung gian như ORM, tự động ánh xạ dữ liệu.
