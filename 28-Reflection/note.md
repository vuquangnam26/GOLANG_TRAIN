# README: Tổng hợp các phương thức thao tác Slice và Array trong reflect (Go)

Tài liệu này ghi lại toàn bộ những gì đã trao đổi liên quan đến các hàm thao tác slice và array trong Go sử dụng gói `reflect`, cùng với ví dụ cụ thể.

---

## 🧩 1. Index(index int)

**Mô tả:** Truy cập phần tử tại vị trí `index`.

**Ví dụ:**

```go
val := reflect.ValueOf([]string{"A", "B", "C"})
fmt.Println(val.Index(1)) // Output: B
```

---

## 🧩 2. Len()

**Mô tả:** Trả về **độ dài** của slice/array.

**Ví dụ:**

```go
val := reflect.ValueOf([]int{10, 20, 30})
fmt.Println(val.Len()) // Output: 3
```

---

## 🧩 3. Cap()

**Mô tả:** Trả về **capacity** của slice.

**Ví dụ:**

```go
val := reflect.ValueOf(make([]int, 2, 5))
fmt.Println(val.Cap()) // Output: 5
```

---

## 🧩 4. SetLen(n int)

**Mô tả:** Thay đổi độ dài của slice.

**Lưu ý:** Chỉ áp dụng nếu `CanSet()` là `true`, và `n <= cap(slice)`.

**Ví dụ:**

```go
slice := make([]int, 5, 10)
val := reflect.ValueOf(&slice).Elem()
val.SetLen(3)
fmt.Println(slice) // Output: [0 0 0]
```

---

## 🧩 5. SetCap(n int)

**Mô tả:** Thay đổi capacity của slice. ⚠️ Hiếm dùng.

---

## 🧩 6. Slice(lo, hi int)

**Mô tả:** Cắt slice giống như `a[lo:hi]`.

**Ví dụ:**

```go
val := reflect.ValueOf([]int{1, 2, 3, 4})
newSlice := val.Slice(1, 3)
fmt.Println(newSlice.Interface()) // Output: [2 3]
```

---

## 🧩 7. Slice3(lo, hi, max int)

**Mô tả:** Cắt slice với capacity mới giống `a[lo:hi:max]`.

**Ví dụ:**

```go
val := reflect.ValueOf([]int{1, 2, 3, 4})
newSlice := val.Slice3(1, 3, 4)
fmt.Println(newSlice.Interface()) // Output: [2 3]
fmt.Println(newSlice.Cap())       // Output: 3
```

---

## 📌 Tổng kết nhanh

| Phương thức           | Mục đích                                  |
| --------------------- | ----------------------------------------- |
| `Index(i)`            | Lấy phần tử tại chỉ mục `i`               |
| `Len()`               | Trả về độ dài của slice/array             |
| `Cap()`               | Trả về capacity của slice                 |
| `SetLen(n)`           | Đặt lại độ dài slice (chỉ nếu có thể set) |
| `SetCap(n)`           | Đặt lại capacity (hiếm dùng)              |
| `Slice(lo, hi)`       | Cắt slice kiểu `a[lo:hi]`                 |
| `Slice3(lo, hi, max)` | Cắt slice kiểu `a[lo:hi:max]`             |

---

> Ghi chú: Các thao tác này thuộc gói `reflect` trong Go, chủ yếu dùng trong tình huống cần xử lý dynamic type (kiểu động), ví dụ như viết thư viện chung, ORM, v.v.

# README: Thao tác Map trong Reflect (Go)

Tài liệu này tổng hợp các hàm thao tác map sử dụng gói `reflect` trong Golang, kèm theo ví dụ cụ thể.

---

## 1. MapKeys()

**Mô tả:** Trả về một slice `[]reflect.Value` chứa tất cả key của map.

**Ví dụ:**

```go
m := map[string]int{"a": 1, "b": 2}
val := reflect.ValueOf(m)
for _, key := range val.MapKeys() {
    fmt.Println("Key:", key, "Value:", val.MapIndex(key))
}
```

---

## 2. MapIndex(key reflect.Value)

**Mô tả:** Trả về giá trị tương ứng với `key`. Trả về zero value nếu không tìn thấy, có thể kiểm tra bằng `.IsValid()`.

**Ví dụ:**

```go
m := map[string]int{"a": 1}
val := reflect.ValueOf(m)
key := reflect.ValueOf("a")
fmt.Println(val.MapIndex(key)) // Output: 1

key2 := reflect.ValueOf("x")
fmt.Println(val.MapIndex(key2).IsValid()) // Output: false
```

---

## 3. MapRange()

**Mô tả:** Trả về iterator (\*reflect.MapIter) dùng để duyệt các phần tử map.

**Ví dụ:**

```go
m := map[string]int{"a": 1, "b": 2}
val := reflect.ValueOf(m)
iter := val.MapRange()
for iter.Next() {
    fmt.Printf("Key: %v, Value: %v\n", iter.Key(), iter.Value())
}
```

---

## 4. SetMapIndex(key, val reflect.Value)

**Mô tả:** Gán giá trị mới cho key. Nếu value là zero Value thì xóa key.

**Ví dụ:**

```go
m := map[string]int{"a": 1}
val := reflect.ValueOf(m)

key := reflect.ValueOf("b")
value := reflect.ValueOf(100)
val.SetMapIndex(key, value) // Thêm "b": 100
fmt.Println(m) // map[a:1 b:100]

val.SetMapIndex(key, reflect.Value{}) // Xóa "b"
fmt.Println(m) // map[a:1]
```

---

## 5. Len()

**Mô tả:** Trả về số lượng key-value trong map.

**Ví dụ:**

```go
m := map[string]int{"a": 1, "b": 2}
val := reflect.ValueOf(m)
fmt.Println(val.Len()) // Output: 2
```

---

## 6. Hàm hỗ trợ: `setMap` và `removeFromMap`

### 📌 `setMap(m interface{}, key interface{}, val interface{})`

**Chức năng:** Gán key-value mới vào map.

```go
func setMap(m interface{}, key interface{}, val interface{}) {
    mapValue := reflect.ValueOf(m)
    keyValue := reflect.ValueOf(key)
    valValue := reflect.ValueOf(val)
    if (mapValue.Kind() == reflect.Map &&
        mapValue.Type().Key() == keyValue.Type() &&
        mapValue.Type().Elem() == valValue.Type()) {
        mapValue.SetMapIndex(keyValue, valValue)
    } else {
        Printfln("Not a map or mismatched types")
    }
}
```

**Ví dụ:**

```go
m := map[string]int{"a": 1}
setMap(m, "b", 100)
fmt.Println(m) // Output: map[a:1 b:100]
```

---

### 🗑️ `removeFromMap(m interface{}, key interface{})`

**Chức năng:** Xóa phần tử khỏi map bằng key.

```go
func removeFromMap(m interface{}, key interface{}) {
    mapValue := reflect.ValueOf(m)
    keyValue := reflect.ValueOf(key)
    if (mapValue.Kind() == reflect.Map && mapValue.Type().Key() == keyValue.Type()) {
        mapValue.SetMapIndex(keyValue, reflect.Value{})
    }
}
```

**Ví dụ:**

```go
m := map[string]int{"a": 1, "b": 2}
removeFromMap(m, "b")
fmt.Println(m) // Output: map[a:1]
```

---

## Tóm tắt

| Hàm                | Chức năng                      |
| ------------------ | ------------------------------ |
| `MapKeys()`        | Truy xuất danh sách các key    |
| `MapIndex(key)`    | Lấy value tương ứng key        |
| `MapRange()`       | Duyệt map bằng iterator        |
| `SetMapIndex(k,v)` | Gán value cho key hoặc xóa key |
| `Len()`            | Trả về số lượng key-value      |
| `setMap()`         | Hàm helper thêm key-value      |
| `removeFromMap()`  | Hàm helper xóa key khỏi map    |

---

> Ghi chú: Các hàm này thích hợp khi xây dựng framework, middleware hoặc làm việc với dữ liệu dynamic trong Go.
