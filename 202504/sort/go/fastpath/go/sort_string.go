package fastpath

import (
	"runtime"
	"sync"
)

// insertionSort は指定された配列を挿入ソートでソートする
func insertionSort(data []string) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// SortStringFast は文字列スライスを並列3-way QuickSortでソートする
func SortStringFast(data []string) {
	if len(data) < 2 {
		return
	}

	// 並列処理の数を決定
	numCPU := runtime.NumCPU()
	threshold := len(data) / numCPU
	if threshold < 1000 {
		// 小さすぎる配列は並列化しない
		quickSort3Way(data, 0, len(data)-1)
		return
	}

	// 並列処理用のWaitGroup
	var wg sync.WaitGroup

	// 最初の分割を実行
	pivot := data[len(data)/2]
	lt, gt := partition3Way(data, pivot)

	// 3つの部分を並列にソート
	if lt > 0 {
		wg.Add(1)
		go func(left, right int) {
			defer wg.Done()
			quickSort3Way(data, left, right)
		}(0, lt-1)
	}

	if gt-lt > 0 {
		wg.Add(1)
		go func(left, right int) {
			defer wg.Done()
			quickSort3Way(data, left, right)
		}(lt, gt-1)
	}

	if len(data)-gt > 0 {
		wg.Add(1)
		go func(left, right int) {
			defer wg.Done()
			quickSort3Way(data, left, right)
		}(gt, len(data)-1)
	}

	wg.Wait()
}

// quickSort3Way は3-way QuickSortを実行
func quickSort3Way(data []string, low, high int) {
	if high <= low {
		return
	}

	// 分割サイズが小さい場合は挿入ソートを使用
	if high-low <= 10 {
		insertionSort(data[low : high+1])
		return
	}

	// ピボットを選択（中央値）
	mid := (low + high) / 2
	// 3つの要素の中央値をピボットとして選択
	if data[high] < data[low] {
		data[low], data[high] = data[high], data[low]
	}
	if data[mid] < data[low] {
		data[low], data[mid] = data[mid], data[low]
	}
	if data[high] < data[mid] {
		data[mid], data[high] = data[high], data[mid]
	}
	pivot := data[mid]

	// 分割を行う範囲を限定
	lt, gt := partition3Way(data[low:high+1], pivot)
	lt += low  // lowのオフセットを加算
	gt += low  // lowのオフセットを加算

	// 再帰的にソート
	if lt-1 > low {
		quickSort3Way(data, low, lt-1)  // ピボットより小さい部分
	}
	if gt < high {
		quickSort3Way(data, gt, high)   // ピボットより大きい部分
	}
}

// partition3Way は配列を3つの部分に分割
// 戻り値は (lt, gt) で:
// - data[0:lt] はピボットより小さい
// - data[lt:gt] はピボットと等しい
// - data[gt:] はピボットより大きい
func partition3Way(data []string, pivot string) (int, int) {
	lt, i, gt := 0, 0, len(data)

	// 配列が空の場合は早期リターン
	if len(data) == 0 {
		return 0, 0
	}

	for i < gt {
		if data[i] < pivot {
			data[lt], data[i] = data[i], data[lt]
			lt++
			i++
		} else if data[i] > pivot {
			gt--
			data[i], data[gt] = data[gt], data[i]
		} else {
			i++
		}
	}
	return lt, gt
}
