package main

import (
	"fmt"
	"go_SortingAlgorithms/utils"
	"time"
)

// ShellSort sắp xếp mảng số nguyên theo thuật toán Shell Sort
func ShellSort(arr []int) {
	n := len(arr)

	// Bắt đầu với một khoảng cách (gap) lớn, sau đó giảm dần
	for gap := n / 2; gap > 0; gap /= 2 {
		// Thực hiện Insertion Sort cho từng "mảng con" với khoảng cách gap
		// Các phần tử cách nhau 'gap' vị trí sẽ được so sánh và sắp xếp
		for i := gap; i < n; i++ {
			temp := arr[i]

			// Dịch chuyển các phần tử của mảng con đã sắp xếp lên phía sau
			// cho đến khi tìm được vị trí đúng cho temp
			var j int
			for j = i; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}

			// Chèn temp vào vị trí đúng trong mảng con
			arr[j] = temp
		}
	}
}

func main() {
	// 1. Kiểm thử tính đúng đắn với dữ liệu nhỏ
	smallArr := []int{64, 34, 25, 12, 22, 11, 90, 5, 21, 7, 23, 19, 10, 2, 15, 30, 35, 1, 0, 100, 50, 60, 4, 3, 99, 88, 77, 66}
	fmt.Println("--- Kiểm thử tính đúng đắn ---")
	fmt.Println("Mảng ban đầu (nhỏ):", smallArr)
	ShellSort(smallArr)
	fmt.Println("Mảng sau khi sắp xếp:", smallArr)
	fmt.Println()

	// 2. Đo hiệu năng với dữ liệu lớn
	const DataSize = 10000 // 10.000 phần tử
	fmt.Println("--- Đo hiệu năng (Benchmark) ---")
	fmt.Printf("Đang sinh %d phần tử ngẫu nhiên...\n", DataSize)
	largeArr := utils.GetStandardData(DataSize)

	fmt.Println("Bắt đầu sắp xếp...")
	start := time.Now()
	ShellSort(largeArr)
	elapsed := time.Since(start)

	fmt.Printf("✅ Hoàn thành Shell Sort với %d phần tử.\n", DataSize)
	fmt.Printf("⏱️  Thời gian thực thi: %s\n", elapsed)
}
