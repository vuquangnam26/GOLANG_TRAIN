# README: Phản chiếu Struct trong Go (reflect.StructField)

Tài liệu này tổng hợp các hàm và thuộc tính liên quan đến việc phản chiếu struct trong Golang bằng gói `reflect`, đặc biệt với kiểu `reflect.StructField`.

---

## 1. Các hàm truy xuất field của struct

### `NumField()`

- **Chức năng:** Trả về số lượng field trong struct.

```go
reflect.TypeOf(Person{}).NumField()
```

### `Field(index int)`

- **Chức năng:** Trả về field tại vị trí index.

```go
field := t.Field(0)
fmt.Println(field.Name)
```

### `FieldByIndex(indices []int)`

- **Chức năng:** Truy xuất field lồng nhau (nested).

```go
t.FieldByIndex([]int{1, 0}) // ví dụ với struct Address.City
```

### `FieldByName(name string)`

- **Chức năng:** Truy xuất field theo tên.

```go
field, found := t.FieldByName("Age")
```

### `FieldByNameFunc(func)`

- **Chức năng:** Truy xuất theo điều kiện hàm.

```go
field, _ := t.FieldByNameFunc(func(s string) bool {
    return strings.ToLower(s) == "name"
})
```

---

## 2. Cấu trúc `reflect.StructField`

| Trường      | Ý nghĩa                                                         |
| ----------- | --------------------------------------------------------------- |
| `Name`      | Tên field trong struct                                          |
| `PkgPath`   | Tên package. Rỗng nếu field là exported (public)                |
| `Type`      | Kiểu dữ liệu của field (kiểu `reflect.Type`)                    |
| `Tag`       | Tag của field (struct tag như `json:"field"`)                   |
| `Index`     | Chỉ số field (cho phép truy xuất field lồng nhau)               |
| `Anonymous` | `true` nếu field là embedded (ẩn danh - struct nhúng trực tiếp) |

---

## 3. Ví dụ minh họa

```go
type Address struct {
    City string
}
type Person struct {
    Name string `json:"name"`
    Age  int
    Address
}

func main() {
    t := reflect.TypeOf(Person{})
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Name: %s, Type: %s, Tag: %s, Anonymous: %v\n",
            field.Name, field.Type, field.Tag, field.Anonymous)
    }
}
```

---

## 4. Tổng kết

| Hàm / Trường            | Ý nghĩa                          |
| ----------------------- | -------------------------------- |
| `NumField()`            | Trả về số lượng field            |
| `Field(i)`              | Truy xuất field tại vị trí i     |
| `FieldByName(name)`     | Truy xuất field theo tên         |
| `FieldByIndex([]int)`   | Truy xuất field lồng nhau        |
| `FieldByNameFunc()`     | Truy xuất theo điều kiện         |
| `StructField.Tag`       | Lấy tag của field                |
| `StructField.Type`      | Kiểu dữ liệu field               |
| `StructField.Anonymous` | Kiểm tra field có embedded không |

---

> Những API này rất hữu dụng trong các tình huống cần xử lý dữ liệu dynamic hoặc tạo các thư viện/ORM/framework tự động ánh xạ dữ liệu.
