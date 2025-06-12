### 📄 Ghi chú I/O và xử lý dữ liệu trong Go

#### 📦 `io` Package

| Hàm                              | Mô tả                                                                                                                     |
| -------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `Copy(w, r)`                     | Sao chép dữ liệu từ `Reader` sang `Writer` cho đến EOF hoặc lỗi. Trả về số byte và lỗi (nếu có).                          |
| `CopyBuffer(w, r, buffer)`       | Giống như `Copy`, nhưng sử dụng buffer người dùng truyền vào.                                                             |
| `CopyN(w, r, count)`             | Sao chép chính xác `count` byte từ `Reader` sang `Writer`.                                                                |
| `ReadAll(r)`                     | Đọc toàn bộ dữ liệu từ `Reader` cho đến EOF. Trả về slice byte và lỗi (nếu có).                                           |
| `ReadAtLeast(r, byteSlice, min)` | Đọc ít nhất `min` byte vào slice. Báo lỗi nếu không đủ.                                                                   |
| `ReadFull(r, byteSlice)`         | Đọc đúng số byte bằng kích thước slice. Lỗi nếu gặp EOF trước.                                                            |
| `WriteString(w, str)`            | Ghi một chuỗi vào `Writer`.                                                                                               |
| `Pipe()`                         | Tạo ra một cặp `PipeReader` và `PipeWriter`, thường dùng để kết nối giữa 2 goroutine hoặc giữa các luồng đọc/ghi dữ liệu. |
| `MultiReader(...readers)`        | Tạo một `Reader` đọc lần lượt từ nhiều `Reader` khác nhau theo thứ tự.                                                    |
| `MultiWriter(...writers)`        | Tạo một `Writer` ghi cùng lúc vào nhiều `Writer`.                                                                         |
| `LimitReader(r, limit)`          | Trả về một `Reader` đọc tối đa `limit` byte rồi EOF.                                                                      |

---

### 🔄 Các hàm đặc biệt trong xử lý I/O nâng cao

#### 1. `Pipe()`

* **Tạo ra** một cặp `PipeReader` và `PipeWriter` — chúng kết nối với nhau như một đường ống.
* **Ứng dụng chính**: giúp truyền dữ liệu giữa hai goroutine. Một bên ghi (`Writer`) và một bên đọc (`Reader`) cùng lúc.
* 📌 **Lưu ý quan trọng**:

    * Pipe hoạt động **đồng bộ (synchronous)**.
    * Khi gọi `PipeWriter.Write(...)`, hàm sẽ **block (chặn)** lại cho đến khi dữ liệu được đọc từ `PipeReader`.
    * Nếu bạn sử dụng cả `Writer` và `Reader` trong cùng một goroutine, chương trình có thể **bị deadlock (treo)**.
    * ✅ Vì vậy, hãy sử dụng **PipeReader và PipeWriter trong các goroutine khác nhau**.

📝 **Giải thích:**

> "Pipes are synchronous, such that the `PipeWriter.Write` method will block until the data is read from the pipe. This means that the `PipeWriter` needs to be used in a different goroutine from the reader to prevent the application from deadlocking."

```go
r, w := io.Pipe()

go func() {
    w.Write([]byte("hello"))
    w.Close()
}()
data, _ := io.ReadAll(r)
fmt.Println(string(data)) // "hello"
```

#### 2. `MultiReader(...readers)`

* Ghép **nhiều `Reader`** lại thành một `Reader` duy nhất.
* Khi đọc dữ liệu, nó sẽ **đọc lần lượt từng `Reader`**, đến hết cái này mới đọc cái kế tiếp.

```go
r1 := strings.NewReader("Hello, ")
r2 := strings.NewReader("Go!")
r := io.MultiReader(r1, r2)
data, _ := io.ReadAll(r)
fmt.Println(string(data)) // "Hello, Go!"
```

#### 3. `MultiWriter(...writers)`

* Tạo một `Writer` ghi **đồng thời vào tất cả các `Writer`** được truyền vào.

```go
var b1, b2 strings.Builder
w := io.MultiWriter(&b1, &b2)
w.Write([]byte("Test"))
fmt.Println(b1.String(), b2.String()) // "Test Test"
```

📌 **Lưu ý quan trọng khi dùng `MultiWriter(...)`:**

* Cần truyền vào **con trỏ tới các `Writer`** nếu `Writer` đó là một struct (ví dụ `strings.Builder`) để đảm bảo dữ liệu thực sự được ghi đúng.
* Nếu truyền giá trị (không dùng `&`), bạn sẽ chỉ ghi vào bản sao, không phải đối tượng thật.

🧪 **Ví dụ mở rộng:**

```go
var w1, w2, w3 strings.Builder
combinedWriter := io.MultiWriter(&w1, &w2, &w3)
combinedWriter.Write([]byte("Logging!"))
fmt.Println(w1.String()) // "Logging!"
fmt.Println(w2.String()) // "Logging!"
fmt.Println(w3.String()) // "Logging!"
```

#### 4. `LimitReader(r, limit)`

* Tạo một `Reader` **chỉ cho phép đọc tối đa `limit` byte** từ `r`.

```go
r := strings.NewReader("1234567890")
lr := io.LimitReader(r, 4)
data, _ := io.ReadAll(lr)
fmt.Println(string(data)) // "1234"
```

---

### 🧱 strings.Builder và nguyên tắc dùng Reader/Writer

* `strings.Builder` cho phép ghi các byte và sau đó tạo chuỗi bằng `.String()`.
* Writer sẽ trả lỗi nếu không thể ghi toàn bộ dữ liệu, nhưng với `Builder` thì khả năng xảy ra lỗi là rất thấp vì nó ghi vào bộ nhớ.

#### Ví dụ đơn giản:

```go
var builder strings.Builder
builder.WriteString("Hello, ")
builder.WriteString("Go!")
fmt.Println(builder.String()) // "Hello, Go!"
```

#### Khi truyền `Writer` hoặc `Reader` vào hàm:

* Nên truyền **bằng con trỏ** (dùng `&`) để không bị copy giá trị.
* Hầu hết các method của `Writer`/`Reader` hoạt động trên con trỏ.

```go
func processData(r io.Reader, w io.Writer) {
    io.Copy(w, r)
}

reader := strings.NewReader("data")
var builder strings.Builder
processData(reader, &builder)
fmt.Println(builder.String())
```

* `strings.NewReader(...)` trả về con trỏ sẵn, nên không cần `&`.
* `strings.Builder` là struct, nên cần dùng `&builder`.

#### 🛠 Ví dụ thực tế: xử lý từng byte từ Reader và ghi vào Builder

```go
package main

import (
	"io"
	"strings"
	"fmt"
)

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func processData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2) // tạo buffer 2 byte
	for {
		count, err := reader.Read(b) // đọc vào buffer
		if count > 0 {
			writer.Write(b[0:count]) // ghi đúng số byte đã đọc vào writer
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break // kết thúc nếu đọc hết dữ liệu
		}
	}
}

func main() {
	r := strings.NewReader("Kayak") // tạo reader từ chuỗi
	var builder strings.Builder // tạo builder để ghi chuỗi
	processData(r, &builder) // truyền reader và con trỏ builder vào
	Printfln("String builder contents: %s", builder.String())
}
```

---

### 🔧 Phân tích logic: `GenanrateData` và `ConsumeData`

```go
func GenanrateData(writer io.Writer) {
	data := []byte("kayak , lifeJacket")
	writeSize := 4
	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}
		count, err := writer.Write(data[i:end])
		Printfln("Wrote %v byte(s): %v", count, string(data[i:end]))
		if err != nil {
			Printfln("Error: %v", err.Error())
		}
	}
}

func ConsumeData(reader io.Reader) {
	data := make([]byte, 0, 10)
	slice := make([]byte, 2)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			Printfln("Read data: %v", string(slice[0:count]))
			data = append(data, slice[0:count]...)
		}
		if err == io.EOF {
			break
		}
	}
	Printfln("Read data: %v", string(data))
}
```

#### 🧠 Giải thích:

* `GenanrateData(writer)` chia dữ liệu thành từng phần nhỏ (4 byte) và ghi từng phần vào `Writer`.

    * Dùng `slice := data[i:end]` để xử lý từng phần.
    * `writer.Write(...)` ghi phần đó vào đích (ví dụ: buffer, file, pipe...).

* `ConsumeData(reader)` đọc từng phần nhỏ (2 byte) từ `Reader`.

    * Ghi nhận và in ra từng đoạn dữ liệu đã đọc.
    * Gộp các phần lại vào `data` rồi in ra toàn bộ cuối cùng.

✅ Hai hàm này mô phỏng quy trình **ghi dữ liệu theo khối nhỏ** và **đọc dữ liệu theo khối nhỏ**, thường gặp trong xử lý luồng dữ liệu hoặc stream lớn.
---

## 📖 `io.TeeReader` – Đọc và nhân bản dữ liệu đồng thời

### 🧠 Khái niệm

`io.TeeReader` là một hàm trong Go tạo ra một `Reader` **sao chép dữ liệu mà nó đọc ra một Writer khác**. Khi dữ liệu được đọc từ `TeeReader`, nó vừa:

* Trả về dữ liệu như một `Reader` bình thường.
* Đồng thời **ghi dữ liệu đó vào một `Writer` khác**.

### 🧪 Ví dụ minh họa:

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func ConsumeData(r io.Reader) {
	data, _ := io.ReadAll(r)
	fmt.Printf("Consumed data: %s\n", data)
}

func main() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")

	// Nối 3 Reader thành một
	concatReader := io.MultiReader(r1, r2, r3)

	// Tạo Writer để ghi lại nội dung được đọc
	var writer strings.Builder

	// Tạo TeeReader để "phản chiếu" dữ liệu đọc sang writer
	teeReader := io.TeeReader(concatReader, &writer)

	ConsumeData(teeReader)
	fmt.Printf("Echo data: %v\n", writer.String())
}
```

### 💡 Tưởng tượng trực quan:

Hãy tưởng tượng bạn **đọc sách thành tiếng** và **có một người bên cạnh đang ghi lại những gì bạn đọc**:

* `Reader` = bạn đọc sách.
* `TeeReader` = bạn vừa đọc vừa để người khác ghi lại.
* `Writer` = người đang chép lại từng lời bạn đọc.

---

## 📦 `bufio.Reader` – Đọc dữ liệu có đệm

### 🔧 Các phương thức hữu ích:

| Method       | Description                                              |
| ------------ | -------------------------------------------------------- |
| `Buffered()` | Trả về số byte hiện có trong bộ đệm (chưa đọc ra ngoài). |
| `Discard(n)` | Bỏ qua `n` byte tiếp theo trong bộ đệm.                  |
| `Peek(n)`    | Lấy `n` byte tiếp theo mà không xoá khỏi bộ đệm.         |
| `Reset(r)`   | Gán `Reader` mới, xóa dữ liệu đệm hiện tại.              |
| `Size()`     | Trả về kích thước tổng bộ đệm.                           |

### 🧪 Ví dụ:

```go
reader := bufio.NewReaderSize(strings.NewReader("Example"), 10)
peeked, _ := reader.Peek(3) // lấy trước 3 byte
fmt.Println(string(peeked)) // "Exa"
```

### 💡 Tưởng tượng trực quan:

Hãy tưởng tượng `bufio.Reader` giống như một **khay đựng thư đến**:

* `Buffered()` = có bao nhiêu thư sẵn sàng để đọc.
* `Peek(n)` = xem trước nội dung mà không lấy ra.
* `Discard(n)` = vứt bỏ thư không cần đọc.
* `Reset(r)` = thay đổi nguồn thư mới.

---

## ✍️ `bufio.Writer` – Ghi dữ liệu có đệm

### 🛠 Các hàm khởi tạo:

| Function                 | Description                                                  |
| ------------------------ | ------------------------------------------------------------ |
| `NewWriter(w)`           | Tạo `bufio.Writer` với kích thước đệm mặc định (4096 bytes). |
| `NewWriterSize(w, size)` | Tạo `bufio.Writer` với kích thước đệm tùy chỉnh.             |

### 🔍 Các phương thức:

| Method        | Description                                    |
| ------------- | ---------------------------------------------- |
| `Available()` | Số byte còn trống trong buffer.                |
| `Buffered()`  | Số byte đã ghi vào buffer (chưa flush).        |
| `Flush()`     | Ghi nội dung buffer ra `Writer` gốc.           |
| `Reset(w)`    | Đặt lại `Writer` mới và xóa nội dung hiện tại. |
| `Size()`      | Dung lượng tổng của buffer.                    |

### 🧪 Ví dụ:

```go
writer := bufio.NewWriterSize(os.Stdout, 10)
writer.WriteString("Hello")
fmt.Println("Buffered:", writer.Buffered())
writer.WriteString("World")
writer.Flush() // in ra HelloWorld
```

### 💡 Tưởng tượng trực quan:

Tưởng tượng `bufio.Writer` giống như **chai nước**:

* `Write()` = rót nước vào chai.
* `Flush()` = đổ nước từ chai vào bình.
* `Buffered()` = lượng nước đang chứa trong chai.
* `Available()` = còn chỗ trống không?
* `Reset()` = thay đổi bình mới.
* `Size()` = dung tích tối đa của chai.

---
