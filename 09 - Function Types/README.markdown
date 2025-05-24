# README: Tính Chất của Closure trong Go (Golang)

**Closure** (hàm đóng) là một tính năng quan trọng trong Go, cho phép tạo ra các hàm ẩn danh (anonymous functions) có khả năng truy cập các biến từ phạm vi bên ngoài. Trong tài liệu này, tôi sẽ trình bày chi tiết các tính chất của closure trong Go, liên hệ với đoạn code bạn cung cấp, và cung cấp ví dụ minh họa để làm rõ.

## 1. Định nghĩa Closure trong Go
Closure trong Go là một hàm ẩn danh được định nghĩa bên trong một hàm khác, có khả năng truy cập và sử dụng các biến từ phạm vi bao quanh (scope của hàm cha hoặc biến toàn cục). Closure "đóng gói" (capture) các biến này, cho phép hàm tiếp tục sử dụng chúng ngay cả sau khi hàm cha kết thúc.

- **Ví dụ cơ bản**:
```go
package main
import "fmt"

func outer() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

func main() {
    counter := outer()
    fmt.Println(counter()) // In: 1
    fmt.Println(counter()) // In: 2
}
```

## 2. Tính Chất của Closure trong Go

### 2.1. Truy cập biến từ phạm vi bên ngoài
Closure có thể truy cập các biến từ:
- **Biến cục bộ** trong hàm cha.
- **Tham số** của hàm cha.
- **Biến toàn cục** được định nghĩa ngoài hàm.

**Trong code của bạn**:
```go
func priceCalcFactory(threshold, rate float64) calcFunc {
    return func(price float64) float64 {
        if prizeGiveaway { // Biến toàn cục
            return 0
        } else if price > threshold { // Tham số của hàm cha
            return price + (price * rate) // Tham số của hàm cha
        }
        return price
    }
}
```
- Closure truy cập:
  - `prizeGiveaway` (biến toàn cục).
  - `threshold` và `rate` (tham số của `priceCalcFactory`).

### 2.2. Tham chiếu, không sao chép giá trị
Closure trong Go **tham chiếu** (reference) đến các biến, không sao chép giá trị của chúng tại thời điểm tạo. Do đó, closure sử dụng giá trị **hiện tại** của biến khi được gọi.

**Trong code của bạn**:
```go
var prizeGiveaway = false

func main() {
    prizeGiveaway = false
    waterCalc := priceCalcFactory(100, 0.2) // Tạo khi prizeGiveaway = false
    prizeGiveaway = true // Thay đổi giá trị
    soccerCalc := priceCalcFactory(50, 0.1)
    fmt.Println(waterCalc(275)) // In: 0 (vì prizeGiveaway = true khi gọi)
}
```
- **Giải thích**:
  - Khi tạo `waterCalc`, `prizeGiveaway = false`, nhưng closure không lưu giá trị này.
  - Khi gọi `waterCalc(275)`, closure kiểm tra `prizeGiveaway` (lúc này là `true`), nên trả về 0.
  - Điều này giải thích tại sao bạn thấy `prizeGiveaway = true` cả hai lần khi in trong vòng lặp.

### 2.3. Lưu trữ trạng thái riêng cho biến cục bộ/tham số
Mỗi closure có thể lưu trữ trạng thái riêng cho các biến cục bộ hoặc tham số của hàm cha. Điều này cho phép tạo nhiều closure với các giá trị khác nhau.

**Trong code của bạn**:
```go
waterCalc := priceCalcFactory(100, 0.2) // Lưu threshold = 100, rate = 0.2
soccerCalc := priceCalcFactory(50, 0.1)  // Lưu threshold = 50, rate = 0.1
```
- Mỗi closure (`waterCalc`, `soccerCalc`) giữ các giá trị riêng của `threshold` và `rate`.
- Tuy nhiên, cả hai đều chia sẻ `prizeGiveaway` (biến toàn cục), nên đều thấy `prizeGiveaway = true` khi được gọi.

### 2.4. Tồn tại sau khi hàm cha kết thúc
Closure tiếp tục tồn tại và giữ quyền truy cập các biến được capture ngay cả sau khi hàm cha (như `priceCalcFactory`) kết thúc.

**Ví dụ**:
```go
func makeAdder(add int) func(int) int {
    return func(x int) int {
        return x + add // add vẫn tồn tại dù makeAdder đã kết thúc
    }
}

func main() {
    add5 := makeAdder(5)
    fmt.Println(add5(10)) // In: 15
}
```
- Closure `add5` giữ giá trị `add = 5` và tiếp tục sử dụng nó.

**Trong code của bạn**:
- Sau khi `priceCalcFactory` trả về `waterCalc` và `soccerCalc`, các closure này vẫn giữ `threshold` và `rate`, đồng thời tham chiếu `prizeGiveaway`.

### 2.5. Linh hoạt trong việc tạo hàm động
Closure cho phép tạo các hàm với hành vi khác nhau dựa trên các biến được capture, rất hữu ích trong việc tạo các hàm động hoặc xử lý logic phức tạp.

**Trong code của bạn**:
- `priceCalcFactory` tạo ra các hàm (`waterCalc`, `soccerCalc`) với các ngưỡng và tỷ lệ khác nhau, giúp áp dụng logic giá khác nhau cho các danh mục sản phẩm.

## 3. Ứng dụng của Closure trong Go
- **Tạo hàm động**: Như trong code của bạn, dùng closure để tạo các hàm tính giá linh hoạt.
- **Quản lý trạng thái**: Giữ trạng thái riêng cho từng closure (ví dụ: bộ đếm, bộ lọc).
- **Xử lý callback**: Dùng trong các API hoặc xử lý sự kiện.
- **Trì hoãn thực thi**: Cho phép trì hoãn việc thực thi một hàm với trạng thái được lưu trước.

## 4. Lưu ý khi sử dụng Closure trong Go
- **Biến toàn cục**: Nếu closure sử dụng biến toàn cục (như `prizeGiveaway`), tất cả closure sẽ chia sẻ giá trị hiện tại của biến đó. Để tránh điều này, truyền biến như tham số:
```go
func priceCalcFactory(threshold, rate float64, giveaway bool) calcFunc {
    return func(price float64) float64 {
        if giveaway { // Dùng giá trị giveaway lúc tạo
            return 0
        } else if price > threshold {
            return price + (price * rate)
        }
        return price
    }
}
```
- **Hiệu suất**: Closure nhẹ, nhưng cần cẩn thận khi capture nhiều biến hoặc trong các vòng lặp lớn, vì có thể gây rò rỉ bộ nhớ nếu không quản lý tốt.

## 5. Kết quả trong code của bạn
Dựa trên code bạn cung cấp:
```go
package main
import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
    fmt.Println("Product:", product, "Price:", calculator(price))
}

var prizeGiveaway = false

func priceCalcFactory(threshold, rate float64) calcFunc {
    return func(price float64) float64 {
        if prizeGiveaway {
            return 0
        } else if price > threshold {
            return price + (price * rate)
        }
        return price
    }
}

func main() {
    watersportsProducts := map[string]float64{
        "Kayak":      275,
        "Lifejacket": 48.95,
    }
    soccerProducts := map[string]float64{
        "Soccer Ball": 19.50,
        "Stadium":     79500,
    }
    prizeGiveaway = false
    waterCalc := priceCalcFactory(100, 0.2)
    prizeGiveaway = true
    soccerCalc := priceCalcFactory(50, 0.1)
    for product, price := range watersportsProducts {
        printPrice(product, price, waterCalc)
    }
    for product, price := range soccerProducts {
        printPrice(product, price, soccerCalc)
    }
}
```
- **Kết quả**:
```
Product: Kayak Price: 0
Product: Lifejacket Price: 0
Product: Soccer Ball Price: 0
Product: Stadium Price: 0
```
- **Lý do**: Do tính chất **tham chiếu** của closure, cả `waterCalc` và `soccerCalc` dùng giá trị hiện tại của `prizeGiveaway` (là `true`) khi được gọi, dẫn đến tất cả giá là 0.

## 6. Kết luận
Closure trong Go là một công cụ mạnh mẽ với các tính chất:
1. Truy cập biến từ phạm vi bên ngoài (toàn cục, cục bộ, tham số).
2. Tham chiếu giá trị hiện tại, không sao chép.
3. Lưu trữ trạng thái riêng cho biến cục bộ/tham số.
4. Tồn tại sau khi hàm cha kết thúc.
5. Linh hoạt trong việc tạo hàm động.

Hiểu rõ các tính chất này giúp bạn sử dụng closure hiệu quả, đặc biệt trong các trường hợp như code của bạn, nơi closure được dùng để tạo logic tính giá linh hoạt. Nếu bạn cần thêm ví dụ hoặc muốn sửa code để đạt kết quả khác, hãy liên hệ!