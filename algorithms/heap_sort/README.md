# Thuật toán Heap Sort (Sắp xếp vun đống)

## Giới thiệu
Heap Sort là một thuật toán sắp xếp dựa trên cấu trúc dữ liệu Binary Heap (Đống nhị phân). Nó tương tự như Selection Sort ở chỗ chia mảng thành phần đã sắp xếp và chưa sắp xếp, nhưng nó sử dụng Heap để tìm phần tử lớn nhất nhanh hơn.

## Ý tưởng chính
1. Xây dựng một Max-Heap từ dữ liệu đầu vào.
2. Tại thời điểm này, phần tử lớn nhất nằm ở gốc của heap (vị trí đầu mảng).
3. Hoán đổi phần tử cuối cùng của heap với phần tử gốc.
4. Giảm kích thước của heap đi 1.
5. "Vun đống" (Heapify) lại phần tử gốc mới để đảm bảo tính chất Max-Heap.
6. Lặp lại bước 3-5 cho đến khi heap chỉ còn 1 phần tử.

## Độ phức tạp
- **Thời gian:** O(n log n) trong mọi trường hợp (tốt nhất, trung bình, xấu nhất).
- **Không gian:** O(1) (In-place sort).

## Ưu điểm
- Hiệu suất ổn định, không bị suy biến thành O(n^2) như Quick Sort.
- Không tốn thêm bộ nhớ phụ như Merge Sort.
- Thích hợp cho các hệ thống nhúng (Embedded Systems).

## Triển khai (Golang)
File `main.go` minh họa cách xây dựng Heap và sắp xếp mảng.

## Cách chạy
```bash
go run algorithms/heap_sort/main.go
```
