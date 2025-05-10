#include "radix_sort.h"
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

#define BUCKET_SIZE 10

void radix_sort_int(int32_t* arr, size_t len) {
    if (len < 2) return;

    int32_t* temp = (int32_t*)malloc(len * sizeof(int32_t));
    if (temp == NULL) {
        printf("Error: Failed to allocate memory for temp array\n");
        return;
    }

    // 最大値と最小値を見つける
    int32_t max = arr[0];
    int32_t min = arr[0];
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
    for (uint32_t exp = 1; max/exp > 0; exp *= 10) {
        int count[10] = {0};  // 各桁の出現回数をカウント

        // 各桁の出現回数をカウント
        for (size_t i = 0; i < len; i++) {
            int digit = (arr[i]/exp) % 10;
            count[digit]++;
        }

        // 累積和を計算（開始位置を求める）
        for (int i = 1; i < 10; i++) {
            count[i] += count[i-1];
        }

        // 後ろから走査して安定ソートを実現
        for (int i = len-1; i >= 0; i--) {
            int digit = (arr[i]/exp) % 10;
            temp[count[digit]-1] = arr[i];
            count[digit]--;
        }

        // 結果を元の配列にコピー
        memcpy(arr, temp, len * sizeof(int32_t));
    }

    // オフセットを元に戻す
    if (offset > 0) {
        for (size_t i = 0; i < len; i++) {
            arr[i] -= offset;
        }
    }

    free(temp);
}
