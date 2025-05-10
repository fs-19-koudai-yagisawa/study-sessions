package impl

// 標準ライブラリとの比較検証用
import "sort"

// StandardSort は標準ライブラリのsortパッケージを使用する実装
type StandardSort struct{}

func (s *StandardSort) Sort(data interface{}) {
	switch v := data.(type) {
	case []int:
		sort.Ints(v)
	case []int64:
		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
	case []float64:
		sort.Float64s(v)
	case []string:
		sort.Strings(v)
	}
}
