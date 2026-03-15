package main

import (
	"fmt"
	"go_SortingAlgorithms/utils"
	"time"
)

// BubbleSort sắp xếp mảng số nguyên theo thuật toán Bubble Sort
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortVisual sắp xếp và in ra quá trình thực hiện
func BubbleSortVisual(arr []int) {
	n := len(arr)
	fmt.Printf("Ban đầu: %v\n", arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		fmt.Printf("\n--- Vòng lặp thứ %d ---\n", i+1)

		for j := 0; j < n-i-1; j++ {
			// In ra so sánh
			// fmt.Printf("So sánh %d và %d: ", arr[j], arr[j+1])
			if arr[j] > arr[j+1] {
				// Hoán đổi
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
				fmt.Printf("👉 Hoán đổi %d <-> %d: %v\n", arr[j+1], arr[j], arr)
			} else {
				// fmt.Printf("Giữ nguyên\n")
			}
		}

		fmt.Printf("📍 Kết thúc vòng %d: %v\n", i+1, arr)

		if !swapped {
			fmt.Println("🏁 Không có hoán đổi nào trong vòng này -> Mảng đã sắp xếp xong.")
			break
		}
	}
}

func main() {
	// 1. Kiểm thử tính đúng đắn với dữ liệu nhỏ
	smallArr := []int{64, 34, 25, 12, 22, 11, 90} // Sử dụng mảng nhỏ hơn để demo dễ nhìn
	fmt.Println("--- Demo quá trình sắp xếp (Bubble Sort Visual) ---")
	BubbleSortVisual(smallArr)
	fmt.Println()

	// 2. Đo hiệu năng với dữ liệu lớn
	const DataSize = 10000 // 10.000 phần tử
	fmt.Println("--- Đo hiệu năng (Benchmark) ---")
	fmt.Printf("Đang sinh %d phần tử ngẫu nhiên...\n", DataSize)
	largeArr := utils.GetStandardData(DataSize)

	fmt.Println("Bắt đầu sắp xếp...")
	start := time.Now()
	BubbleSort(largeArr)
	elapsed := time.Since(start)

	fmt.Printf("✅ Hoàn thành Bubble Sort với %d phần tử.\n", DataSize)
	fmt.Printf("⏱️  Thời gian thực thi: %s\n", elapsed)
}
