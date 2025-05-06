package impl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	utils "study-session/utils/go"
)

// Operation はHashMapに対する操作を表す
type Operation struct {
	Action string      `json:"action"`
	Key    interface{} `json:"key"`
	Value  interface{} `json:"value,omitempty"`
	Debug  bool        `json:"debug,omitempty"`
}

// loadHashMapTestData は入力ファイルと期待値ファイルを読み込む
func loadHashMapTestData(fileDir string) ([]Operation, map[string]interface{}, error) {
	// 入力データの読み込み
	inputData, err := ioutil.ReadFile(strings.Join([]string{fileDir, "input.txt"}, "/"))
	if err != nil {
		return nil, nil, fmt.Errorf("入力ファイルの読み込みに失敗しました: %v", err)
	}

	// 入力データのパース
	var operations []Operation
	err = json.Unmarshal(inputData, &operations)
	if err != nil {
		return nil, nil, fmt.Errorf("入力データのJSONパースに失敗しました: %v", err)
	}

	// 期待値の読み込み
	expectedData, err := ioutil.ReadFile(strings.Join([]string{fileDir, "expected.txt"}, "/"))
	if err != nil {
		return operations, nil, fmt.Errorf("期待値ファイルの読み込みに失敗しました: %v", err)
	}

	// 期待値のパース
	var expectedOutput map[string]interface{}
	err = json.Unmarshal(expectedData, &expectedOutput)
	if err != nil {
		return operations, nil, fmt.Errorf("期待値データのJSONパースに失敗しました: %v", err)
	}

	return operations, expectedOutput, nil
}

// 操作数に基づいて最適なバケット数を計算する
func calculateOptimalBucketSize(operationCount int) int {
	// 操作数が少ない場合はデフォルトのバケット数を使用
	if operationCount < 100 {
		return DefaultBucketSize
	}

	// 操作数に基づいてバケット数を計算
	// 一般的に、予想される要素数の1.3倍程度が効率的
	// 実際の要素数は操作数より少ない可能性があるので、操作数の約10分の7を使用
	estimatedSize := int(float64(operationCount) * 0.7)

	// 2のべき乗に丸める（ハッシュマップのサイズは2のべき乗が効率的）
	bucketSize := 16 // 最小値
	for bucketSize < estimatedSize {
		bucketSize *= 2
	}

	return bucketSize
}

// MeasureHashMapPerformance はHashMapの性能と正当性を計測する
func MeasureHashMapPerformance(fileDir string, iterations int) map[string]interface{} {
	var err error
	operations, expectedOutput, err := loadHashMapTestData(fileDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 操作数に基づいて最適なバケット数を計算
	optimalBucketSize := calculateOptimalBucketSize(len(operations))
	hashMap := NewHashMap(optimalBucketSize)

	fmt.Printf("HashMap実装のパフォーマンス計測と正当性検証:\n")
	fmt.Printf("操作数: %d\n", len(operations))
	fmt.Printf("繰り返し回数: %d\n", iterations)

	// 処理時間とメモリ使用量を計測
	results := utils.MeasurePerformance("HashMap", func() {
		for i := 0; i < iterations; i++ {
			// 複数回反復する場合は新しいインスタンスで開始
			if i > 0 {
				hashMap = NewHashMap(optimalBucketSize)
			}

			for _, op := range operations {
				switch op.Action {
				case "put":
					hashMap.Put(op.Key, op.Value)
				case "get":
					value, exists := hashMap.Get(op.Key)
					if iterations == 1 && op.Debug {
						fmt.Printf("取得: %s => %s (存在: %v)\n", op.Key, value, exists)
					}
				case "remove":
					hashMap.Remove(op.Key)
				}
			}
		}
	})

	// 正当性検証
	actualEntries := hashMap.GetAllEntries()
	valid := utils.VerifyResult("HashMap", actualEntries, expectedOutput)
	results["valid"] = valid

	return results
}
