package utils

import (
	"math/rand"
)

// GetStandardData trả về một mảng số nguyên ngẫu nhiên với kích thước size.
// Sử dụng seed cố định để đảm bảo dữ liệu giống nhau giữa các lần chạy và giữa các thuật toán khác nhau.
func GetStandardData(size int) []int {
	// Sử dụng seed cố định (ví dụ: 42)
	r := rand.New(rand.NewSource(42))
	
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = r.Intn(size * 10) // Random trong khoảng [0, size*10)
	}
	return arr
}
