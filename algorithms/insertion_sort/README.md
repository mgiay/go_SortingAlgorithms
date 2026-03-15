# Thuật toán Insertion Sort (Sắp xếp chèn)

## Giới thiệu
Insertion Sort là một thuật toán sắp xếp đơn giản, hoạt động tương tự như cách bạn sắp xếp các lá bài trên tay. Mảng được chia thành hai phần: phần đã sắp xếp (bên trái) và phần chưa sắp xếp (bên phải). Các giá trị từ phần chưa sắp xếp được chọn và đặt vào đúng vị trí trong phần đã sắp xếp.

## Ý tưởng chính
1. Bắt đầu từ phần tử thứ 2 (chỉ số 1) của mảng.
2. So sánh phần tử này với các phần tử trong danh sách con đã sắp xếp (bên trái).
3. Di chuyển tất cả các phần tử lớn hơn giá trị đang xét sang phải một vị trí để tạo khoảng trống.
4. Chèn giá trị đang xét vào vị trí đó.
5. Lặp lại cho đến khi hết mảng.

## Độ phức tạp
- **Thời gian:**
  - Tốt nhất: O(n) (khi mảng đã sắp xếp sẵn).
  - Trung bình: O(n^2).
  - Xấu nhất: O(n^2).
- **Không gian:** O(1) (In-place sort).

## Ưu điểm
- Rất hiệu quả với các mảng nhỏ.
- Thích nghi tốt (adaptive): nếu mảng đã gần sắp xếp, nó chạy rất nhanh.
- Ổn định (Stable): giữ nguyên thứ tự tương đối của các phần tử bằng nhau.

## Triển khai (Golang)
File `main.go` minh họa cách hoạt động của thuật toán.

## Cách chạy
```bash
go run algorithms/insertion_sort/main.go
```
