#include "radix_sort.h"
#include <string.h>

#define BUCKET_SIZE 256
#define RADIX_BITS 8

void radix_sort_uint64(uint64_t* arr, size_t len) {
    if (len < 2) return;

    uint64_t* temp = (uint64_t*)malloc(len * sizeof(uint64_t));
    uint32_t* counts = (uint32_t*)calloc(BUCKET_SIZE, sizeof(uint32_t));
    uint32_t* pos = (uint32_t*)malloc(BUCKET_SIZE * sizeof(uint32_t));

    // 8バイト（64ビット）を8ビットずつ8回処理
    for (int shift = 0; shift < 64; shift += RADIX_BITS) {
        // カウントをリセット
        memset(counts, 0, BUCKET_SIZE * sizeof(uint32_t));

        // 各桁の出現回数をカウント
        for (size_t i = 0; i < len; i++) {
            uint32_t digit = (arr[i] >> shift) & (BUCKET_SIZE - 1);
            counts[digit]++;
        }

        // 累積和を計算して位置を決定
        pos[0] = 0;
        for (int i = 1; i < BUCKET_SIZE; i++) {
            pos[i] = pos[i-1] + counts[i-1];
        }

        // 要素を一時配列に移動
        for (size_t i = 0; i < len; i++) {
            uint32_t digit = (arr[i] >> shift) & (BUCKET_SIZE - 1);
            temp[pos[digit]++] = arr[i];
        }

        // 結果を元の配列にコピー
        memcpy(arr, temp, len * sizeof(uint64_t));
    }

    free(temp);
    free(counts);
    free(pos);
}
