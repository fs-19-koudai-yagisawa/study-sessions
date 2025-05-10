package fastpath

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

void radix_sort_int32(int32_t* arr, size_t len) {
    if (len < 2) return;

    // 最大値と最小値を見つける
    int32_t max = arr[0], min = arr[0];
    for (size_t i = 1; i < len; i++) {
        if (arr[i] > max) max = arr[i];
        if (arr[i] < min) min = arr[i];
    }

    // オフセットを計算（負の数を扱うため）
    int32_t offset = 0;
    if (min < 0) {
        offset = -min;
        max += offset;
        for (size_t i = 0; i < len; i++) {
            arr[i] += offset;
        }
    }

    // 各桁でソート
    int32_t* output = (int32_t*)malloc(len * sizeof(int32_t));
    if (output == NULL) return;
    
    int32_t count[256] = {0};

    // 8ビットずつ処理
    for (int shift = 0; shift < 32; shift += 8) {
        memset(count, 0, sizeof(count));

        // カウント
        for (size_t i = 0; i < len; i++) {
            count[(arr[i] >> shift) & 0xFF]++;
        }

        // 累積和
        for (int i = 1; i < 256; i++) {
            count[i] += count[i-1];
        }

        // 出力配列の構築
        for (int i = len-1; i >= 0; i--) {
            output[--count[(arr[i] >> shift) & 0xFF]] = arr[i];
        }

        // 元の配列にコピー
        memcpy(arr, output, len * sizeof(int32_t));
    }

    // オフセットを元に戻す
    if (offset > 0) {
        for (size_t i = 0; i < len; i++) {
            arr[i] -= offset;
        }
    }

    free(output);
}

void radix_sort_int64(int64_t* arr, size_t len) {
    if (len < 2) return;

    // 最大値と最小値を見つける
    int64_t max = arr[0], min = arr[0];
    for (size_t i = 1; i < len; i++) {
        if (arr[i] > max) max = arr[i];
        if (arr[i] < min) min = arr[i];
    }

    // オフセットを計算（負の数を扱うため）
    int64_t offset = 0;
    if (min < 0) {
        offset = -min;
        max += offset;
        for (size_t i = 0; i < len; i++) {
            arr[i] += offset;
        }
    }

    // 各桁でソート
    int64_t* output = (int64_t*)malloc(len * sizeof(int64_t));
    if (output == NULL) return;
    
    int64_t count[256] = {0};

    // 8ビットずつ処理
    for (int shift = 0; shift < 64; shift += 8) {
        memset(count, 0, sizeof(count));

        // カウント
        for (size_t i = 0; i < len; i++) {
            count[(arr[i] >> shift) & 0xFF]++;
        }

        // 累積和
        for (int i = 1; i < 256; i++) {
            count[i] += count[i-1];
        }

        // 出力配列の構築
        for (int i = len-1; i >= 0; i--) {
            output[--count[(arr[i] >> shift) & 0xFF]] = arr[i];
        }

        // 元の配列にコピー
        memcpy(arr, output, len * sizeof(int64_t));
    }

    // オフセットを元に戻す
    if (offset > 0) {
        for (size_t i = 0; i < len; i++) {
            arr[i] -= offset;
        }
    }

    free(output);
}

#cgo CFLAGS: -O3
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// RadixSortInt は32ビット整数スライスをRadix Sortでソートする
func RadixSortInt(data []int32) {
	if len(data) < 2 {
		return
	}

	// スライスヘッダから直接メモリを参照
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	ptr := (*C.int32_t)(unsafe.Pointer(header.Data))
	
	// C関数を呼び出し
	C.radix_sort_int32(ptr, C.size_t(len(data)))
}

// RadixSortInt64 は64ビット整数スライスをRadix Sortでソートする
func RadixSortInt64(data []int64) {
	if len(data) < 2 {
		return
	}

	// スライスヘッダから直接メモリを参照
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	ptr := (*C.int64_t)(unsafe.Pointer(header.Data))
	
	// C関数を呼び出し
	C.radix_sort_int64(ptr, C.size_t(len(data)))
}
