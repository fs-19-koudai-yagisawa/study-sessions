package impl

import (
	"study-session/sort/go/fastpath/go"
)

// SortImplementation はソートアルゴリズムの基本実装を提供する
type SortImplementation struct{}

// Sort はデータをソートする
func (s *SortImplementation) Sort(data interface{}) interface{} {
	// 直接型チェックを試みる
	if intSlice, ok := data.([]int); ok {
		// 整数配列を直接ソート
		// int64に変換
		int64Slice := make([]int64, len(intSlice))
		for i, v := range intSlice {
			int64Slice[i] = int64(v)
		}
		fastpath.RadixSortInt64(int64Slice)
		// 結果をintに戻す
		for i, v := range int64Slice {
			intSlice[i] = int(v)
		}
		return intSlice
	}

	if floatSlice, ok := data.([]float64); ok {
		// float64配列を直接ソート
		fastpath.RadixSortFloat(floatSlice)
		return floatSlice
	}

	if strSlice, ok := data.([]string); ok {
		// 文字列配列を直接ソート
		fastpath.SortStringFast(strSlice)
		return strSlice
	}

	// interface{}配列の場合
	if v, ok := data.([]interface{}); ok {
		if len(v) == 0 {
			return v
		}

		// 最初の要素の型に基づいて処理
		switch v[0].(type) {
		case int:
			// 整数配列にコピー
			int64Slice := make([]int64, len(v))
			for i, val := range v {
				int64Slice[i] = int64(val.(int))
			}
			// 整数配列をソート
			fastpath.RadixSortInt64(int64Slice)
			// interface{}配列に戻す
			result := make([]interface{}, len(v))
			for i, val := range int64Slice {
				result[i] = int(val)
			}
			return result

		case float64:
			// float64配列にコピー
			floatSlice := make([]float64, len(v))
			for i, val := range v {
				floatSlice[i] = val.(float64)
			}
			// float64配列をソート
			fastpath.RadixSortFloat(floatSlice)
			// interface{}配列に戻す
			result := make([]interface{}, len(v))
			for i, val := range floatSlice {
				result[i] = val
			}
			return result

		case string:
			// 文字列配列にコピー
			strSlice := make([]string, len(v))
			for i, val := range v {
				strSlice[i] = val.(string)
			}
			// 文字列配列をソート
			fastpath.SortStringFast(strSlice)
			// interface{}配列に戻す
			result := make([]interface{}, len(v))
			for i, val := range strSlice {
				result[i] = val
			}
			return result

		default:
			return v
		}
	}
	return data
}
