package impl

import (
	"encoding/binary"
	"fmt"
)

// Entry はキーと値のペアを表す
type Entry struct {
	Key   interface{}
	Value interface{}
}

// HashMapImplementation はスライスを使用したHashMapの基本実装を提供する
type HashMapImplementation struct {
	buckets [][]Entry // バケットのスライス
	size    int       // 現在の要素数
}

// NewHashMap は新しいHashMapを作成する
func NewHashMap(bucketSize int) *HashMapImplementation {
	return &HashMapImplementation{
		buckets: make([][]Entry, bucketSize),
		size:    0,
	}
}

// hashKey はキーのハッシュ値を計算する
func (h *HashMapImplementation) hashKey(key string) int {
	const seed = 0x9747b28c // 任意のシード値
	const m = 0x5bd1e995
	const r = 24

	data := []byte(key)
	length := len(data)
	hash := uint32(seed ^ length)

	for len(data) >= 4 {
		k := binary.LittleEndian.Uint32(data[:4])
		k *= m
		k ^= k >> r
		k *= m

		hash *= m
		hash ^= k

		data = data[4:]
	}

	switch len(data) {
	case 3:
		hash ^= uint32(data[2]) << 16
		fallthrough
	case 2:
		hash ^= uint32(data[1]) << 8
		fallthrough
	case 1:
		hash ^= uint32(data[0])
		hash *= m
	}

	hash ^= hash >> 13
	hash *= m
	hash ^= hash >> 15

	return int(hash)
}

// Put はキーと値のペアを格納する
func (h *HashMapImplementation) Put(key, value interface{}) {
	hash := h.hashKey(fmt.Sprintf("%v", key))
	index := hash % len(h.buckets)

	// バケット内を探索して既存のキーを更新
	for i, entry := range h.buckets[index] {
		if entry.Key == key {
			h.buckets[index][i].Value = value
			return
		}
	}

	// 新しいエントリを追加
	h.buckets[index] = append(h.buckets[index], Entry{Key: key, Value: value})
	h.size++
}

// Get はキーに対応する値を取得する
func (h *HashMapImplementation) Get(key interface{}) (interface{}, bool) {
	hash := h.hashKey(fmt.Sprintf("%v", key))
	index := hash % len(h.buckets)

	// バケット内を探索
	for _, entry := range h.buckets[index] {
		if entry.Key == key {
			return entry.Value, true
		}
	}

	return nil, false
}

// Remove はキーに対応するエントリを削除する
func (h *HashMapImplementation) Remove(key interface{}) bool {
	hash := h.hashKey(fmt.Sprintf("%v", key))
	index := hash % len(h.buckets)

	// バケット内を探索してエントリを削除
	for i, entry := range h.buckets[index] {
		if entry.Key == key {
			h.buckets[index] = append(h.buckets[index][:i], h.buckets[index][i+1:]...)
			h.size--
			return true
		}
	}

	return false
}

// Size は現在の要素数を取得する
func (h *HashMapImplementation) Size() int {
	return h.size
}

// GetAllEntries は全てのエントリを取得する（テスト用）
func (h *HashMapImplementation) GetAllEntries() map[string]interface{} {
	result := make(map[string]interface{})
	for _, bucket := range h.buckets {
		for _, entry := range bucket {
			result[fmt.Sprintf("%v", entry.Key)] = entry.Value
		}
	}
	return result
}
