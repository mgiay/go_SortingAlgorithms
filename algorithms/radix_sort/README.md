# Thuật toán Radix Sort (Sắp xếp cơ số)

## Giới thiệu
Radix Sort là một thuật toán sắp xếp không dựa trên so sánh (non-comparative sort). Thay vì so sánh các phần tử với nhau, nó phân loại các số dựa trên từng chữ số (digit) của chúng.

## Ý tưởng chính
1. Xác định số lượng chữ số tối đa trong mảng (dựa vào số lớn nhất).
2. Bắt đầu từ chữ số có ý nghĩa thấp nhất (hàng đơn vị) đến chữ số có ý nghĩa cao nhất.
3. Sử dụng một thuật toán sắp xếp ổn định (thường là Counting Sort) để sắp xếp các số dựa trên chữ số hiện tại đang xét.
4. Sau khi duyệt qua tất cả các hàng (đơn vị, chục, trăm...), mảng sẽ được sắp xếp hoàn toàn.

## Độ phức tạp
- **Thời gian:** O(nk), trong đó n là số phần tử và k là số lượng chữ số tối đa của số lớn nhất.
- **Không gian:** O(n + k) cho mảng phụ của Counting Sort.

## Ưu điểm
- Có thể nhanh hơn các thuật toán so sánh (O(n log n)) khi k nhỏ.
- Ổn định (Stable).

## Nhược điểm
- Không sắp xếp được số thực (float) hoặc số âm một cách trực tiếp nếu không có biến đổi.
- Tốn bộ nhớ hơn Quick Sort in-place.

## Triển khai (Golang)
File `main.go` triển khai thuật toán LSD (Least Significant Digit) Radix Sort cho các số nguyên dương.

## Cách chạy
```bash
go run algorithms/radix_sort/main.go
```
