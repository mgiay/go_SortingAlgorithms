# Danh sách các thuật toán sắp xếp nổi tiếng và bối cảnh ra đời

Tài liệu này tổng hợp các thuật toán sắp xếp (sorting algorithms) nổi tiếng, được sử dụng rộng rãi trong thực tế và lịch sử phát triển khoa học máy tính.

---

## 1. Quicksort (Sắp xếp nhanh)

- **Tác giả:** Sir Tony Hoare (C.A.R. Hoare)
- **Quốc tịch:** Anh (United Kingdom)
- **Năm công bố:** 1959 (phát triển), 1961 (công bố chính thức)
- **Tuổi khi công bố:** Khoảng 27 tuổi (sinh năm 1934)
- **Hoàn cảnh ra đời:**
  Tony Hoare khi đó là một lập trình viên trẻ làm việc tại Moscow (Liên Xô) trong một dự án dịch máy tiếng Nga sang tiếng Anh cho Phòng thí nghiệm Vật lý Quốc gia Anh (NPL). Ông cần sắp xếp các từ trong câu tiếng Nga để tra cứu trong từ điển từ vựng được lưu trữ trên băng từ (magnetic tape). Các thuật toán hiện có quá chậm hoặc không phù hợp với kích thước bộ nhớ hạn chế lúc bấy giờ.
- **Ý nghĩa và tình huống sử dụng:**
  Quicksort là một trong những thuật toán sắp xếp hiệu quả nhất và được sử dụng phổ biến nhất cho đến ngày nay (ví dụ: `qsort` trong C, `std::sort` trong C++ thường dùng biến thể của nó). Nó sử dụng chiến lược "Chia để trị" (Divide and Conquer). Dù trường hợp xấu nhất là O(n^2), nhưng trung bình nó chạy rất nhanh O(n log n) và hoạt động tốt trên bộ nhớ cache của máy tính hiện đại.

## 2. Merge Sort (Sắp xếp trộn)

- **Tác giả:** John von Neumann
- **Quốc tịch:** Mỹ (gốc Hungary)
- **Năm công bố:** 1945
- **Tuổi khi công bố:** 42 tuổi (sinh năm 1903)
- **Hoàn cảnh ra đời:**
  John von Neumann đã viết ra thuật toán này trên một bản thảo dài 23 trang cho máy tính EDVAC. Đây là thời kỳ sơ khai của máy tính điện tử, và việc xử lý dữ liệu lớn đòi hỏi một phương pháp có thể dự đoán được hiệu suất và xử lý tuần tự tốt.
- **Ý nghĩa và tình huống sử dụng:**
  Merge Sort là thuật toán sắp xếp ổn định (stable sort) điển hình với độ phức tạp O(n log n) trong mọi trường hợp. Nó đặc biệt hữu ích khi sắp xếp dữ liệu lớn không thể chứa hết trong RAM (external sorting) vì nó truy cập dữ liệu một cách tuần tự, rất phù hợp với băng từ hoặc ổ đĩa cứng.

## 3. Timsort

- **Tác giả:** Tim Peters
- **Quốc tịch:** Mỹ (United States)
- **Năm công bố:** 2002
- **Tuổi khi công bố:** Khoảng 35-40 tuổi (Không công bố chính thức, ông bắt đầu sự nghiệp kỹ sư phần mềm cuối thập niên 1980)
- **Hoàn cảnh ra đời:**
  Được tạo ra để thay thế thuật toán sắp xếp cũ của ngôn ngữ lập trình Python. Tim Peters nhận thấy rằng dữ liệu trong thực tế thường có các đoạn con đã được sắp xếp sẵn (gọi là "runs"). Ông muốn tận dụng điều này thay vì sắp xếp lại từ đầu như Quicksort hay Merge Sort thuần túy.
- **Ý nghĩa và tình huống sử dụng:**
  Timsort là một thuật toán lai giữa Merge Sort và Insertion Sort. Nó cực kỳ hiệu quả với dữ liệu thực tế (real-world data). Hiện nay, Timsort là thuật toán sắp xếp mặc định trong **Python** (`list.sort`), **Java** (cho Arrays.sort với kiểu Object), và **Android SDK**. Nó rất nhanh, ổn định và tối ưu cho dữ liệu hỗn loạn nhưng có tính thứ tự cục bộ.

## 4. Heapsort (Sắp xếp vun đống)

- **Tác giả:** J. W. J. Williams (John William Joseph Williams)
- **Quốc tịch:** Anh (United Kingdom)
- **Năm công bố:** 1964
- **Tuổi khi công bố:** 34 tuổi (sinh năm 1930)
- **Hoàn cảnh ra đời:**
  Williams phát minh ra cấu trúc dữ liệu Heap (đống) chính là để phục vụ cho thuật toán sắp xếp này khi ông làm việc cho công ty Elliot Brothers (London). Ông muốn cải thiện việc lựa chọn phần tử lớn nhất/nhỏ nhất một cách nhanh chóng mà không cần quét toàn bộ mảng.
- **Ý nghĩa và tình huống sử dụng:**
  Heapsort đảm bảo độ phức tạp O(n log n) ngay cả trong trường hợp xấu nhất (điểm yếu của Quicksort) và không tốn thêm bộ nhớ phụ như Merge Sort (in-place). Tuy nhiên, trên thực tế nó thường chậm hơn Quicksort do tính chất truy cập bộ nhớ không liên tục (kém cache locality). Nó thường được dùng trong các hệ thống nhúng hoặc nơi bộ nhớ cực kỳ hạn chế và cần đảm bảo thời gian chạy tối đa cố định (real-time systems).

## 5. Shellsort

- **Tác giả:** Donald Shell
- **Quốc tịch:** Mỹ (United States)
- **Năm công bố:** 1959
- **Tuổi khi công bố:** 35 tuổi (sinh năm 1924)
- **Hoàn cảnh ra đời:**
  Donald Shell muốn cải tiến thuật toán Insertion Sort (Sắp xếp chèn) khi ông đang làm việc tại General Electric. Insertion Sort rất chậm khi các phần tử nhỏ nằm ở cuối mảng vì chúng phải di chuyển từng bước một về đầu. Shell đưa ra ý tưởng so sánh các phần tử cách xa nhau để đưa chúng về đúng vị trí nhanh hơn.
- **Ý nghĩa và tình huống sử dụng:**
  Đây là thuật toán đầu tiên phá vỡ rào cản độ phức tạp O(n^2). Shellsort rất dễ cài đặt và hiệu quả với các mảng kích thước trung bình. Trong các thư viện C chuẩn cũ (uClibc), nó từng được sử dụng cho `qsort`. Ngày nay nó ít phổ biến hơn Quicksort/Timsort nhưng vẫn hữu dụng trong các ứng dụng nhúng vì code rất ngắn gọn và không dùng đệ quy (tránh tràn stack).

## 6. Radix Sort (Sắp xếp cơ số)

- **Tác giả:** Herman Hollerith
- **Quốc tịch:** Mỹ (United States)
- **Năm công bố:** 1887
- **Tuổi khi công bố:** 27 tuổi (sinh năm 1860)
- **Hoàn cảnh ra đời:**
  Herman Hollerith phát triển máy sắp xếp thẻ đục lỗ (tabulating machine) để phục vụ cho việc thống kê dân số Hoa Kỳ năm 1890. Thuật toán này không so sánh trực tiếp các giá trị mà phân loại dựa trên từng chữ số (hoặc ký tự) của dữ liệu.
- **Ý nghĩa và tình huống sử dụng:**
  Radix Sort có thể đạt độ phức tạp tuyến tính O(nk) trong một số trường hợp, nhanh hơn các thuật toán so sánh O(n log n). Nó được dùng nhiều trong xử lý chuỗi, sắp xếp số nguyên kích thước cố định, hoặc trong các bài toán cần tốc độ cực cao trên phần cứng chuyên dụng (như GPU sorting).

## 7. Introsort (Introspective Sort)

- **Tác giả:** David Musser
- **Quốc tịch:** Mỹ (United States)
- **Năm công bố:** 1997
- **Tuổi khi công bố:** Khoảng 51 tuổi (ước tính sinh năm 1946, nhận bằng PhD năm 1971)
- **Hoàn cảnh ra đời:**
  Musser muốn kết hợp ưu điểm của Quicksort (nhanh trung bình) và Heapsort (an toàn trong trường hợp xấu nhất). Quicksort có thể bị suy biến thành O(n^2) với một số dữ liệu nhất định ("Quicksort killer").
- **Ý nghĩa và tình huống sử dụng:**
  Introsort bắt đầu bằng Quicksort, nhưng nếu độ sâu đệ quy vượt quá một mức giới hạn (cho thấy đang gặp trường hợp xấu), nó chuyển sang Heapsort. Đây là thuật toán mặc định được sử dụng trong thư viện chuẩn **C++ STL** (`std::sort`) và **.NET Framework**. Nó đảm bảo hiệu năng cao và an toàn tuyệt đối.

## 8. Bubble Sort (Sắp xếp nổi bọt)

- **Tác giả:** (Không rõ người đầu tiên, phổ biến từ những năm 1950)
- **Năm công bố:** 1956 (được phân tích bởi Edward Friend)
- **Quốc tịch/Tuổi:** Không xác định cụ thể do nguồn gốc dân gian/cộng đồng.
- **Hoàn cảnh ra đời:**
  Là một trong những thuật toán đơn giản nhất, mô phỏng việc các phần tử "nhẹ" (nhỏ) nổi lên trên giống như bong bóng khí.
- **Ý nghĩa và tình huống sử dụng:**
  Mặc dù hiệu suất rất tệ O(n^2) và hiếm khi được dùng trong sản phẩm thực tế xử lý dữ liệu lớn, Bubble Sort có ý nghĩa lịch sử và giáo dục to lớn. Nó là bài học vỡ lòng về thuật toán cho hầu hết lập trình viên. Trong thực tế, nó chỉ hữu ích khi mảng đã gần như được sắp xếp hoàn toàn hoặc kích thước rất nhỏ, và người viết code cần một giải pháp "mì ăn liền" nhanh gọn nhất.

## 9. Insertion Sort (Sắp xếp chèn)

- **Tác giả:** (Kỹ thuật cơ bản, xuất hiện từ thời sắp xếp thẻ bài thủ công)
- **Quốc tịch/Tuổi:** Không áp dụng.
- **Hoàn cảnh ra đời:**
  Mô phỏng cách con người sắp xếp các lá bài trên tay: rút một lá bài và chèn nó vào đúng vị trí trong các lá đã xếp.
- **Ý nghĩa và tình huống sử dụng:**
  Rất hiệu quả với các mảng kích thước nhỏ (n < 20-50). Do đó, các thuật toán phức tạp như Quicksort hay Merge Sort thường chuyển sang dùng Insertion Sort khi chia nhỏ mảng đến kích thước này để tối ưu hóa tốc độ.

## 10. Smoothsort

- **Tác giả:** Edsger W. Dijkstra
- **Quốc tịch:** Hà Lan (Netherlands)
- **Năm công bố:** 1981
- **Tuổi khi công bố:** 51 tuổi (sinh năm 1930)
- **Hoàn cảnh ra đời:**
  Dijkstra muốn tạo ra một biến thể của Heapsort nhưng tận dụng được lợi thế khi mảng đã được sắp xếp một phần (điều mà Heapsort gốc không làm được).
- **Ý nghĩa và tình huống sử dụng:**
  Smoothsort có độ phức tạp O(n) trong trường hợp tốt nhất (mảng đã sắp xếp) và O(n log n) trong trường hợp xấu nhất. Mặc dù lý thuyết rất hay, nhưng do cấu trúc phức tạp và khó cài đặt, nó ít được sử dụng phổ biến hơn Timsort hay Introsort trong các thư viện chuẩn ngày nay.
