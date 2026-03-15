package main

import (
	"fmt"
	"go_SortingAlgorithms/utils"
	"time"
)

// partition phân hoạch mảng thành hai phần dựa trên pivot
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Chọn pivot là phần tử cuối
	i := low - 1       // Vị trí của phần tử nhỏ hơn

	for j := low; j < high; j++ {
		// Nếu phần tử hiện tại nhỏ hơn hoặc bằng pivot
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Hoán đổi
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// QuickSort thực hiện thuật toán sắp xếp nhanh
func QuickSort(arr []int, low, high int) {
	if low < high {
		// pi là chỉ số phân hoạch (partitioning index)
		pi := partition(arr, low, high)

		// Sắp xếp đệ quy các phần tử trước và sau phân hoạch
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func main() {
	// 1. Kiểm thử tính đúng đắn với dữ liệu nhỏ
	smallArr := []int{64, 34, 25, 12, 22, 11, 90, 5, 21, 7, 23, 19, 10, 2, 15, 30, 35, 1, 0, 100, 50, 60, 4, 3, 99, 88, 77, 66}
	fmt.Println("--- Kiểm thử tính đúng đắn ---")
	fmt.Println("Mảng ban đầu (nhỏ):", smallArr)
	QuickSort(smallArr, 0, len(smallArr)-1)
	fmt.Println("Mảng sau khi sắp xếp:", smallArr)
	fmt.Println()

	// 2. Đo hiệu năng với dữ liệu lớn
	const DataSize = 10000 // 10.000 phần tử
	fmt.Println("--- Đo hiệu năng (Benchmark) ---")
	fmt.Printf("Đang sinh %d phần tử ngẫu nhiên...\n", DataSize)
	largeArr := utils.GetStandardData(DataSize)

	fmt.Println("Bắt đầu sắp xếp...")
	start := time.Now()
	QuickSort(largeArr, 0, len(largeArr)-1)
	elapsed := time.Since(start)

	fmt.Printf("✅ Hoàn thành Quick Sort với %d phần tử.\n", DataSize)
	fmt.Printf("⏱️  Thời gian thực thi: %s\n", elapsed)
}
