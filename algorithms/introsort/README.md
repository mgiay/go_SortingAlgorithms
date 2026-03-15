# Thuật toán Introsort (Introspective Sort)

## Giới thiệu
Introsort là một thuật toán sắp xếp lai (hybrid sorting algorithm) kết hợp giữa Quick Sort, Heap Sort và Insertion Sort. Nó bắt đầu với Quick Sort, nhưng chuyển sang Heap Sort khi độ sâu đệ quy vượt quá một mức nhất định (để tránh trường hợp xấu nhất của Quick Sort), và chuyển sang Insertion Sort khi kích thước mảng con nhỏ.

## Ý tưởng chính
1. **Quick Sort:** Sử dụng làm thuật toán chính để phân chia mảng.
2. **Heap Sort:** Nếu độ sâu đệ quy của Quick Sort vượt quá `2 * log(n)`, chuyển sang Heap Sort. Điều này đảm bảo độ phức tạp trong trường hợp xấu nhất luôn là O(n log n).
3. **Insertion Sort:** Khi kích thước của mảng con nhỏ (thường là < 16), sử dụng Insertion Sort vì nó nhanh hơn các thuật toán phức tạp trên dữ liệu nhỏ.

## Độ phức tạp
- **Thời gian:** O(n log n) trong trường hợp xấu nhất.
- **Không gian:** O(log n) cho stack đệ quy.

## Ưu điểm
- Kết hợp tốc độ trung bình tuyệt vời của Quick Sort và sự an toàn trong trường hợp xấu nhất của Heap Sort.
- Là thuật toán mặc định trong `std::sort` của C++.

## Triển khai (Golang)
File `main.go` minh họa cách kết hợp 3 thuật toán này.

## Cách chạy
```bash
go run algorithms/introsort/main.go
```
