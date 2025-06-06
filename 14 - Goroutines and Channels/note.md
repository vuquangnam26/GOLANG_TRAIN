Ghi Chú Về Goroutines và Channels Trong Go
- các hàm chạy không đồng bộ , khi logic trong  một hàm có channel thực hiện xong thì sẽ kết thúc một routien
1. Goroutines Là Gì?

Goroutine là một đơn vị thực thi nhẹ (lightweight thread) trong Go, cho phép chạy các hàm đồng thời (concurrently).
Cách sử dụng: Thêm từ khóa go trước lời gọi hàm, ví dụ: go myFunction().
Lưu ý:
Goroutine không đảm bảo thứ tự thực thi, phụ thuộc vào lập lịch của Go runtime.
Khi goroutine chính (main) kết thúc, chương trình dừng, bất kể các goroutine khác còn chạy hay không.



2. Channels Là Gì?

Channel là cơ chế đồng bộ và giao tiếp giữa các goroutine, cho phép gửi và nhận dữ liệu.
Cú pháp khai báo: ch := make(chan Type) (unbuffered) hoặc ch := make(chan Type, capacity) (buffered).
Lưu ý:
Unbuffered channel: Gửi (ch <- value) và nhận (<-ch) chặn nhau cho đến khi cả hai sẵn sàng.
Buffered channel: Gửi không chặn cho đến khi bộ đệm đầy; nhận không chặn nếu có dữ liệu trong bộ đệm.



3. Các Trường Hợp Sử Dụng Goroutines và Channels
   3.1. Trường Hợp Sử Dụng Goroutines

Thực hiện tác vụ song song:
Dùng khi cần chạy nhiều tác vụ độc lập cùng lúc, ví dụ: xử lý nhiều request mạng, tính toán trên các tập dữ liệu lớn.
Ví dụ: Tính tổng phụ của nhiều danh mục sản phẩm đồng thời (như trong CalcStoreTotal).


Xử lý tác vụ nền:
Chạy các tác vụ không cần kết quả ngay lập tức, như ghi log, gửi email.
Ví dụ: go writeLog("message").


Tăng hiệu suất:
Phân chia công việc CPU-bound hoặc I/O-bound thành các phần nhỏ chạy đồng thời.
Ví dụ: Xử lý song song các file trong thư mục.



3.2. Trường Hợp Sử Dụng Channels

Đồng bộ hóa goroutines:
Dùng để đảm bảo các goroutine hoàn thành trước khi tiếp tục.
Ví dụ: Trong CalcStoreTotal, dùng channel để nhận tổng phụ từ các goroutine trước khi in tổng cuối cùng.


Truyền dữ liệu giữa goroutines:
Gửi kết quả từ một goroutine đến goroutine khác an toàn, tránh race condition.
Ví dụ: Gửi tổng phụ từ TotalPrice qua channel đến goroutine chính.


Điều phối luồng công việc (worker pool):
Phân phối công việc cho một nhóm goroutine (workers) xử lý đồng thời.
Ví dụ: Xử lý hàng loạt URL tải về, mỗi goroutine lấy URL từ channel và xử lý.


Tín hiệu (signaling):
Dùng channel để báo hiệu sự kiện, như hoàn thành tác vụ hoặc lỗi.
Ví dụ: Channel kiểu chan struct{} để báo hiệu goroutine hoàn thành.



3.3. Kết Hợp Goroutines và Channels

Xử lý tác vụ đồng thời và thu thập kết quả:
Khởi chạy nhiều goroutine để thực hiện công việc, dùng channel để thu thập kết quả.
Ví dụ: Tính tổng phụ các danh mục sản phẩm, gửi kết quả qua channel (như trong mã sửa lỗi của bạn).


Tránh deadlock:
Đảm bảo số lần gửi và nhận trên channel khớp nhau.
Ví dụ sai trong mã gốc: Nhận len(data)+1 lần trong khi chỉ gửi len(data) lần, gây deadlock.


Kiểm soát thứ tự thực thi:
Dùng channel để đảm bảo các bước trong quy trình thực thi đúng thứ tự.
Ví dụ: Nhận kết quả từ tất cả goroutine trước khi in tổng.



4. Lưu Ý Quan Trọng Khi Dùng Goroutines và Channels

Tránh Deadlock:
Đảm bảo tất cả goroutine gửi dữ liệu vào channel đều có nhận tương ứng và ngược lại.
Ví dụ sai: Vòng lặp nhận len(data)+1 lần trong khi chỉ có len(data) goroutine gửi.
Cách sửa: Khởi chạy tất cả goroutine trước, sau đó nhận đúng số lần (for range data).


Đóng Channel:
Đóng channel (close(ch)) khi không còn dữ liệu để gửi, tránh goroutine nhận bị chặn mãi mãi.
Kiểm tra channel đóng bằng v, ok := <-ch.


Sử Dụng WaitGroup Thay Thế:
Nếu không cần truyền dữ liệu, sync.WaitGroup có thể thay channel để đợi goroutine hoàn thành.
Ví dụ: Dùng WaitGroup để đợi tất cả TotalPrice hoàn thành, cập nhật tổng bằng mutex.


Race Condition:
Tránh truy cập biến chia sẻ trực tiếp trong goroutine mà không có khóa (mutex).
Channel là cách an toàn hơn để truyền dữ liệu.


Buffered vs Unbuffered Channels:
Unbuffered: Phù hợp khi cần đồng bộ chặt chẽ (như trong CalcStoreTotal).
Buffered: Dùng khi muốn gửi dữ liệu mà không cần nhận ngay, nhưng phải cẩn thận với kích thước bộ đệm.



5. Ví Dụ Minh Họa

Mã gốc (sai): Vòng lặp lồng trong CalcStoreTotal gây deadlock vì nhận quá nhiều giá trị trước khi khởi chạy tất cả goroutine.
Mã sửa:func CalcStoreTotal(data ProductData) {
var storeTotal float64
var channel chan float64 = make(chan float64)
for category, group := range data {
go group.TotalPrice(category, channel)
}
for range data {
storeTotal += <-channel
}
fmt.Println("Total:", ToCurrency(storeTotal))
}


Trường hợp khác:
Worker pool: Tạo nhiều goroutine nhận công việc từ một channel công việc (job channel) và gửi kết quả vào channel kết quả (result channel).
Pipeline: Dùng nhiều channel để xử lý dữ liệu qua các giai đoạn (stages).



6. Khi Nào Không Nên Dùng Goroutines/Channels

Tác vụ tuần tự đơn giản: Nếu tác vụ không cần đồng thời, dùng hàm thông thường để tránh phức tạp.
Overhead của goroutine: Với tác vụ rất nhỏ, overhead của việc tạo goroutine có thể lớn hơn lợi ích.
Dữ liệu không cần chia sẻ: Nếu goroutine không cần giao tiếp, có thể không cần channel.

