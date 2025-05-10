package impl

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	utils "study-session/utils/go"
)

// loadSortTestData は入力ファイルと期待値ファイルを読み込む
func loadSortTestData(fileDir string) ([]interface{}, []interface{}, error) {
	// 入力データの読み込み
	inputData, err := ioutil.ReadFile(strings.Join([]string{fileDir, "input.txt"}, "/"))
	if err != nil {
		return nil, nil, fmt.Errorf("入力ファイルの読み込みに失敗しました: %v", err)
	}

	// 入力データのパースと配列への変換
	var array []interface{}
	inputStr := strings.TrimSpace(string(inputData))
	inputStr = strings.Trim(inputStr, "[]")
	if inputStr != "" {
		for _, str := range strings.Split(inputStr, ",") {
			str = strings.TrimSpace(str)
			str = strings.Trim(str, "\"")
			// 数値として解釈を試みる
			if num, err := strconv.ParseInt(str, 10, 64); err == nil {
				// 整数として扱う
				array = append(array, int(num))
			} else if num, err := strconv.ParseFloat(str, 64); err == nil {
				// 浮動小数点数として扱う
				array = append(array, num)
			} else {
				// 数値でない場合は文字列として扱う
				array = append(array, str)
			}
		}
	}

	// 期待値の読み込み
	expectedData, err := ioutil.ReadFile(strings.Join([]string{fileDir, "expected.txt"}, "/"))
	if err != nil {
		return array, nil, fmt.Errorf("期待値ファイルの読み込みに失敗しました: %v", err)
	}

	// 期待値のパースと配列への変換
	var expectedOutput []interface{}
	expectedStr := strings.TrimSpace(string(expectedData))
	expectedStr = strings.Trim(expectedStr, "[]")
	if expectedStr != "" {
		for _, str := range strings.Split(expectedStr, ",") {
			str = strings.TrimSpace(str)
			str = strings.Trim(str, "\"")
			// 数値として解釈を試みる
			if num, err := strconv.ParseInt(str, 10, 64); err == nil {
				// 整数として扱う
				expectedOutput = append(expectedOutput, int(num))
			} else if num, err := strconv.ParseFloat(str, 64); err == nil {
				// 浮動小数点数として扱う
				expectedOutput = append(expectedOutput, num)
			} else {
				// 数値でない場合は文字列として扱う
				expectedOutput = append(expectedOutput, str)
			}
		}
	}

	return array, expectedOutput, nil
}

// MeasureSortPerformance はSortの性能と正当性を計測する
func MeasureSortPerformance(fileDir string, iterations int) map[string]interface{} {
	var err error
	array, expectedOutput, err := loadSortTestData(fileDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	sorter := &SortImplementation{}

	fmt.Printf("Sort実装のパフォーマンス計測と正当性検証:\n")
	fmt.Printf("配列サイズ: %d\n", len(array))
	fmt.Printf("繰り返し回数: %d\n", iterations)

	var sorted []interface{}

	// 処理時間とメモリ使用量を計測
	results := utils.MeasurePerformance("Sort", func() {
		for i := 0; i < iterations; i++ {
			// デバッグ出力
			if iterations == 1 && len(array) > 0 {
				fmt.Printf("最初の要素の型: %T\n", array[0])
			}

			// ソート前の表示
			if iterations == 1 {
				fmt.Printf("ソート前の先頭5要素: ")
				for j := 0; j < 5 && j < len(array); j++ {
					fmt.Printf("%v ", array[j])
				}
				fmt.Println()
			}

			// 配列のコピーを作成
			arrayCopy := make([]interface{}, len(array))
			copy(arrayCopy, array)

			// ソートを実行
			result := sorter.Sort(arrayCopy)

			// ソート後の表示
			if iterations == 1 {
				fmt.Printf("ソート後の先頭5要素: ")
				switch v := result.(type) {
				case []int:
					for j := 0; j < 5 && j < len(v); j++ {
						fmt.Printf("%v ", v[j])
					}
				case []float64:
					for j := 0; j < 5 && j < len(v); j++ {
						fmt.Printf("%v ", v[j])
					}
				case []string:
					for j := 0; j < 5 && j < len(v); j++ {
						fmt.Printf("%v ", v[j])
					}
				case []interface{}:
					for j := 0; j < 5 && j < len(v); j++ {
						fmt.Printf("%v ", v[j])
					}
				}
				fmt.Println()
			}

			// 結果を保存
			if iface, ok := result.([]interface{}); ok {
				sorted = iface
			} else {
				// 型固有の配列をinterface{}配列に変換
				switch v := result.(type) {
				case []int:
					sorted = make([]interface{}, len(v))
					for i, val := range v {
						sorted[i] = val
					}
				case []float64:
					sorted = make([]interface{}, len(v))
					for i, val := range v {
						sorted[i] = val
					}
				case []string:
					sorted = make([]interface{}, len(v))
					for i, val := range v {
						sorted[i] = val
					}
				}
			}
		}
	})

	// 正当性検証
	valid := utils.VerifyResult("Sort", sorted, expectedOutput)
	results["valid"] = valid

	return results
}
