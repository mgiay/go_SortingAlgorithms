# Thuật toán Quick Sort (Sắp xếp nhanh)

## Giới thiệu
Quick Sort là một thuật toán sắp xếp dựa trên nguyên tắc "Chia để trị" (Divide and Conquer). Nó chọn một phần tử làm "chốt" (pivot) và phân hoạch mảng thành hai phần: một phần chứa các phần tử nhỏ hơn pivot và phần còn lại chứa các phần tử lớn hơn pivot. Sau đó, thuật toán lặp lại quy trình này cho hai mảng con một cách đệ quy.

## Ý tưởng chính
1. Chọn một phần tử làm pivot (có thể là đầu, cuối, giữa, hoặc ngẫu nhiên).
2. Phân hoạch (Partitioning): Sắp xếp lại mảng sao cho tất cả các phần tử nhỏ hơn pivot đứng trước pivot, và tất cả các phần tử lớn hơn pivot đứng sau pivot.
3. Đệ quy (Recursion): Áp dụng bước trên cho mảng con bên trái và mảng con bên phải của pivot.
4. Điều kiện dừng: Khi mảng con chỉ còn 0 hoặc 1 phần tử.

## Độ phức tạp
- **Thời gian:**
  - Tốt nhất: O(n log n).
  - Trung bình: O(n log n).
  - Xấu nhất: O(n^2) (khi mảng đã sắp xếp hoặc pivot luôn là phần tử lớn nhất/nhỏ nhất).
- **Không gian:** O(log n) (do stack đệ quy).

## Ưu điểm
- Rất nhanh trong thực tế.
- Hoạt động tốt với bộ nhớ cache.
- Có thể thực hiện in-place (tốn ít bộ nhớ phụ).

## Triển khai (Golang)
File `main.go` sử dụng cách chọn pivot là phần tử cuối cùng (Lomuto partition scheme).

## Cách chạy
```bash
go run algorithms/quicksort/main.go
```
