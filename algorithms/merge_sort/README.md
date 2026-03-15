# Thuật toán Merge Sort (Sắp xếp trộn)

## Giới thiệu
Merge Sort là một thuật toán sắp xếp dựa trên nguyên tắc "Chia để trị" (Divide and Conquer). Nó chia mảng thành hai nửa, sắp xếp đệ quy từng nửa, và sau đó trộn (merge) hai nửa đã sắp xếp lại với nhau.

## Ý tưởng chính
1. Chia (Divide): Chia mảng thành hai nửa bằng nhau.
2. Trị (Conquer): Gọi đệ quy Merge Sort cho từng nửa mảng.
3. Trộn (Combine): Gộp hai nửa mảng đã sắp xếp thành một mảng duy nhất.

## Độ phức tạp
- **Thời gian:** O(n log n) trong mọi trường hợp (tốt nhất, trung bình, xấu nhất).
- **Không gian:** O(n) (cần mảng phụ để lưu kết quả tạm thời).

## Ưu điểm
- Ổn định (Stable): Giữ nguyên thứ tự của các phần tử bằng nhau.
- Hiệu suất ổn định, không phụ thuộc vào dữ liệu đầu vào.
- Thích hợp cho sắp xếp danh sách liên kết (Linked List) hoặc dữ liệu lớn không nằm gọn trong RAM (External Sorting).

## Triển khai (Golang)
File `main.go` minh họa cách triển khai Merge Sort bằng đệ quy.

## Cách chạy
```bash
go run algorithms/merge_sort/main.go
```
