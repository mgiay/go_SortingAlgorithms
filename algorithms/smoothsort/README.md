# Thuật toán Smoothsort (Phiên bản đơn giản hóa)

## Giới thiệu
Smoothsort là một thuật toán sắp xếp dựa trên so sánh, là một biến thể của Heapsort. Nó được phát minh bởi Edsger W. Dijkstra vào năm 1981.

## Ý tưởng chính (Smoothsort gốc)
Smoothsort cải tiến Heapsort bằng cách sử dụng một cấu trúc dữ liệu heap đặc biệt dựa trên các số Leonardo (thay vì heap nhị phân thông thường). Điều này cho phép thuật toán tận dụng lợi thế khi mảng đầu vào đã được sắp xếp một phần (adaptive).

- **Số Leonardo:** L(0)=1, L(1)=1, L(n) = L(n-1) + L(n-2) + 1.
- **Thích nghi:** Khi mảng đã sắp xếp, độ phức tạp của Smoothsort tiệm cận O(n).

## Lưu ý về triển khai
Do độ phức tạp rất cao của thuật toán Smoothsort gốc (sử dụng Leonardo Heap và các thao tác bitwise phức tạp để quản lý rừng heap), trong dự án này chúng tôi cung cấp một phiên bản **đơn giản hóa** hoạt động dựa trên Binary Heap (tương tự Heapsort). Điều này đảm bảo tính đúng đắn và ổn định của mã nguồn, đồng thời vẫn minh họa được tư tưởng sắp xếp dựa trên cấu trúc cây.

Nếu bạn cần sử dụng Smoothsort nguyên bản cho các hệ thống yêu cầu độ thích nghi cao (adaptive), hãy cân nhắc sử dụng các thư viện chuyên dụng đã được kiểm thử kỹ lưỡng.

## Độ phức tạp (Phiên bản này)
- **Thời gian:** O(n log n).
- **Không gian:** O(1).

## Cách chạy
```bash
go run algorithms/smoothsort/main.go
```
