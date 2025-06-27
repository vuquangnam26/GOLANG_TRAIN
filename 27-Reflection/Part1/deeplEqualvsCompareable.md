## 🔍 So sánh `reflect.DeepEqual` vs `==` kèm `reflect.Type.Comparable()` trong Go

### 🧪 Mục đích

So sánh sự khác nhau giữa hai phương pháp kiểm tra hai giá trị có bằng nhau trong Go:

- Dùng toán tử `==` (có kiểm tra `Type.Comparable()`)
- Dùng `reflect.DeepEqual`

---

### ✅ Sử dụng `==` và `reflect.Type.Comparable()`

```go
if reflect.TypeOf(a).Comparable() && reflect.TypeOf(b).Comparable() {
    if a == b {
        fmt.Println("Equal")
    }
}
```

- An toàn khi biết trước kiểu dữ liệu.
- Nhanh hơn `reflect.DeepEqual`.
- KHÔNG dùng được với slice, map, func (gây panic).

---

### ✅ Sử dụng `reflect.DeepEqual`

```go
if reflect.DeepEqual(a, b) {
    fmt.Println("Deep Equal")
}
```

- Dùng được với hầu hết các kiểu dữ liệu.
- Chấp nhận so sánh những giá trị phức tạp: slice, array, struct...
- Chậm hơn do duyệt đệ quy.

---

### 💡 Đề xuất khi nào dùng cái nào?

| Tình huống                   | Nên dùng              |
| ---------------------------- | --------------------- |
| So sánh primitive type       | `==` + `Comparable()` |
| So sánh slice, array         | `reflect.DeepEqual`   |
| So sánh map, struct phức tạp | `reflect.DeepEqual`   |
| Tối ưu hiệu năng cao         | `==` với kiểm tra     |

---

### 📌 Lưu ý khi sử dụng

- `reflect.DeepEqual` coi `nil` và zero value khác nhau:

```go
var a []int = nil
b := []int{}
fmt.Println(reflect.DeepEqual(a, b)) // false
```

- `==` chỉ hoạt động khi kiểu là `comparable`.

---

### 🧪 Ví dụ so sánh an toàn trong slice:

```go
func contains(slice interface{}, target interface{}) bool {
    sliceVal := reflect.ValueOf(slice)
    targetType := reflect.TypeOf(target)

    if sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem().Comparable() && targetType.Comparable() {
        for i := 0; i < sliceVal.Len(); i++ {
            if sliceVal.Index(i).Interface() == target {
                return true
            }
        }
    } else {
        for i := 0; i < sliceVal.Len(); i++ {
            if reflect.DeepEqual(sliceVal.Index(i).Interface(), target) {
                return true
            }
        }
    }
    return false
}
```

---

### 🔗 Tài liệu tham khảo

- [reflect.DeepEqual](https://pkg.go.dev/reflect#DeepEqual)
- [reflect.Type.Comparable](https://pkg.go.dev/reflect#Type.Comparable)
