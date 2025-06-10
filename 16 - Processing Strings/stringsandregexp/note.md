# 📘 Xử lý chuỗi & các hàm Builder trong Go

## ✂️ Hàm Trim (Cắt chuỗi)

| Hàm                              | Mô tả                                        | Ví dụ                           | Kết quả   |
| -------------------------------- | -------------------------------------------- | ------------------------------- | --------- |
| `strings.TrimSpace(s)`           | Xoá khoảng trắng đầu và cuối chuỗi           | `TrimSpace(" hello ")`          | `"hello"` |
| `strings.Trim(s, cutset)`        | Xoá tất cả ký tự đầu/cuối nằm trong `cutset` | `Trim("...hi...", ".")`         | `"hi"`    |
| `strings.TrimLeft(s, cutset)`    | Xoá ký tự bên trái theo `cutset`             | `TrimLeft("...hi", ".")`        | `"hi"`    |
| `strings.TrimRight(s, cutset)`   | Xoá ký tự bên phải theo `cutset`             | `TrimRight("hi...", ".")`       | `"hi"`    |
| `strings.TrimPrefix(s, prefix)`  | Xoá tiền tố nếu có                           | `TrimPrefix("goLang", "go")`    | `"Lang"`  |
| `strings.TrimSuffix(s, suffix)`  | Xoá hậu tố nếu có                            | `TrimSuffix("hello.go", ".go")` | `"hello"` |
| `strings.TrimFunc(s, func)`      | Cắt tuỳ chỉnh bằng hàm                       | Ví dụ xoá khoảng trắng unicode  | tuỳ logic |
| `strings.TrimLeftFunc(s, func)`  | Cắt bên trái với hàm                         | -                               | -         |
| `strings.TrimRightFunc(s, func)` | Cắt bên phải với hàm                         | -                               | -         |

---

## 🔎 Regex (Biểu thức chính quy - gói regexp)

### Match và Find cơ bản

| Hàm                        | Mô tả                                      | Ví dụ                   | Kết quả            |
| -------------------------- | ------------------------------------------ | ----------------------- | ------------------ |
| `MatchString(s)`           | Trả về `true` nếu pattern khớp chuỗi       | `MatchString("abc123")` | `true/false`       |
| `FindStringIndex(s)`       | Vị trí `[start, end]` của kết quả đầu tiên | -                       | `[0, 3]`           |
| `FindAllStringIndex(s, n)` | Tất cả các vị trí khớp                     | -                       | `[[0, 3], [5, 8]]` |
| `FindString(s)`            | Kết quả chuỗi đầu tiên khớp                | -                       | `"abc"`            |
| `FindAllString(s, n)`      | Tất cả chuỗi khớp                          | -                       | `["abc", "def"]`   |
| `Split(s, n)`              | Cắt chuỗi bằng biểu thức regex             | -                       | `["a", "b"]`       |

### Submatch (Nhóm kết quả)

| Hàm                                | Mô tả                                       |
| ---------------------------------- | ------------------------------------------- |
| `FindStringSubmatch(s)`            | Trả về mảng gồm kết quả khớp + các nhóm phụ |
| `FindAllStringSubmatch(s, n)`      | Tất cả khớp và nhóm phụ                     |
| `FindStringSubmatchIndex(s)`       | Giống trên nhưng trả về chỉ số              |
| `FindAllStringSubmatchIndex(s, n)` | Tất cả khớp + chỉ số                        |
| `NumSubexp()`                      | Số lượng nhóm phụ                           |
| `SubexpNames()`                    | Tên các nhóm phụ theo thứ tự định nghĩa     |
| `SubexpIndex(name)`                | Vị trí nhóm phụ theo tên                    |

### Ví dụ biểu thức có tên nhóm:

```go
pattern := `(?P<day>\d{2})-(?P<month>\d{2})-(?P<year>\d{4})`
re := regexp.MustCompile(pattern)
matches := re.FindStringSubmatch("09-06-2025")
names := re.SubexpNames()
```

Truy cập kết quả như:

```go
map["day"] => "09", ["month"] => "06", ["year"] => "2025"
```

---

## 🧪 Hàm Scan / Sscanf / Fscan

| Hàm                                 | Mô tả                                                       |
| ----------------------------------- | ----------------------------------------------------------- |
| `Scan(...vals)`                     | Đọc đầu vào chuẩn, tách theo khoảng trắng, gán vào các biến |
| `Scanln(...vals)`                   | Giống `Scan` nhưng dừng khi gặp newline (`\n`)              |
| `Scanf(template, ...vals)`          | Đọc đầu vào theo mẫu định dạng (`template`)                 |
| `Fscan(reader, ...vals)`            | Đọc dữ liệu từ `reader`, tách khoảng trắng                  |
| `Fscanln(reader, ...vals)`          | Giống `Fscan` nhưng dừng tại newline                        |
| `Fscanf(reader, template, ...vals)` | Đọc từ `reader` theo mẫu định dạng                          |
| `Sscan(str, ...vals)`               | Scan từ chuỗi `str`, tách khoảng trắng                      |
| `Sscanf(str, template, ...vals)`    | Scan từ chuỗi `str` theo mẫu định dạng                      |
| `Sscanln(str, ...vals)`             | Scan chuỗi và dừng tại newline                              |

### Ví dụ:

```go
var name string
var age int
fmt.Sscanf("Alice 25", "%s %d", &name, &age)
fmt.Println(name, age) // Alice 25
```

---

## 🆚 So sánh Trim - Replace - Fields

| Trường hợp sử dụng   | Hàm phù hợp                               |
| -------------------- | ----------------------------------------- |
| Xoá khoảng trắng     | `TrimSpace()`                             |
| Xoá ký tự cụ thể     | `Trim(), TrimLeft(), TrimRight()`         |
| Xoá tiền tố / hậu tố | `TrimPrefix(), TrimSuffix()`              |
| Làm sạch bằng regex  | `regexp.MustCompile().ReplaceAllString()` |
| Tách từ trong chuỗi  | `strings.Fields()`                        |

---

## 📚 Ghi chú

* `Trim` xoá từng ký tự, không phải chuỗi con.
* `TrimSpace` đặc biệt chỉ xoá các ký tự trắng unicode.
* `Scan` phù hợp đọc dòng đầu vào dạng nhập liệu.
* Nên dùng `MustCompile` nếu chắc chắn regex đúng.

---

> 📄 File này dùng để ghi chú nhanh về cách xử lý chuỗi, regex, và đọc dữ liệu đầu vào trong Go.
