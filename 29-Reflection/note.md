# Reflect: Thao Tác & Phân Tích Hàm (Function) trong Go

Tài liệu này ghi chép lại các phương thức trong package `reflect` dùng để phân tích hàm trong Golang.

---

## 1. NumIn()

**Chức năng:** Trả về số lượng tham số (đối số) mà hàm nhận.

```go
fn := func(a int, b string) {}
t := reflect.TypeOf(fn)
fmt.Println(t.NumIn()) // Output: 2
```

---

## 2. In(index int)

**Chức năng:** Truy xuất Type của tham số tại index.

```go
fmt.Println(t.In(0)) // Output: int
fmt.Println(t.In(1)) // Output: string
```

---

## 3. IsVariadic()

**Chức năng:** Trả về true nếu hàm nhận số lượng tham số biến thiết (...).

```go
fn := func(nums ...int) {}
t := reflect.TypeOf(fn)
fmt.Println(t.IsVariadic()) // Output: true
```

---

## 4. NumOut()

**Chức năng:** Trả về số lượng giá trị trả về (output).

```go
fn := func() (int, string) { return 1, "ok" }
t := reflect.TypeOf(fn)
fmt.Println(t.NumOut()) // Output: 2
```

---

## 5. Out(index int)

**Chức năng:** Truy xuất Type của giá trị trả về tại index.

```go
fmt.Println(t.Out(0)) // Output: int
fmt.Println(t.Out(1)) // Output: string
```

---

## Tóm tắt

| Phương Thức    | Mô tả                           |
| -------------- | ------------------------------- |
| `NumIn()`      | Số tham số được hàm nhận        |
| `In(i)`        | Type của tham số tại i          |
| `IsVariadic()` | Hàm có phải nhận số biến thiết? |
| `NumOut()`     | Số giá trị hàm trả về           |
| `Out(i)`       | Type của giá trị trả về tại i   |

---

> File này phù hợp cho những ai muốn phân tích các hàm một cách dynamic hoặc xây dựng framework trong Go.

# README: Thao tác Map và Method trong Reflect (Go)

Tài liệu này tổng hợp các hàm thao tác `map` và `method` sử dụng gói `reflect` trong Golang, kèm theo ví dụ cụ thể.

---

## 📌 PHẦN 1: THAO TÁC MAP

### 1. MapKeys()

**Mô tả:** Trả về một slice `[]reflect.Value` chứa tất cả key của map.

```go
m := map[string]int{"a": 1, "b": 2}
val := reflect.ValueOf(m)
for _, key := range val.MapKeys() {
    fmt.Println("Key:", key, "Value:", val.MapIndex(key))
}
```

### 2. MapIndex(key reflect.Value)

**Mô tả:** Trả về giá trị tương ứng với `key`. Trả về zero value nếu không tìm thấy, có thể kiểm tra bằng `.IsValid()`.

### 3. MapRange()

**Mô tả:** Trả về iterator (`*reflect.MapIter`) dùng để duyệt các phần tử map.

```go
m := map[string]int{"a": 1, "b": 2}
val := reflect.ValueOf(m)
iter := val.MapRange()
for iter.Next() {
    fmt.Printf("Key: %v, Value: %v\n", iter.Key(), iter.Value())
}
```

### 4. SetMapIndex(key, val reflect.Value)

**Mô tả:** Gán giá trị mới cho key. Nếu value là zero Value thì xóa key.

### 5. Len()

**Mô tả:** Trả về số lượng key-value trong map.

### ➕ Các hàm thao tác Map bằng Reflection:

```go
func setMap(m interface{}, key interface{}, val interface{}) {
    mapValue := reflect.ValueOf(m)
    keyValue := reflect.ValueOf(key)
    valValue := reflect.ValueOf(val)
    if (mapValue.Kind() == reflect.Map &&
        mapValue.Type().Key() == keyValue.Type() &&
        mapValue.Type().Elem() == valValue.Type()) {
        mapValue.SetMapIndex(keyValue, valValue)
    }
}

func removeFromMap(m interface{}, key interface{}) {
    mapValue := reflect.ValueOf(m)
    keyValue := reflect.ValueOf(key)
    if (mapValue.Kind() == reflect.Map && mapValue.Type().Key() == keyValue.Type()) {
        mapValue.SetMapIndex(keyValue, reflect.Value{})
    }
}
```

---

## 📌 PHẦN 2: THAO TÁC METHOD TRÊN STRUCT

### 1. NumMethod()

**Mô tả:** Trả về số lượng method export được của kiểu struct.

### 2. Method(index)

**Mô tả:** Trả về method tại vị trí index (kiểu `reflect.Method`).

### 3. MethodByName(name)

**Mô tả:** Trả về method theo tên.

```go
t := reflect.TypeOf(obj)
m, ok := t.MethodByName("Hello")
if ok {
    m.Func.Call([]reflect.Value{reflect.ValueOf(obj)})
}
```

### ✅ Struct Method có cấu trúc:

```go
type Method struct {
    Name    string
    PkgPath string // Nếu không phải exported thì có package path
    Type    reflect.Type
    Func    reflect.Value // method như function
    Index   int           // vị trí
}
```

---

## 📌 Ví dụ sử dụng `makeMapperFunc`

Tạo function wrapper dùng reflection:

```go
func makeMapperFunc(mapper interface{}) interface{} {
    mapVal := reflect.ValueOf(mapper)
    if mapVal.Kind() == reflect.Func && mapVal.Type().NumIn() == 1 && mapVal.Type().NumOut() == 1 {
        inType := reflect.SliceOf(mapVal.Type().In(0))
        outType := reflect.SliceOf(mapVal.Type().Out(0))

        funcType := reflect.FuncOf([]reflect.Type{inType}, []reflect.Type{outType}, false)

        funcVal := reflect.MakeFunc(funcType, func(params []reflect.Value) []reflect.Value {
            srcSliceVal := params[0]
            resultsSliceVal := reflect.MakeSlice(outType, srcSliceVal.Len(), srcSliceVal.Len())

            for i := 0; i < srcSliceVal.Len(); i++ {
                r := mapVal.Call([]reflect.Value{srcSliceVal.Index(i)})
                resultsSliceVal.Index(i).Set(r[0])
            }

            return []reflect.Value{resultsSliceVal}
        })

        return funcVal.Interface()
    }
    return nil
}
```

---

> 📘 **Lưu ý:** Các thao tác `reflect` rất mạnh mẽ nhưng dễ lỗi runtime và khó đọc, nên chỉ dùng khi xử lý dynamic, generic, hoặc viết framework, middleware.
