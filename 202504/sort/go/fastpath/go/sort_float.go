package fastpath

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

void radix_sort_uint64(uint64_t* arr, size_t len) {
    if (len < 2) return;

    uint64_t* temp = (uint64_t*)malloc(len * sizeof(uint64_t));
    uint32_t* counts = (uint32_t*)calloc(256, sizeof(uint32_t));
    uint32_t* pos = (uint32_t*)malloc(256 * sizeof(uint32_t));

    // 8バイト（64ビット）を8ビットずつ8回処理
    for (int shift = 0; shift < 64; shift += 8) {
        // カウントをリセット
        memset(counts, 0, 256 * sizeof(uint32_t));

        // 各桁の出現回数をカウント
        for (size_t i = 0; i < len; i++) {
            uint32_t digit = (arr[i] >> shift) & 0xFF;
            counts[digit]++;
        }

        // 累積和を計算して位置を決定
        pos[0] = 0;
        for (int i = 1; i < 256; i++) {
            pos[i] = pos[i-1] + counts[i-1];
        }

        // 要素を一時配列に移動
        for (size_t i = 0; i < len; i++) {
            uint32_t digit = (arr[i] >> shift) & 0xFF;
            temp[pos[digit]++] = arr[i];
        }

        // 結果を元の配列にコピー
        memcpy(arr, temp, len * sizeof(uint64_t));
    }

    free(temp);
    free(counts);
    free(pos);
}

void radix_sort_float64(uint64_t* arr, size_t len) {
    if (len < 2) return;

    uint64_t* output = (uint64_t*)malloc(len * sizeof(uint64_t));
    if (output == NULL) return;

    uint64_t* count = (uint64_t*)malloc(256 * sizeof(uint64_t));
    if (count == NULL) {
        free(output);
        return;
    }

    // 各バイトでソート
    for (int shift = 0; shift < 64; shift += 8) {
        memset(count, 0, 256 * sizeof(uint64_t));

        // カウント
        for (size_t i = 0; i < len; i++) {
            uint64_t byte = (arr[i] >> shift) & 0xFF;
            count[byte]++;
        }

        // 累積和
        for (int i = 1; i < 256; i++) {
            count[i] += count[i-1];
        }

        // 出力配列の構築
        for (int i = len-1; i >= 0; i--) {
            uint64_t byte = (arr[i] >> shift) & 0xFF;
            output[count[byte]-1] = arr[i];
            count[byte]--;
        }

        // 元の配列にコピー
        memcpy(arr, output, len * sizeof(uint64_t));
    }

    free(output);
    free(count);
}

#cgo CFLAGS: -O3
*/
import "C"
import (
	"math"
	"reflect"
	"unsafe"
)

// SortFloatFast はfloat64型スライスをRadix Sortでソートする
func SortFloatFast(data []float64) {
	if len(data) < 2 {
		return
	}

	// float64をuint64のビット表現に変換
	bits := make([]uint64, len(data))
	for i, v := range data {
		bits[i] = math.Float64bits(v)
		// IEEE754の符号ビットを反転して正しい順序になるように調整
		if bits[i]>>63 != 0 {
			bits[i] = ^bits[i]
		}
	}

	// スライスヘッダから直接メモリを参照
	header := (*reflect.SliceHeader)(unsafe.Pointer(&bits))
	ptr := (*C.uint64_t)(unsafe.Pointer(header.Data))
	
	// C関数を呼び出し
	C.radix_sort_uint64(ptr, C.size_t(len(bits)))

	// ビット表現を元のfloat64に戻す
	for i := range bits {
		if bits[i]>>63 != 0 {
			bits[i] = ^bits[i]
		}
		data[i] = math.Float64frombits(bits[i])
	}
}

// RadixSortFloat はfloat64型スライスをRadix Sortでソートする
func RadixSortFloat(data []float64) {
	if len(data) < 2 {
		return
	}

	// float64をuint64に変換
	for i := range data {
		bits := math.Float64bits(data[i])
		// 符号ビットを反転
		if bits>>63 == 1 {
			bits ^= 0xFFFFFFFFFFFFFFFF
		}
		*(*uint64)(unsafe.Pointer(&data[i])) = bits
	}

	// Radix Sort
	C.radix_sort_float64((*C.uint64_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)))

	// uint64をfloat64に戻す
	for i := range data {
		bits := *(*uint64)(unsafe.Pointer(&data[i]))
		// 符号ビットを反転
		if bits>>63 == 0 {
			bits ^= 0xFFFFFFFFFFFFFFFF
		}
		data[i] = math.Float64frombits(bits)
	}
}
