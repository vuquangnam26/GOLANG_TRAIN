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

---

## 📂 Làm việc với File: `WriteFile` vs `OpenFile`

### 🔹 `os.WriteFile(name, data []byte, perm fs.FileMode)`

- **Tạo hoặc ghi đè file** với nội dung từ `[]byte`.
- Nhanh, đơn giản, dùng cho file nhỏ.

✅ Ví dụ:

```go
err := os.WriteFile("example.txt", []byte("Hello!"), 0644)
if err != nil {
    log.Fatal(err)
}
```

### 🔹 `os.OpenFile(name, flag, perm)`

- Mở file nâng cao với các cờ (`READ`, `WRITE`, `CREATE`,...)

✅ Ví dụ:

```go
file, err := os.OpenFile("example.txt", os.O_CREATE|os.O_RDWR, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()
file.WriteString("More data")
```

### 📌 So sánh nhanh:

| Tiêu chí          | `os.WriteFile`                      | `os.OpenFile`                                  |
| ----------------- | ----------------------------------- | ---------------------------------------------- |
| Mục đích chính    | Ghi nhanh toàn bộ nội dung vào file | Mở file với nhiều tùy chọn (ghi, đọc, tạo,...) |
| Ghi nối (append)? | ❌ Không hỗ trợ                     | ✅ Có thể dùng với `os.O_APPEND`               |
| Ghi đè dữ liệu?   | ✅ Ghi đè hoàn toàn                 | ✅ Tuỳ chọn qua flag `O_TRUNC` hoặc `O_APPEND` |
| Thích hợp cho     | File nhỏ, ghi 1 lần                 | Ghi từng phần, xử lý nâng cao                  |

---

## 🔖 Các cờ mở file (`File Opening Flags`)

| Tên cờ (Flag) | Mô tả (Description)                                                                                             |
| ------------- | --------------------------------------------------------------------------------------------------------------- |
| `O_RDONLY`    | Mở file chỉ để đọc – có thể đọc từ file nhưng **không ghi được**.                                               |
| `O_WRONLY`    | Mở file chỉ để ghi – có thể ghi vào file nhưng **không đọc được**.                                              |
| `O_RDWR`      | Mở file để vừa đọc vừa ghi.                                                                                     |
| `O_APPEND`    | Ghi nội dung mới vào **cuối file** (thêm nội dung).                                                             |
| `O_CREATE`    | Tạo file nếu **file chưa tồn tại**.                                                                             |
| `O_EXCL`      | Dùng cùng với `O_CREATE` để đảm bảo chỉ tạo file **mới**. Nếu file đã tồn tại → báo lỗi.                        |
| `O_SYNC`      | Ghi dữ liệu một cách **đồng bộ**, đảm bảo dữ liệu được ghi xuống thiết bị lưu trữ **trước khi hàm ghi trả về**. |
| `O_TRUNC`     | **Xóa toàn bộ nội dung hiện tại** trong file khi mở.                                                            |

---

## 🔹 Thao tác viết vào file (các phương thức `File`)

| Tên hàm                  | Mô tả                                                                           |
| ------------------------ | ------------------------------------------------------------------------------- |
| `Seek(offset, how)`      | Thiết lập vị trí cho thao tác đọc/ghi tiếp theo.                                |
| `Write(slice)`           | Ghi một mảng byte vào file. Trả về số byte đã ghi và lỗi (nếu có).              |
| `WriteAt(slice, offset)` | Ghi mảng byte vào vị trí cụ thể trong file. Tương ứng với `ReadAt`.             |
| `WriteString(str)`       | Ghi chuỗi vào file. Tiện lợi vì nó chuyển chuỗi sang byte slice và gọi `Write`. |

---

## 🔹 Tạo File với `Create` và `CreateTemp`

| Tên hàm                 | Mô tả                                                                                                                                               |
| ----------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Create(name)`          | Tạo và mở file với cờ `O_RDWR`, `O_CREATE`, `O_TRUNC`. Nếu file tồn tại → nội dung cũ bị xóa. Trả về đối tượng `File` và lỗi nếu có.                |
| `CreateTemp(dir, name)` | Tạo file tạm với tên ngẫu nhiên trong thư mục chỉ định (`dir`). Dùng các cờ `O_RDWR`, `O_CREATE`, `O_EXCL`. File **không bị xóa tự động** khi đóng. |

📌 Ghi nhớ:

- `Create` tiện để khởi tạo file mới hoặc ghi đè nội dung.
- `CreateTemp` dùng để tạo file tạm nhưng không tự xóa → cần dọn thủ công.
- Tên file tạm chứa chuỗi ngẫu nhiên để tránh trùng lặp.

---

## 🔄 So sánh `Create` / `CreateTemp` vs `Write`, `WriteAt`, `WriteString`

| Tiêu chí                       | `Create` / `CreateTemp`                                        | `Write` / `WriteAt` / `WriteString`                              |
| ------------------------------ | -------------------------------------------------------------- | ---------------------------------------------------------------- |
| Chức năng chính                | **Tạo** và mở một file mới (thường dùng trước khi ghi dữ liệu) | **Ghi dữ liệu** vào file đã mở                                   |
| Trả về gì?                     | Một đối tượng `*os.File` dùng để thao tác đọc/ghi              | Số byte đã ghi + lỗi nếu có                                      |
| Có tạo file nếu chưa tồn tại?  | ✅ `Create` và `CreateTemp` đều tạo file mới                   | ❌ Phải có file trước (mở bằng `OpenFile` hoặc `Create`)         |
| Có xóa nội dung file cũ không? | ✅ `Create` có `O_TRUNC`: xóa nội dung cũ nếu file tồn tại     | ❌ Không – trừ khi kết hợp với flag trong `OpenFile` (`O_TRUNC`) |
| Có ghi dữ liệu không?          | ❌ Không ghi gì – chỉ tạo file                                 | ✅ Ghi nội dung từ byte slice hoặc string vào file               |
| Thường dùng khi nào?           | Khi cần **khởi tạo file mới** hoặc **tạo file tạm**            | Khi đã có file mở và cần ghi dữ liệu cụ thể                      |
| Ví dụ                          | `file := os.Create("out.txt")`                                 | `file.Write([]byte("Hello"))`, `file.WriteString("Hello")`       |

---

## 📚 Các hàm định vị thư mục trong gói `os`

| Tên hàm           | Mô tả                                                          |
| ----------------- | -------------------------------------------------------------- |
| `Getwd()`         | Trả về thư mục hiện tại (working directory) và lỗi nếu có.     |
| `UserHomeDir()`   | Trả về thư mục home của người dùng và lỗi nếu có.              |
| `UserCacheDir()`  | Trả về thư mục cache mặc định của người dùng và lỗi nếu có.    |
| `UserConfigDir()` | Trả về thư mục cấu hình mặc định của người dùng và lỗi nếu có. |
| `TempDir()`       | Trả về thư mục mặc định dùng cho file tạm và lỗi nếu có.       |

---

## 🧱 Các hàm thao tác đường dẫn trong `path/filepath`

| Hàm                      | Mô tả                                                             |
| ------------------------ | ----------------------------------------------------------------- |
| `Abs(path)`              | Trả về đường dẫn tuyệt đối (dựa trên working directory hiện tại). |
| `IsAbs(path)`            | Trả về `true` nếu đường dẫn là tuyệt đối.                         |
| `Base(path)`             | Trả về phần cuối cùng của đường dẫn (tên file hoặc thư mục).      |
| `Clean(path)`            | Dọn dẹp đường dẫn: bỏ ký tự `..`, `.` dư thừa, hoặc dấu `/` dư.   |
| `Dir(path)`              | Trả về phần thư mục (không bao gồm tên file cuối).                |
| `EvalSymlinks(path)`     | Trả về đường dẫn thực tế nếu có symbolic link.                    |
| `Ext(path)`              | Trả về phần đuôi mở rộng (extension) của file (vd: `.txt`).       |
| `FromSlash(path)`        | Chuyển `/` thành separator hệ điều hành (`\` trên Windows).       |
| `ToSlash(path)`          | Chuyển separator hệ điều hành thành `/`.                          |
| `Join(elem1, elem2,...)` | Nối các thành phần thành 1 đường dẫn hợp lệ với hệ điều hành.     |
| `Match(pattern, path)`   | Trả về true nếu `path` khớp với `pattern` (vd: `*.txt`).          |
| `Split(path)`            | Tách path thành 2 phần: thư mục và tên file.                      |
| `SplitList(path)`        | Tách đường dẫn nhiều thành phần (PATH env var) thành slice.       |
| `VolumeName(path)`       | Trả về tên volume nếu có (vd: `C:` trên Windows).                 |

---

## 👀 Một số thao tác file/directory khác trong `os`

| Hàm                     | Mô tả                                                                                                         |
| ----------------------- | ------------------------------------------------------------------------------------------------------------- |
| `MkdirTemp(dir, name)`  | Tạo thư mục tạm với tên ngẫu nhiên trong `dir`. Nếu `name` chứa `*`, phần đó được thay bằng chuỗi ngẫu nhiên. |
| `Remove(name)`          | Xoá file hoặc thư mục (nếu rỗng) được chỉ định.                                                               |
| `RemoveAll(name)`       | Xoá file hoặc thư mục cùng với toàn bộ nội dung con của nó.                                                   |
| `Rename(old, new)`      | Đổi tên file hoặc thư mục từ `old` sang `new`.                                                                |
| `Symlink(old, new)`     | Tạo symbolic link `new` trỏ đến `old`.                                                                        |
| `Chdir(dir)`            | Đổi thư mục làm việc hiện tại sang `dir`.                                                                     |
| `Mkdir(name, perms)`    | Tạo một thư mục với tên và quyền truy cập chỉ định.                                                           |
| `MkdirAll(name, perms)` | Tạo thư mục cùng với các thư mục cha nếu chưa tồn tại.                                                        |

## 📁 Đọc thư mục với `ReadDir`

### 🔹 `ReadDir(name)`

- Đọc nội dung của thư mục `name` và trả về một slice chứa các đối tượng `DirEntry`.
- Mỗi `DirEntry` đại diện cho một file hoặc thư mục con.

### ✅ Ví dụ:

```go
entries, err := os.ReadDir("mydir")
if err != nil {
    log.Fatal(err)
}

for _, entry := range entries {
    fmt.Println("Name:", entry.Name())
    fmt.Println("IsDir:", entry.IsDir())
    fmt.Println("Type:", entry.Type())
    info, _ := entry.Info()
    fmt.Println("Size:", info.Size())
    fmt.Println("ModTime:", info.ModTime())
}
```

---

## 🔍 Giao diện `DirEntry`

| Phương thức | Mô tả                                                               |
| ----------- | ------------------------------------------------------------------- |
| `Name()`    | Trả về tên của file/thư mục.                                        |
| `IsDir()`   | Trả về `true` nếu là thư mục.                                       |
| `Type()`    | Trả về `fs.FileMode` mô tả loại và quyền truy cập.                  |
| `Info()`    | Trả về `FileInfo` chứa thêm thông tin như kích thước, thời gian,... |

## 📘 Giao diện `FileInfo`

| Phương thức | Mô tả                                               |
| ----------- | --------------------------------------------------- |
| `Name()`    | Tên của file hoặc thư mục.                          |
| `Size()`    | Kích thước file (int64).                            |
| `Mode()`    | Quyền truy cập và kiểu file (`fs.FileMode`).        |
| `ModTime()` | Thời điểm file được sửa đổi lần cuối (`time.Time`). |

---

## 📄 Kiểm tra file với `Stat`

### 🔹 `Stat(path)`

- Trả về đối tượng `FileInfo` chứa thông tin về file hoặc thư mục tại `path`.

✅ Ví dụ:

```go
info, err := os.Stat("myfile.txt")
if err != nil {
    log.Fatal(err)
}

fmt.Println("Name:", info.Name())
fmt.Println("Size:", info.Size())
fmt.Println("Mode:", info.Mode())
fmt.Println("Last modified:", info.ModTime())
```

📌 Ghi nhớ:

- `Stat` hữu ích để kiểm tra file có tồn tại không, và lấy thông tin như thời gian cập nhật, quyền truy cập, kích thước.
  🔍 Khớp và tìm đường dẫn: Match vs Glob
  Hàm Mô tả
  Match(pattern, name) So khớp một chuỗi tên với mẫu (pattern). Trả về bool cho biết có khớp hay không và error nếu mẫu không hợp lệ.
  Glob(pathPattern) Tìm tất cả file/thư mục khớp với mẫu (_.txt, data/_.json,...). Trả về slice các đường dẫn khớp và error nếu có lỗi trong quá trình.
