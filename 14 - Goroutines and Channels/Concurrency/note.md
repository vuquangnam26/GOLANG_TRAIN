# README: Tổng hợp kiến thức Channel & Goroutines trong Go

## I. Goroutine là gì?

* Goroutine là luồng nhẹ (lightweight thread) do Go runtime quản lý.
* Tạo bằng từ khóa `go`, ví dụ:

  ```go
  go sayHello()
  ```

## II. Channel là gì?

* Là phương tiện giao tiếp an toàn giữa các goroutines.
* Tạo channel:

  ```go
  ch := make(chan int)           // unbuffered
  ch := make(chan int, 3)        // buffered (đệm 3 giá trị)
  ```
* Gửi và nhận:

  ```go
  ch <- 5        // Gửi 5 vào channel
  value := <-ch  // Nhận từ channel
  ```

## III. Giới hạn chiều của channel

* Có thể giới hạn chiều sử dụng:

  ```go
  func sendData(ch chan<- int)     // chỉ gửi
  func receiveData(ch <-chan int)  // chỉ nhận
  ```

## IV. Buffered vs Unbuffered Channel

| Loại       | Đặc điểm                            |
| ---------- | ----------------------------------- |
| Unbuffered | Gửi block đến khi có goroutine nhận |
| Buffered   | Gửi không block nếu buffer chưa đầy |

## V. Dùng `select` để xử lý nhiều channel

```go
select {
case val := <-ch1:
    fmt.Println("Received", val)
case ch2 <- 10:
    fmt.Println("Sent")
default:
    fmt.Println("Nothing ready")
}
```

* `select` chọn ngẫu nhiên 1 case sẵn sàng, nếu không có:

    * Có `default` → thực hiện `default`
    * Không có → block

## VI. Đóng channel

* Gửi xong thì `close(channel)`
* Nhận:

  ```go
  for val := range ch {
    fmt.Println(val)
  }
  ```
* Kiểm tra channel đóng:

  ```go
  val, ok := <-ch
  if !ok {
    fmt.Println("Đã đóng")
  }
  ```

## VII. Tránh lỗi khi channel đóng

* Channel đóng vẫn có thể nhận (trả zero-value)
* Dùng `ok` để kiểm tra channel đã đóng hay chưa

## VIII. Gửi nhiều giá trị và bỏ qua khi channel đầy

```go
select {
case ch <- val:
    fmt.Println("Sent")
default:
    fmt.Println("Discarded")
}
```

## IX. Điều phối channel và thoát vòng lặp

* Dùng biến đếm số lượng channel đang mở
* Khi channel đóng thì giảm biến đếm
* Khi biến đếm = 0 thì `break` vòng lặp

## X. Ghi chú thực chiến

| Tình huống                       | Giải pháp                                   |
| -------------------------------- | ------------------------------------------- |
| Gửi không có người nhận          | Dùng buffered channel hoặc goroutine nhận   |
| Gửi khi đầy buffer               | Dùng `select` + `default` để tránh block    |
| Nhận từ channel đã đóng          | Trả zero-value, dùng `ok` để kiểm tra       |
| Channel block mãi không kết thúc | Dùng `close`, `break`, hoặc `goto` để thoát |

---

> Tổng hợp này giúp bạn hiểu và áp dụng hiệu quả channel và goroutines trong lập trình song song với Go.
