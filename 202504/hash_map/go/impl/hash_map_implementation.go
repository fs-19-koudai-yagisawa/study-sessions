package impl

import (
	"fmt"
)

// デフォルト値
const (
	DefaultBucketSize = 128 // 大量データに対応するために増加
	DefaultLoadFactor = 0.75
)

// entry はキーと値のペアを表す
type entry struct {
	key   string
	value interface{}
}

// bucket はチェイニング方式で衝突を解決するためのバケット
type bucket struct {
	entries []*entry
}

// HashMapImplementation はハッシュマップの基本実装を提供する
type HashMapImplementation struct {
	buckets    []*bucket // バケット配列
	size       int       // 現在の要素数
	bucketSize int       // バケット配列のサイズ
	loadFactor float64   // リサイズのトリガーとなる負荷係数
}

// NewHashMap は新しいHashMapを作成する
func NewHashMap(bucketSize int) *HashMapImplementation {
	// バケットサイズが0以下の場合はデフォルト値を使用
	if bucketSize <= 0 {
		bucketSize = DefaultBucketSize
	}

	// バケット配列を初期化
	buckets := make([]*bucket, bucketSize)
	for i := 0; i < bucketSize; i++ {
		buckets[i] = &bucket{entries: make([]*entry, 0)}
	}

	return &HashMapImplementation{
		buckets:    buckets,
		size:       0,
		bucketSize: bucketSize,
		loadFactor: DefaultLoadFactor,
	}
}

// hashKey はキーのハッシュ値を計算する
func (h *HashMapImplementation) hashKey(key string) int {
	// 多項式ハッシュ関数を実装
	// 素数を使用してより良い分散を得る
	const prime = 31
	hash := 0

	// 各文字のASCII値を使用してハッシュ値を計算
	for i := 0; i < len(key); i++ {
		hash = (hash*prime + int(key[i])) % h.bucketSize
	}

	return hash
}

// Put はキーと値のペアを格納する
func (h *HashMapImplementation) Put(key, value interface{}) {
	// キーを文字列に変換
	strKey := fmt.Sprintf("%v", key)

	// 負荷係数をチェックし、必要に応じてリサイズ
	if float64(h.size)/float64(h.bucketSize) >= h.loadFactor {
		h.resize()
	}

	// キーのハッシュ値を計算
	index := h.hashKey(strKey)

	// バケット内でキーを検索
	for i, e := range h.buckets[index].entries {
		if e.key == strKey {
			// 既存のキーの場合は値を更新
			h.buckets[index].entries[i].value = value
			return
		}
	}

	// 新しいキーの場合はエントリを追加
	h.buckets[index].entries = append(h.buckets[index].entries, &entry{key: strKey, value: value})
	h.size++
}

// Get はキーに対応する値を取得する
func (h *HashMapImplementation) Get(key interface{}) (interface{}, bool) {
	// キーを文字列に変換
	strKey := fmt.Sprintf("%v", key)

	// キーのハッシュ値を計算
	index := h.hashKey(strKey)

	// バケット内でキーを検索
	for _, e := range h.buckets[index].entries {
		if e.key == strKey {
			return e.value, true
		}
	}

	// キーが見つからない場合
	return nil, false
}

// Remove はキーに対応するエントリを削除する
func (h *HashMapImplementation) Remove(key interface{}) bool {
	// キーを文字列に変換
	strKey := fmt.Sprintf("%v", key)

	// キーのハッシュ値を計算
	index := h.hashKey(strKey)

	// バケット内でキーを検索
	entries := h.buckets[index].entries
	for i, e := range entries {
		if e.key == strKey {
			// エントリを削除（スライスから要素を削除）
			h.buckets[index].entries = append(entries[:i], entries[i+1:]...)
			h.size--
			return true
		}
	}

	// キーが見つからない場合
	return false
}

// resize はバケットサイズを拡張する
func (h *HashMapImplementation) resize() {
	// 古いバケットを保存
	oldBuckets := h.buckets

	// バケットサイズを2倍に拡張
	h.bucketSize *= 2
	h.buckets = make([]*bucket, h.bucketSize)
	for i := 0; i < h.bucketSize; i++ {
		h.buckets[i] = &bucket{entries: make([]*entry, 0)}
	}

	// サイズをリセット（Put操作で再カウント）
	oldSize := h.size
	h.size = 0

	// 古いバケットからすべてのエントリを新しいバケットに再配置
	for _, bucket := range oldBuckets {
		for _, e := range bucket.entries {
			h.Put(e.key, e.value)
		}
	}

	// サイズを復元（Put操作で重複カウントを防ぐため）
	h.size = oldSize
}

// Size は現在の要素数を取得する
func (h *HashMapImplementation) Size() int {
	return h.size
}

// GetAllEntries は全てのエントリを取得する（テスト用）
func (h *HashMapImplementation) GetAllEntries() map[string]interface{} {
	result := make(map[string]interface{})

	// すべてのバケットを走査
	for _, bucket := range h.buckets {
		for _, e := range bucket.entries {
			result[e.key] = e.value
		}
	}

	return result
}
