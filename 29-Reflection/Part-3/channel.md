# 📘 Go reflect: Kiểm tra chiều và kiểu dữ liệu của Channel

Tài liệu này giải thích cách sử dụng package `reflect` trong Go để kiểm tra **channel direction** (chiều gửi/nhận của channel) và **kiểu dữ liệu chứa trong channel** thông qua hai phương thức: `ChanDir()` và `Elem()`.

---

## 🔍 `ChanDir()` – Kiểm tra chiều của channel

Phương thức `ChanDir()` trả về giá trị cho biết channel đó được dùng để **gửi**, **nhận**, hay **cả hai**.

### ✅ Bảng giá trị `ChanDir`:

| Giá trị hằng số (`reflect`) | Ý nghĩa                       | Biểu diễn dạng chuỗi |
| --------------------------- | ----------------------------- | -------------------- |
| `reflect.RecvDir`           | Chỉ nhận dữ liệu              | `<-chan T`           |
| `reflect.SendDir`           | Chỉ gửi dữ liệu               | `chan<- T`           |
| `reflect.BothDir`           | Gửi và nhận dữ liệu (2 chiều) | `chan T`             |

---

## 📦 `Elem()` – Kiểm tra kiểu dữ liệu chứa trong channel

Phương thức `Elem()` trả về **kiểu dữ liệu** mà channel chứa (loại dữ liệu được gửi/nhận qua channel).

Ví dụ:

```go
ch := make(chan int)
t := reflect.TypeOf(ch)
fmt.Println(t.Elem()) // Output: int
```

# 📘 Go reflect: Các phương thức thao tác với channel bằng reflection

Gói `reflect` trong Go cho phép bạn thao tác với channel một cách linh hoạt trong runtime. Dưới đây là các phương thức quan trọng của `reflect.Value` khi làm việc với channel.

---

## 📦 1. `Send(val reflect.Value)`

- **Mục đích**: Gửi một giá trị vào channel.
- **Chặn**: Có. Hàm sẽ chờ đến khi gửi được.
- **Cách dùng**:

```go
v := reflect.ValueOf(myChan)
v.Send(reflect.ValueOf(10))
```

| Thành phần           | Mục đích                                   |
| -------------------- | ------------------------------------------ |
| `reflect.Select`     | Chạy select động từ slice các `SelectCase` |
| `reflect.SelectCase` | Mô tả từng dòng case trong select          |
| `Chan`               | Channel sẽ thao tác                        |
| `Dir`                | Hướng: gửi, nhận, hay default              |
| `Send`               | Giá trị sẽ gửi nếu là thao tác `Send`      |
| `SelectSend`         | Gửi dữ liệu vào channel                    |
| `SelectRecv`         | Nhận dữ liệu từ channel                    |
| `SelectDefault`      | Mặc định nếu không channel nào sẵn sàng    |
