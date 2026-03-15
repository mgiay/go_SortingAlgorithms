# Sorting Algorithms Collection in Golang

Dự án này là bộ sưu tập các thuật toán sắp xếp nổi tiếng được triển khai bằng ngôn ngữ **Go (Golang)**. Bên cạnh việc cung cấp mã nguồn chuẩn, dự án còn tích hợp các công cụ **Benchmark** hiệu năng và một **Web Visualizer** sống động với âm thanh.

---

## 📚 Danh sách Thuật toán

Dự án bao gồm 10 thuật toán sắp xếp phổ biến, mỗi thuật toán được đặt trong thư mục riêng tại `algorithms/` với tài liệu chi tiết:

1.  **[Bubble Sort](algorithms/bubble_sort/README.md):** Đơn giản, dễ hiểu, phù hợp cho giáo dục.
2.  **[Insertion Sort](algorithms/insertion_sort/README.md):** Hiệu quả với dữ liệu nhỏ hoặc gần sắp xếp.
3.  **[Shell Sort](algorithms/shell_sort/README.md):** Cải tiến của Insertion Sort với các khoảng cách (gap).
4.  **[Quick Sort](algorithms/quicksort/README.md):** Nhanh, phổ biến, sử dụng chiến lược chia để trị.
5.  **[Merge Sort](algorithms/merge_sort/README.md):** Ổn định, hiệu suất O(n log n) đảm bảo.
6.  **[Heap Sort](algorithms/heap_sort/README.md):** Sử dụng cấu trúc Heap, không tốn bộ nhớ phụ.
7.  **[Radix Sort](algorithms/radix_sort/README.md):** Sắp xếp không so sánh, dựa trên chữ số.
8.  **[Introsort](algorithms/introsort/README.md):** Lai giữa Quick, Heap và Insertion Sort (dùng trong C++ STL).
9.  **[Timsort](algorithms/timsort/README.md):** Lai giữa Merge và Insertion Sort (dùng trong Python, Java).
10. **[Smoothsort](algorithms/smoothsort/README.md):** Biến thể thích nghi của Heapsort (phiên bản đơn giản hóa).

---

## 🛠️ Tính năng nổi bật

### 1. Benchmark & Kiểm thử
Mỗi thuật toán đều được tích hợp sẵn code để:
- **Kiểm thử đúng đắn:** Chạy trên mảng nhỏ cố định để kiểm tra logic.
- **Đo hiệu năng:** Chạy trên mảng lớn (10.000 phần tử) được sinh ngẫu nhiên với seed cố định (đảm bảo công bằng khi so sánh).

Cách chạy thử một thuật toán:
```bash
go run algorithms/quicksort/main.go
```

### 2. Script Build Đa nền tảng (`go_build.sh`)
Script `go_build.sh` giúp biên dịch mã nguồn thành file thực thi (binary) tối ưu cho nhiều hệ điều hành:
- **Hỗ trợ:** Linux (amd64/arm64), Windows, macOS (Intel/Apple Silicon).
- **Tối ưu:** Tự động nén file (UPX), loại bỏ thông tin debug để giảm kích thước.
- **Bảo mật:** Tự động tạo file checksum `SHA256SUMS` để kiểm tra toàn vẹn.

Cách sử dụng:
```bash
./go_build.sh algorithms/bubble_sort/main.go
```
Kết quả sẽ nằm trong thư mục `bin/bubble_sort/`.

### 3. Web Visualizer (Trực quan hóa)
Một ứng dụng web giúp bạn "nhìn" và "nghe" thấy thuật toán hoạt động.

*   **Công nghệ:** Go Backend (SSE) + HTML/JS Frontend + Web Audio API.
*   **Tính năng:**
    *   Mô phỏng Bubble, Quick, Heap, Merge, Insertion, Selection Sort.
    *   **Âm thanh Piano:** Mỗi nốt nhạc tương ứng với giá trị số đang được xử lý.
    *   **Điều khiển:** Tùy chỉnh tốc độ, bật/tắt âm thanh.

Cách chạy Visualizer:
```bash
go run visualizer/main.go
```
Sau đó mở trình duyệt tại: `http://localhost:8080`

---

## 📂 Cấu trúc Dự án

```
.
├── algorithms/          # Mã nguồn các thuật toán
│   ├── bubble_sort/
│   ├── quicksort/
│   └── ...
├── bin/                 # Thư mục chứa file binary sau khi build
├── utils/               # Các hàm tiện ích (sinh dữ liệu ngẫu nhiên)
├── visualizer/          # Mã nguồn Web Visualizer
│   ├── index.html
│   ├── main.go
│   └── README.md
├── go_build.sh          # Script build tự động
├── go.mod               # Go module file
└── README.md            # Tài liệu tổng hợp (File này)
```

---

## 🚀 Bắt đầu ngay

1.  **Clone dự án:**
    ```bash
    git clone https://github.com/yourusername/go_SortingAlgorithms.git
    cd go_SortingAlgorithms
    ```

2.  **Chạy thử Bubble Sort:**
    ```bash
    go run algorithms/bubble_sort/main.go
    ```

3.  **Mở Visualizer:**
    ```bash
    go run visualizer/main.go
    ```

---
*Dự án được thực hiện bởi Tony Cao - 2026*
