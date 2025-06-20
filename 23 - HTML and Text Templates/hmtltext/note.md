# 📜 Template Notes in Go

## 🔹 Load Template từ file

### ✅ `template.ParseFiles(filenames...)`

- Dùng khi bạn có danh sách file cụ thể, không có quy luật đặt tên.

```go
tmpl, err := template.ParseFiles("header.html", "content.html", "footer.html")
```

### ✅ `template.ParseGlob(pattern)`

- Dùng để load theo pattern, ví dụ tất cả file `.html` trong thư mục `templates`:

```go
tmpl, err := template.ParseGlob("templates/*.html")
```

### ✅ So sánh

| Tiêu chí                    | ParseFiles           | ParseGlob              |
| --------------------------- | -------------------- | ---------------------- |
| Cách chỉ định template      | Cứng tên file cụ thể | Theo pattern (đồng bộ) |
| Khi file không theo quy tắc | ✅                   | ❌ Không phù hợp       |
| Gọi nhanh theo wildcard     | ❌                   | ✅                     |

## 🔹 Dùng Template

### ✅ `Execute(w io.Writer, data interface{})`

- Render template chính

```go
tmpl.Execute(os.Stdout, data)
```

### ✅ `ExecuteTemplate(w, "templateName", data)`

- Render template theo đích danh

```go
tmpl.ExecuteTemplate(os.Stdout, "page.html", data)
```

### ✅ `Lookup(name)`

- Truy xuất template theo tên:

```go
t := tmpl.Lookup("footer.html")
t.Execute(os.Stdout, data)
```

### ✅ `Templates()`

- Trả về danh sách template đã parse:

```go
for _, t := range tmpl.Templates() {
    fmt.Println("Template:", t.Name())
}
```

## 📊 Tình huống sử dụng phổ biến

| Tình huống                               | Hàm đề xuất       | Lý do                                  |
| ---------------------------------------- | ----------------- | -------------------------------------- |
| Web app có nhiều trang giao diện         | `ParseGlob`       | Load tất cả template 1 lần, dễ quản lý |
| Chỉ render 1-2 trang đơn giản            | `ParseFiles`      | Gọn, rõ ràng, không cần pattern        |
| Tách nhỏ template (layout, content, ...) | `ExecuteTemplate` | Giúp kiểm soát từng phần hiển thị      |
| Tùy biến dữ liệu đầu vào (struct/map)    | `Execute`         | Có thể truyền bất kỳ kiểu dữ liệu nào  |

## 📄 Ví dụ thực tế

### 📁 Cấu trúc:

```
templates/
├── layout.html
├── index.html
├── footer.html
```

### 📑 `layout.html`

```html
{{ define "layout" }}
<html>
  <body>
    {{ template "content" . }} {{ template "footer" . }}
  </body>
</html>
{{ end }}
```

### 📑 `index.html`

```html
{{ define "content" }}
<h1>{{ .Title }}</h1>
<p>{{ .Message }}</p>
{{ end }}
```

### 📑 `footer.html`

```html
{{ define "footer" }}
<footer>Copyright 2025</footer>
{{ end }}
```

### ✨ Go code:

```go
package main

import (
    "html/template"
    "os"
)

func main() {
    tmpl, err := template.ParseGlob("templates/*.html")
    if err != nil {
        panic(err)
    }
    data := map[string]string{
        "Title":   "Welcome!",
        "Message": "This is a templating demo.",
    }
    tmpl.ExecuteTemplate(os.Stdout, "layout", data)
}
```

---

Bạn có thể dùng các hàm `ParseFiles`, `ParseGlob`, `ExecuteTemplate` thường xuyên khi build web app hoặc gen nội dung động từ template.

## 🧪 Example: Conditional Template Logic in Go

### 🧩 template.html:

```html
<h1>There are {{ len . }} products in the source data.</h1>
<h1>First product: {{ index . 0 }}</h1>
{{ range . }} {{ if lt .Price 100.00 }}
<h1>
  Name: {{ .Name }}, Category: {{ .Category }}, Price: {{- printf "$%.2f" .Price
  }}
</h1>
{{ end }} {{ end }}
```

### 💡 Giải thích:

- `{{ range . }}`: duyệt qua toàn bộ mảng dữ liệu đầu vào.
- `{{ if lt .Price 100.00 }}`: chỉ render phần tử nếu `.Price < 100`.
- `lt` là viết tắt của "less than" (nhỏ hơn).
- `printf` dùng để định dạng số thành chuỗi theo định dạng tiền tệ.

### 🧠 Các hàm điều kiện có thể dùng trong template:

| Function       | Description                             |
| -------------- | --------------------------------------- |
| `eq arg1 arg2` | Trả `true` nếu `arg1 == arg2`           |
| `ne arg1 arg2` | Trả `true` nếu `arg1 != arg2`           |
| `lt arg1 arg2` | Trả `true` nếu `arg1 < arg2`            |
| `le arg1 arg2` | Trả `true` nếu `arg1 <= arg2`           |
| `gt arg1 arg2` | Trả `true` nếu `arg1 > arg2`            |
| `ge arg1 arg2` | Trả `true` nếu `arg1 >= arg2`           |
| `and a b`      | Trả `true` nếu cả hai đều đúng          |
| `not a`        | Trả `true` nếu biểu thức `a` là `false` |

### ✅ Kết quả mẫu:

```html
<h1>There are 6 products in the source data.</h1>
<h1>First product: {Name: Kayak, Category: Watersports, Price: 279}</h1>
<h1>Name: Product2, Category: Cat2, Price: $49.95</h1>
<h1>Name: Product3, Category: Cat3, Price: $79.99</h1>
```

### 📌 Ghi nhớ:

- Có thể lồng `if`, `range`, và các hàm template để xử lý logic phức tạp trong file `.html`.
- Dữ liệu truyền vào `Execute()` nên là slice, struct hoặc map có đủ field để dùng trong biểu thức.

# 📜 Ghi chú: Dùng `json.Unmarshal` trong Go

## 🥪 Ví dụ đơn giản:

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

## 🎯 Ví dụ: Unmarshal vào struct có field tag

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func main() {
    jsonData := `{"name": "Kayak", "price": 279}`

    var p Product
    err := json.Unmarshal([]byte(jsonData), &p)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Struct: %+v\n", p)
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

## 💪 So sánh `Unmarshal` vs `Decoder`:

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

...

## 🧠 Ví dụ bổ sung: Template dùng built-in function

### 🧩 template.html:

```html
<h1>There are {{ len . }} products in the source data.</h1>
<h1>First product: {{ index . 0 }}</h1>
{{ range slice . 3 5 }}
<h1>
  Name: {{ .Name }}, Category: {{ .Category }}, Price, {{- printf "$%.2f" .Price
  }}
</h1>
{{ end }}
```

### 💡 Giải thích:

- `{{ len . }}`: lấy độ dài của dữ liệu được truyền vào `Execute()`.
- `{{ index . 0 }}`: lấy phần tử đầu tiên trong danh sách.
- `{{ range slice . 3 5 }}`: duyệt qua phần tử thứ 4 và 5 trong danh sách.
- `{{ .Name }}, {{ .Category }}, {{ .Price }}`: lấy từng trường cụ thể từ struct.
- `{{ printf "$%.2f" .Price }}`: định dạng số thành kiểu tiền tệ.

### ✅ Kết quả đầu ra mẫu:

```html
<h1>There are 6 products in the source data.</h1>
<h1>First product: {Name: Kayak, Category: Watersports, Price: 279}</h1>
<h1>Name: Product4, Category: Cat4, Price, $129.99</h1>
<h1>Name: Product5, Category: Cat5, Price, $59.95</h1>
```

→ Sử dụng các built-in template function để xử lý logic trực tiếp trong HTML, rất tiện lợi cho việc render dữ liệu phức tạp trong Go.

---

## 🧩 Nested Named Template: `define`, `template`, `if`, `else`

```gotemplate
{{ define "currency" }}{{ printf "$%.2f" . }}{{ end }}

{{ define "basicProduct" -}}
  Name: {{ .Name }}, Category: {{ .Category }}, Price,
  {{- template "currency" .Price }}
{{- end }}

{{ define "expensiveProduct" -}}
  Expensive Product {{ .Name }} ({{ template "currency" .Price }})
{{- end }}

<h1>There are {{ len . }} products in the source data.</h1>
<h1>First product: {{ index . 0 }}</h1>
{{ range . -}}
  {{ if lt .Price 100.00 -}}
    <h1>{{ template "basicProduct" . }}</h1>
  {{ else if gt .Price 1500.00 -}}
    <h1>{{ template "expensiveProduct" . }}</h1>
  {{ else -}}
    <h1>Midrange Product: {{ .Name }} ({{ printf "$%.2f" .Price}})</h1>
  {{ end -}}
{{ end }}
```

### 💡 Ý nghĩa:

- `define`: định nghĩa template có tên.
- `template`: gọi lại template theo tên.
- `if`, `else if`, `else`: kiểm tra điều kiện như giá trị `.Price` để quyết định template sử dụng.
- Rất phù hợp khi muốn chia nhỏ template theo logic và tái sử dụng.

### 📈 Tình huống sử dụng thực tế:

- Render sản phẩm theo giá (rẻ, trung, cao cấp).
- Viết layout HTML phức tạp, chia nhiều khối.
- Giảm lặp lại code trong template phức tạp.

---

## 🧱 Template Blocks (Layout kế thừa)

### 🧩 layout.html

```gotemplate
<html>
  <head><title>{{ block "title" . }}Default Title{{ end }}</title></head>
  <body>
    {{ block "content" . }}Default Content{{ end }}
  </body>
</html>
```

### 🧩 child.html

```gotemplate
{{ define "title" }}Product List{{ end }}
{{ define "content" }}
  <h1>There are {{ len . }} products</h1>
  {{ range . }}<div>{{ .Name }} - {{ printf "$%.2f" .Price }}</div>{{ end }}
{{ end }}
```

### 💡 Ý nghĩa:

- `block` định nghĩa vùng có thể được ghi đè.
- `define` trong file khác sẽ thay thế nội dung `block` tương ứng.
- Phù hợp để tái sử dụng layout giống như kế thừa trong HTML template engines.

### ✅ Tình huống sử dụng:

- Dùng `layout.html` cho phần khung.
- Tạo nhiều `child.html` để render các phần nội dung khác nhau: trang chủ, sản phẩm, liên hệ...
