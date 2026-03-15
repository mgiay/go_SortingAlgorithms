# Thuật toán Timsort

## Giới thiệu
Timsort là một thuật toán sắp xếp lai, kết hợp giữa Insertion Sort và Merge Sort. Nó được thiết kế để tận dụng lợi thế của dữ liệu thực tế thường chứa các đoạn con đã được sắp xếp (gọi là "runs").

## Ý tưởng chính
1. **Chia nhỏ:** Chia mảng thành các đoạn nhỏ gọi là RUN (thường kích thước 32 hoặc 64).
2. **Insertion Sort:** Sắp xếp từng RUN bằng Insertion Sort (rất hiệu quả với mảng nhỏ).
3. **Merge Sort:** Trộn các RUN đã sắp xếp lại với nhau theo quy trình của Merge Sort để tạo thành mảng lớn hơn đã sắp xếp.

## Độ phức tạp
- **Thời gian:**
  - Tốt nhất: O(n) (khi dữ liệu đã sắp xếp sẵn).
  - Trung bình: O(n log n).
  - Xấu nhất: O(n log n).
- **Không gian:** O(n) (cần bộ nhớ phụ để merge).

## Ưu điểm
- Rất nhanh trên dữ liệu thực tế.
- Ổn định (Stable).
- Là thuật toán mặc định của Python, Java (cho Objects), và Android.

## Triển khai (Golang)
File `main.go` triển khai một phiên bản đơn giản hóa của Timsort với kích thước RUN cố định là 32.

## Cách chạy
```bash
go run algorithms/timsort/main.go
```
