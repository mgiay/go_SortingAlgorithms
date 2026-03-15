package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	Port      = ":8080"
	ArraySize = 50
)

type EventData struct {
	Arr        []int  `json:"arr,omitempty"`
	CompareIdx []int  `json:"compareIdx,omitempty"` // Index của các phần tử đang so sánh/thao tác
	Status     string `json:"status,omitempty"`
	Finished   bool   `json:"finished,omitempty"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "visualizer/index.html")
	})

	http.HandleFunc("/events", handleSSE)

	fmt.Printf("🚀 Server đang chạy tại http://localhost%s\n", Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		panic(err)
	}
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set headers cho SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	algo := r.URL.Query().Get("algo")
	if algo == "" {
		algo = "bubble"
	}

	// Lấy tham số delay từ URL, mặc định 100ms
	delayStr := r.URL.Query().Get("delay")
	delayDuration := 100 * time.Millisecond
	if delayStr != "" {
		if d, err := time.ParseDuration(delayStr + "ms"); err == nil {
			delayDuration = d
		}
	}

	// Sinh dữ liệu ngẫu nhiên
	arr := generateRandomArray(ArraySize)

	// Channel để nhận sự kiện từ thuật toán
	eventChan := make(chan EventData)

	// Chạy thuật toán trong goroutine riêng
	go func() {
		defer close(eventChan)

		// Gửi trạng thái ban đầu
		eventChan <- EventData{Arr: makeCopy(arr), Status: "Bắt đầu sắp xếp..."}
		time.Sleep(500 * time.Millisecond)

		switch algo {
		case "bubble":
			bubbleSortVisual(arr, eventChan, delayDuration)
		case "quick":
			quickSortVisual(arr, 0, len(arr)-1, eventChan, delayDuration)
		case "heap":
			heapSortVisual(arr, eventChan, delayDuration)
		case "merge":
			mergeSortVisual(arr, eventChan, delayDuration)
		case "insertion":
			insertionSortVisual(arr, eventChan, delayDuration)
		case "selection":
			selectionSortVisual(arr, eventChan, delayDuration)
		default:
			eventChan <- EventData{Status: "Thuật toán chưa được hỗ trợ: " + algo}
		}

		// Gửi trạng thái kết thúc
		eventChan <- EventData{Arr: makeCopy(arr), Status: "Hoàn tất", Finished: true}
	}()

	// Lắng nghe sự kiện và gửi về client
	for event := range eventChan {
		jsonData, _ := json.Marshal(event)
		fmt.Fprintf(w, "data: %s\n\n", jsonData)
		flusher.Flush()
	}
}

// Helper để copy mảng (tránh race condition khi gửi json)
func makeCopy(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	return dest
}

// --- CÁC THUẬT TOÁN ---

func bubbleSortVisual(arr []int, ch chan EventData, delay time.Duration) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			// Gửi sự kiện so sánh
			ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{j, j + 1}, Status: "So sánh"}
			time.Sleep(delay)

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true

				// Gửi sự kiện hoán đổi
				ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{j, j + 1}, Status: "Hoán đổi"}
				time.Sleep(delay)
			}
		}
		if !swapped {
			break
		}
	}
}

func quickSortVisual(arr []int, low, high int, ch chan EventData, delay time.Duration) {
	if low < high {
		pi := partitionVisual(arr, low, high, ch, delay)
		quickSortVisual(arr, low, pi-1, ch, delay)
		quickSortVisual(arr, pi+1, high, ch, delay)
	}
}

func partitionVisual(arr []int, low, high int, ch chan EventData, delay time.Duration) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{j, high}, Status: "So sánh với Pivot"}
		time.Sleep(delay)

		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{i, j}, Status: "Swap (Partition)"}
			time.Sleep(delay)
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{i + 1, high}, Status: "Đặt Pivot về đúng vị trí"}
	time.Sleep(delay)

	return i + 1
}

func heapSortVisual(arr []int, ch chan EventData, delay time.Duration) {
	n := len(arr)

	// Xây dựng heap
	for i := n/2 - 1; i >= 0; i-- {
		heapifyVisual(arr, n, i, ch, delay)
	}

	// Trích xuất phần tử
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{0, i}, Status: "Di chuyển Max về cuối"}
		time.Sleep(delay)

		heapifyVisual(arr, i, 0, ch, delay)
	}
}

func heapifyVisual(arr []int, n int, i int, ch chan EventData, delay time.Duration) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n {
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{l, largest}, Status: "So sánh Heap con trái"}
		time.Sleep(delay)
		if arr[l] > arr[largest] {
			largest = l
		}
	}

	if r < n {
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{r, largest}, Status: "So sánh Heap con phải"}
		time.Sleep(delay)
		if arr[r] > arr[largest] {
			largest = r
		}
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{i, largest}, Status: "Heapify Swap"}
		time.Sleep(delay)
		heapifyVisual(arr, n, largest, ch, delay)
	}
}

func insertionSortVisual(arr []int, ch chan EventData, delay time.Duration) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{i}, Status: "Chọn phần tử chèn"}
		time.Sleep(delay)

		for j >= 0 && arr[j] > key {
			ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{j, j + 1}, Status: "Dịch chuyển"}
			time.Sleep(delay)

			arr[j+1] = arr[j]
			j = j - 1

			ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{j + 1}, Status: "Dịch xong"}
			time.Sleep(delay)
		}
		arr[j+1] = key
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{j + 1}, Status: "Chèn vào vị trí"}
		time.Sleep(delay)
	}
}

func selectionSortVisual(arr []int, ch chan EventData, delay time.Duration) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{i}, Status: "Tìm số nhỏ nhất..."}
		time.Sleep(delay)

		for j := i + 1; j < n; j++ {
			// So sánh tìm min
			ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{minIdx, j}, Status: "So sánh"}
			time.Sleep(delay)

			if arr[j] < arr[minIdx] {
				minIdx = j
				ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{minIdx}, Status: "Tìm thấy số nhỏ mới"}
				time.Sleep(delay)
			}
		}

		// Hoán đổi nếu tìm thấy số nhỏ hơn
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
			ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{i, minIdx}, Status: "Hoán đổi Min về đầu"}
			time.Sleep(delay)
		}
	}
}

func mergeSortVisual(arr []int, ch chan EventData, delay time.Duration) {
	// Merge Sort in-place hoặc cần visualizer phức tạp hơn để hiển thị mảng phụ.
	// Để đơn giản cho demo, ta dùng một phiên bản đệ quy có copy data trở lại mảng chính để visualizer thấy.
	mergeSortRecursive(arr, 0, len(arr)-1, ch, delay)
}

func mergeSortRecursive(arr []int, l, r int, ch chan EventData, delay time.Duration) {
	if l < r {
		m := l + (r-l)/2
		mergeSortRecursive(arr, l, m, ch, delay)
		mergeSortRecursive(arr, m+1, r, ch, delay)
		mergeVisual(arr, l, m, r, ch, delay)
	}
}

func mergeVisual(arr []int, l, m, r int, ch chan EventData, delay time.Duration) {
	n1 := m - l + 1
	n2 := r - m
	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	i, j, k := 0, 0, l

	for i < n1 && j < n2 {
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{k}, Status: "Merge: So sánh và Gộp"}
		time.Sleep(delay)

		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{k - 1}, Status: "Merge: Sao chép phần dư L"}
		time.Sleep(delay)
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
		ch <- EventData{Arr: makeCopy(arr), CompareIdx: []int{k - 1}, Status: "Merge: Sao chép phần dư R"}
		time.Sleep(delay)
	}
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100) + 1
	}
	return arr
}
