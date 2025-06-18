# 🧾 Ghi chú chi tiết dòng code: Decode JSON với map trong Go

```go
reader := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
```

- Tạo một `io.Reader` từ chuỗi JSON. Dữ liệu JSON được mô phỏng như từ file hoặc request body.

```go
m := map[string]interface{}{}
```

- Khởi tạo một map để chứa dữ liệu JSON sau khi decode.
- Dùng `interface{}` để có thể chứa mọi kiểu dữ liệu.

```go
decoder := json.NewDecoder(reader)
```

- Tạo một JSON decoder từ nguồn `reader`.
- Có thể dùng để đọc từng phần tử từ luồng nếu JSON phức tạp.

```go
err := decoder.Decode(&m)
```

- Decode nội dung JSON vào địa chỉ của `m`.
- Nếu JSON không hợp lệ → trả về lỗi trong `err`.

```go
if err != nil {
    Printfln("Error: %v", err.Error())
} else {
    Printfln("Map: %T, %v", m, m)
```

- Kiểm tra lỗi giải mã. Nếu có lỗi, in ra thông báo.
- Nếu không, in kiểu và nội dung của map.

```go
    for k, v := range m {
        Printfln("Key: %v, Value: %v", k, v)
    }
```

- Duyệt từng phần tử trong map `m`.
- In ra từng cặp `key - value`. Giá trị `v` có thể là `float64` do cách `encoding/json` xử lý số.

```go

```

# 🧾 Ghi chú: Dùng `json.Unmarshal` trong Go

## 🧪 Ví dụ đơn giản:

```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    data := []byte(`{"Kayak": 279, "Lifejacket": 49.95}`)

    var m map[string]interface{}
    err := json.Unmarshal(data, &m)

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Map: %T, %v\n", m, m)
        for k, v := range m {
            fmt.Printf("Key: %v, Value: %v\n", k, v)
        }
    }
}
```

## 📌 Ghi chú dòng lệnh:

### `data := []byte(...)`

- Biến `data` chứa chuỗi JSON ở dạng mảng byte.
- `Unmarshal` yêu cầu kiểu `[]byte`, không dùng reader như `Decoder`.

### `var m map[string]interface{}`

- Map có khóa là chuỗi, giá trị là `interface{}` để chứa mọi kiểu dữ liệu (string, int, float...).

### `json.Unmarshal(data, &m)`

- Giải mã JSON thành map `m`.
- Cần truyền địa chỉ (`&m`) để `Unmarshal` có thể ghi dữ liệu.

### `if err != nil {...}`

- Kiểm tra lỗi giải mã JSON. Nếu JSON sai định dạng → trả lỗi.

### `fmt.Printf(...)`

- In thông tin kiểu và giá trị.
- Duyệt map để in từng key-value.

## 🔄 So sánh `Unmarshal` vs `Decoder`:

| Tiêu chí              | `Unmarshal` | `Decoder`                             |
| --------------------- | ----------- | ------------------------------------- |
| Nguồn dữ liệu         | `[]byte`    | `io.Reader` (file, stream, socket...) |
| Đọc JSON từng phần    | Không       | Có (`Token`, `Decode` nhiều lần)      |
| Dễ dùng cho chuỗi nhỏ | ✅          | ❌ (quá nặng với chuỗi nhỏ)           |
| Phù hợp cho file lớn  | ❌          | ✅                                    |

## 📘 Ghi nhớ:

- Dùng `Unmarshal` nếu bạn đã có JSON dạng chuỗi hoặc mảng byte.
- Dùng `Decoder` nếu đọc từ file hoặc stream.
- Khi map chứa số → `Unmarshal` sẽ gán kiểu `float64` mặc định.
