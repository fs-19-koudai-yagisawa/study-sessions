package impl

import (
	"fmt"
	"sort"
)

// SortImplementation はソートアルゴリズムの基本実装を提供する
type SortImplementation struct{}

// SortSliceOrdered は「比較演算子 < が使える型」だけを対象にソートする
func (s *SortImplementation) Sort(data []interface{}) []interface{} {
	// 元スライスをコピー
	newArr := make([]interface{}, len(data))
	copy(newArr, data)
	
	// 空の配列を処理
	if len(newArr) == 0 {
		return newArr
	}

	// 要素の型を検出して適切な比較関数を使用
	sort.Slice(newArr, func(i, j int) bool {
		// 比較対象の型を取得
		valueI := newArr[i]
		valueJ := newArr[j]
		
		// 両方が整数型の場合
		iInt, iIsInt := valueI.(int)
		jInt, jIsInt := valueJ.(int)
		if iIsInt && jIsInt {
			return iInt < jInt
		}
		
		// 両方が浮動小数点型の場合
		iFloat, iIsFloat := valueI.(float64)
		jFloat, jIsFloat := valueJ.(float64)
		if iIsFloat && jIsFloat {
			return iFloat < jFloat
		}
		
		// 整数と浮動小数点の混在比較
		if iIsInt && jIsFloat {
			return float64(iInt) < jFloat
		}
		if iIsFloat && jIsInt {
			return iFloat < float64(jInt)
		}
		
		// 両方が文字列型の場合
		iStr, iIsStr := valueI.(string)
		jStr, jIsStr := valueJ.(string)
		if iIsStr && jIsStr {
			return iStr < jStr
		}
		
		// 異なる型の場合は型名で比較（一貫性を保つため）
		return fmt.Sprintf("%T", valueI) < fmt.Sprintf("%T", valueJ)
	})
	
	return newArr
}
