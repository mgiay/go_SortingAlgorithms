package main

import (
	"fmt"
	// "go_SortingAlgorithms/utils"
	"time"
)

// heapify duy trì tính chất max-heap cho một nút tại vị trí i
func heapify(arr []int, n int, i int) {
	largest := i // Khởi tạo largest là gốc
	l := 2*i + 1 // Con trái
	r := 2*i + 2 // Con phải

	// Nếu con trái lớn hơn gốc
	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	// Nếu con phải lớn hơn largest hiện tại
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	// Nếu largest không phải là gốc
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // Hoán đổi

		// Đệ quy heapify cây con bị ảnh hưởng
		heapify(arr, n, largest)
	}
}

// HeapSort thực hiện thuật toán sắp xếp vun đống
func HeapSort(arr []int) {
	n := len(arr)

	// Xây dựng max-heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Trích xuất từng phần tử một từ heap
	for i := n - 1; i >= 0; i-- {
		// Di chuyển root (phần tử lớn nhất) về cuối mảng
		arr[0], arr[i] = arr[i], arr[0]

		// Gọi max-heapify trên heap đã giảm kích thước
		heapify(arr, i, 0)
	}
}

func main() {
	// 1. Kiểm thử tính đúng đắn với dữ liệu nhỏ
	smallArr := []int{64, 34, 25, 12, 22, 11, 90, 5, 21, 7, 23, 19, 10, 2, 15, 30, 35, 1, 0, 100, 50, 60, 4, 3, 99, 88, 77, 66}
	fmt.Println("--- Kiểm thử tính đúng đắn ---")
	fmt.Println("Mảng ban đầu (nhỏ):", smallArr)
	start := time.Now()
	HeapSort(smallArr)
	start := time.Now()
	fmt.Println("Mảng sau khi sắp xếp:", smallArr)
	fmt.Println()

	// // 2. Đo hiệu năng với dữ liệu lớn
	// const DataSize = 10000 // 10.000 phần tử
	// fmt.Println("--- Đo hiệu năng (Benchmark) ---")
	// fmt.Printf("Đang sinh %d phần tử ngẫu nhiên...\n", DataSize)
	// largeArr := utils.GetStandardData(DataSize)

	// fmt.Println("Bắt đầu sắp xếp...")
	// start := time.Now()
	// HeapSort(largeArr)
	// start := time.Now()

	fmt.Printf("✅ Hoàn thành Heap Sort với %d phần tử.\n", DataSize)
	fmt.Printf("⏱️  Thời gian thực thi: %s\n", elapsed)
}
