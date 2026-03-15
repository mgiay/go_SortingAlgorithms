# Thuật toán Bubble Sort (Sắp xếp nổi bọt)

## Giới thiệu
Bubble Sort là một thuật toán sắp xếp đơn giản nhất. Nó hoạt động bằng cách lặp đi lặp lại việc đi qua danh sách cần sắp xếp, so sánh từng cặp phần tử liền kề và hoán đổi chúng nếu chúng theo thứ tự sai. Quá trình này được lặp lại cho đến khi không cần hoán đổi nào nữa, điều đó có nghĩa là danh sách đã được sắp xếp.

Tên gọi "Nổi bọt" bắt nguồn từ việc các phần tử nhỏ hơn sẽ "nổi" dần lên đầu danh sách qua từng bước (hoặc phần tử lớn hơn "chìm" xuống cuối).

## Ý tưởng chính
1. Bắt đầu từ phần tử đầu tiên, so sánh nó với phần tử thứ hai.
2. Nếu phần tử đầu > phần tử thứ hai, hoán đổi chúng.
3. Tiếp tục so sánh cặp thứ 2 và thứ 3, cứ thế cho đến hết mảng.
4. Sau vòng lặp đầu tiên, phần tử lớn nhất sẽ nằm ở cuối mảng.
5. Lặp lại quy trình cho các phần tử còn lại (trừ phần tử cuối cùng đã được sắp xếp).
6. Dừng lại khi không còn cặp nào cần hoán đổi.

## Độ phức tạp
- **Thời gian:**
  - Tốt nhất: O(n) (khi mảng đã sắp xếp sẵn và có cờ kiểm tra `swapped`).
  - Trung bình: O(n^2).
  - Xấu nhất: O(n^2).
- **Không gian:** O(1) (In-place sort).

## Triển khai (Golang)
File `main.go` chứa mã nguồn triển khai. Chúng tôi sử dụng một biến cờ `swapped` để tối ưu hóa thuật toán, giúp nó dừng sớm nếu mảng đã được sắp xếp.

## Cách chạy
```bash
go run algorithms/bubble_sort/main.go
```
