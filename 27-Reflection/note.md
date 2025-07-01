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

### 🏗️ Hàm tạo kiểu với `reflect`

| Hàm                                   | Giải thích                                                  |
| ------------------------------------- | ----------------------------------------------------------- |
| `reflect.New(type)`                   | Tạo giá trị mới dạng con trỏ trỏ đến giá trị zero của type. |
| `reflect.Zero(type)`                  | Trả về giá trị zero cho type đó.                            |
| `reflect.MakeMap(type)`               | Tạo bản đồ mới với kiểu đã cho.                             |
| `reflect.MakeMapWithSize(type, size)` | Tạo bản đồ với số lượng phần tử khởi tạo.                   |
| `reflect.MakeSlice(type, capacity)`   | Tạo slice với type và capacity chỉ định.                    |
| `reflect.MakeFunc(type, func)`        | Tạo hàm động với các kiểu đầu vào và đầu ra.                |
| `reflect.MakeChan(type, buffer)`      | Tạo channel mới với kích thước buffer.                      |

#### 📌 Ví dụ:

```go
// reflect.New
varType := reflect.TypeOf(123)
ptrVal := reflect.New(varType)
fmt.Println(ptrVal.Elem()) // in ra 0 vì là giá trị zero

// reflect.Zero
zero := reflect.Zero(reflect.TypeOf("Hello"))
fmt.Println(zero.String()) // in ra chuỗi rỗng ""

// reflect.MakeSlice
sliceType := reflect.SliceOf(reflect.TypeOf(""))
slice := reflect.MakeSlice(sliceType, 0, 5)
fmt.Println(slice.Len()) // 0
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
- [reflect.New](https://pkg.go.dev/reflect#New)
- [reflect.MakeSlice](https://pkg.go.dev/reflect#MakeSlice)
- [reflect.MakeMap](https://pkg.go.dev/reflect#MakeMap)
- [reflect.MakeChan](https://pkg.go.dev/reflect#MakeChan)
