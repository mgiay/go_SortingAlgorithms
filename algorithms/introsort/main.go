package main

import (
	"fmt"
	"go_SortingAlgorithms/utils"
	"math"
	"time"
)

// partition phân hoạch mảng
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// heapify cho Heapsort
func heapify(arr []int, n int, i int, offset int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l+offset] > arr[largest+offset] {
		largest = l
	}
	if r < n && arr[r+offset] > arr[largest+offset] {
		largest = r
	}

	if largest != i {
		arr[i+offset], arr[largest+offset] = arr[largest+offset], arr[i+offset]
		heapify(arr, n, largest, offset)
	}
}

// heapSort thực hiện Heapsort trên một đoạn mảng
func heapSort(arr []int, low, high int) {
	n := high - low + 1
	offset := low
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i, offset)
	}
	for i := n - 1; i >= 0; i-- {
		arr[offset], arr[offset+i] = arr[offset+i], arr[offset]
		heapify(arr, i, 0, offset)
	}
}

// insertionSort sắp xếp chèn cho các mảng nhỏ
func insertionSort(arr []int, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1
		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// introsortUtil hàm đệ quy chính của Introsort
func introsortUtil(arr []int, low, high, depthLimit int) {
	size := high - low + 1
	if size < 16 {
		insertionSort(arr, low, high)
		return
	}

	if depthLimit == 0 {
		heapSort(arr, low, high)
		return
	}

	pivot := partition(arr, low, high)
	introsortUtil(arr, low, pivot-1, depthLimit-1)
	introsortUtil(arr, pivot+1, high, depthLimit-1)
}

// IntroSort thực hiện thuật toán Introsort
func IntroSort(arr []int) {
	n := len(arr)
	depthLimit := int(2 * math.Log2(float64(n)))
	introsortUtil(arr, 0, n-1, depthLimit)
}

func main() {
	// 1. Kiểm thử tính đúng đắn với dữ liệu nhỏ
	smallArr := []int{64, 34, 25, 12, 22, 11, 90, 5, 21, 7, 23, 19, 10, 2, 15, 30, 35, 1, 0, 100, 50, 60, 4, 3, 99, 88, 77, 66}
	fmt.Println("--- Kiểm thử tính đúng đắn ---")
	fmt.Println("Mảng ban đầu (nhỏ):", smallArr)
	IntroSort(smallArr)
	fmt.Println("Mảng sau khi sắp xếp:", smallArr)
	fmt.Println()

	// 2. Đo hiệu năng với dữ liệu lớn
	const DataSize = 10000 // 10.000 phần tử
	fmt.Println("--- Đo hiệu năng (Benchmark) ---")
	fmt.Printf("Đang sinh %d phần tử ngẫu nhiên...\n", DataSize)
	largeArr := utils.GetStandardData(DataSize)

	fmt.Println("Bắt đầu sắp xếp...")
	start := time.Now()
	IntroSort(largeArr)
	elapsed := time.Since(start)

	fmt.Printf("✅ Hoàn thành Intro Sort với %d phần tử.\n", DataSize)
	fmt.Printf("⏱️  Thời gian thực thi: %s\n", elapsed)
}
