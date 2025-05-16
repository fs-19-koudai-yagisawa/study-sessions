package impl

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	utils "study-session/utils/go"
)

// parseValue parses a string into an appropriate type (int, float, or string)
func parseValue(val string) interface{} {
	// Try parsing as integer first
	if intVal, err := strconv.Atoi(strings.TrimSpace(val)); err == nil {
		return intVal
	}
	
	// If not an integer, try parsing as float
	if floatVal, err := strconv.ParseFloat(strings.TrimSpace(val), 64); err == nil {
		return floatVal
	}
	
	// If neither integer nor float, return as string
	return strings.TrimSpace(val)
}

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
		for _, itemStr := range strings.Split(inputStr, ",") {
			array = append(array, parseValue(itemStr))
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
		for _, itemStr := range strings.Split(expectedStr, ",") {
			expectedOutput = append(expectedOutput, parseValue(itemStr))
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
			// 配列のコピーを作成
			arrayCopy := make([]interface{}, len(array))
			copy(arrayCopy, array)

			sorted = sorter.Sort(arrayCopy)
			if iterations == 1 {
				// ソート前とソート後の最初の5要素を表示
				fmt.Printf("ソート前の先頭5要素: ")
				for j := 0; j < 5 && j < len(array); j++ {
					fmt.Printf("%v ", array[j])
				}
				fmt.Println()

				fmt.Printf("ソート後の先頭5要素: ")
				for j := 0; j < 5 && j < len(sorted); j++ {
					fmt.Printf("%v ", sorted[j])
				}
				fmt.Println()
			}
		}
	})

	// 正当性検証
	valid := utils.VerifyResult("Sort", sorted, expectedOutput)
	results["valid"] = valid

	return results
}
