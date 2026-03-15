package main

import (
	"fmt"
	"go_SortingAlgorithms/utils"
	"math"
	"time"
)

const RUN = 32

// insertionSort sắp xếp một đoạn con của mảng
func insertionSort(arr []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		temp := arr[i]
		j := i - 1
		for j >= left && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

// merge gộp hai mảng con đã sắp xếp
func merge(arr []int, l, m, r int) {
	len1 := m - l + 1
	len2 := r - m
	left := make([]int, len1)
	right := make([]int, len2)

	for i := 0; i < len1; i++ {
		left[i] = arr[l+i]
	}
	for i := 0; i < len2; i++ {
		right[i] = arr[m+1+i]
	}

	i, j, k := 0, 0, l
	for i < len1 && j < len2 {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len1 {
		arr[k] = left[i]
		k++
		i++
	}
	for j < len2 {
		arr[k] = right[j]
		k++
		j++
	}
}

// TimSort thực hiện thuật toán Timsort
func TimSort(arr []int) {
	n := len(arr)

	// Sắp xếp từng mảng con kích thước RUN bằng Insertion Sort
	for i := 0; i < n; i += RUN {
		end := int(math.Min(float64(i+RUN-1), float64(n-1)))
		insertionSort(arr, i, end)
	}

	// Trộn các RUN lại với nhau
	// size bắt đầu từ RUN, sau đó gấp đôi lên (32, 64, 128...)
	for size := RUN; size < n; size = 2 * size {
		for left := 0; left < n; left += 2 * size {
			mid := left + size - 1
			right := int(math.Min(float64(left+2*size-1), float64(n-1)))

			if mid < right {
				merge(arr, left, mid, right)
			}
		}
	}
}

func main() {
	// 1. Kiểm thử tính đúng đắn với dữ liệu nhỏ
	smallArr := []int{64, 34, 25, 12, 22, 11, 90, 5, 21, 7, 23, 19, 10, 2, 15, 30, 35, 1, 0, 100, 50, 60, 4, 3, 99, 88, 77, 66}
	fmt.Println("--- Kiểm thử tính đúng đắn ---")
	fmt.Println("Mảng ban đầu (nhỏ):", smallArr)
	TimSort(smallArr)
	fmt.Println("Mảng sau khi sắp xếp:", smallArr)
	fmt.Println()

	// 2. Đo hiệu năng với dữ liệu lớn
	const DataSize = 10000 // 10.000 phần tử
	fmt.Println("--- Đo hiệu năng (Benchmark) ---")
	fmt.Printf("Đang sinh %d phần tử ngẫu nhiên...\n", DataSize)
	largeArr := utils.GetStandardData(DataSize)

	fmt.Println("Bắt đầu sắp xếp...")
	start := time.Now()
	TimSort(largeArr)
	elapsed := time.Since(start)

	fmt.Printf("✅ Hoàn thành Tim Sort với %d phần tử.\n", DataSize)
	fmt.Printf("⏱️  Thời gian thực thi: %s\n", elapsed)
}
