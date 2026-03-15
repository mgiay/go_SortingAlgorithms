# Thuật toán Shell Sort (Sắp xếp Shell)

## Giới thiệu
Shell Sort là một cải tiến của Insertion Sort. Trong Insertion Sort, các phần tử chỉ di chuyển từng bước một. Shell Sort cho phép trao đổi các phần tử ở xa nhau để sắp xếp một cách nhanh chóng hơn.

## Ý tưởng chính
1. Sắp xếp các cặp phần tử cách xa nhau một khoảng `gap`.
2. Giảm dần khoảng cách `gap` (thường chia đôi mỗi lần lặp).
3. Lặp lại quá trình sắp xếp chèn (Insertion Sort) với các khoảng cách nhỏ hơn.
4. Bước cuối cùng là `gap = 1`, lúc này thuật toán trở thành Insertion Sort thông thường, nhưng mảng đã gần như được sắp xếp hoàn toàn nên chạy rất nhanh.

## Độ phức tạp
- **Thời gian:** Phụ thuộc vào dãy khoảng cách (gap sequence).
  - Tốt nhất: O(n log n).
  - Trung bình: O(n log n) hoặc O(n^(5/4)).
  - Xấu nhất: O(n^2) (với dãy chia đôi đơn giản), có thể đạt O(n log^2 n) với dãy tối ưu hơn.
- **Không gian:** O(1) (In-place sort).

## Ưu điểm
- Nhanh hơn Insertion Sort nhiều lần với mảng lớn hơn.
- Cài đặt đơn giản, không cần đệ quy.
- Hiệu quả bộ nhớ.

## Triển khai (Golang)
File `main.go` sử dụng dãy khoảng cách chia đôi đơn giản (`n/2`, `n/4`, ... `1`).

## Cách chạy
```bash
go run algorithms/shell_sort/main.go
```
