package main

import (
	"fmt"
	"go_SortingAlgorithms/utils"
	"time"
)

// heapify duy trì tính chất max-heap cho một nút tại vị trí i
func heapify(arr []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// SmoothSortSimplified thực hiện sắp xếp dựa trên Heap
// Lưu ý: Đây là phiên bản đơn giản hóa sử dụng Binary Heap thay vì Leonardo Heap
// do tính phức tạp của Smoothsort gốc.
func SmoothSort(arr []int) {
	n := len(arr)

	// Xây dựng max-heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Trích xuất phần tử
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	// 1. Kiểm thử tính đúng đắn với dữ liệu nhỏ
	smallArr := []int{64, 34, 25, 12, 22, 11, 90, 5, 21, 7, 23, 19, 10, 2, 15, 30, 35, 1, 0, 100, 50, 60, 4, 3, 99, 88, 77, 66}
	fmt.Println("--- Kiểm thử tính đúng đắn ---")
	fmt.Println("Mảng ban đầu (nhỏ):", smallArr)
	SmoothSort(smallArr)
	fmt.Println("Mảng sau khi sắp xếp:", smallArr)
	fmt.Println()

	// 2. Đo hiệu năng với dữ liệu lớn
	const DataSize = 10000 // 10.000 phần tử
	fmt.Println("--- Đo hiệu năng (Benchmark) ---")
	fmt.Printf("Đang sinh %d phần tử ngẫu nhiên...\n", DataSize)
	largeArr := utils.GetStandardData(DataSize)
	
	fmt.Println("Bắt đầu sắp xếp...")
	start := time.Now()
	SmoothSort(largeArr)
	elapsed := time.Since(start)

	fmt.Printf("✅ Hoàn thành Smooth Sort (Simplified) với %d phần tử.\n", DataSize)
	fmt.Printf("⏱️  Thời gian thực thi: %s\n", elapsed)
}
