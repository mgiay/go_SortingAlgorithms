package main

import (
	"fmt"
	"go_SortingAlgorithms/utils"
	"time"
)

// getMax trả về giá trị lớn nhất trong mảng
func getMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

// countSort thực hiện sắp xếp đếm (Counting Sort) theo chữ số được chỉ định bởi exp
func countSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	// Khởi tạo mảng đếm
	for i := 0; i < 10; i++ {
		count[i] = 0
	}

	// Đếm số lần xuất hiện của các chữ số
	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	// Tính tổng tiền tố để xác định vị trí thực sự
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Xây dựng mảng output
	// Duyệt ngược để đảm bảo tính ổn định (stable sort)
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	// Sao chép mảng output vào arr
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// RadixSort thực hiện thuật toán sắp xếp cơ số
func RadixSort(arr []int) {
	// Tìm số lớn nhất để biết số lượng chữ số
	m := getMax(arr)

	// Thực hiện Counting Sort cho từng chữ số.
	// exp là 10^i (1, 10, 100, ...)
	for exp := 1; m/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}
}

func main() {
	// 1. Kiểm thử tính đúng đắn với dữ liệu nhỏ
	smallArr := []int{64, 34, 25, 12, 22, 11, 90, 5, 21, 7, 23, 19, 10, 2, 15, 30, 35, 1, 0, 100, 50, 60, 4, 3, 99, 88, 77, 66}
	fmt.Println("--- Kiểm thử tính đúng đắn ---")
	fmt.Println("Mảng ban đầu (nhỏ):", smallArr)
	RadixSort(smallArr)
	fmt.Println("Mảng sau khi sắp xếp:", smallArr)
	fmt.Println()

	// 2. Đo hiệu năng với dữ liệu lớn
	const DataSize = 10000 // 10.000 phần tử
	fmt.Println("--- Đo hiệu năng (Benchmark) ---")
	fmt.Printf("Đang sinh %d phần tử ngẫu nhiên...\n", DataSize)
	largeArr := utils.GetStandardData(DataSize)

	fmt.Println("Bắt đầu sắp xếp...")
	start := time.Now()
	RadixSort(largeArr)
	elapsed := time.Since(start)

	fmt.Printf("✅ Hoàn thành Radix Sort với %d phần tử.\n", DataSize)
	fmt.Printf("⏱️  Thời gian thực thi: %s\n", elapsed)
}
