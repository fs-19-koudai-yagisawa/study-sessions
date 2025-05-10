#ifndef RADIX_SORT_H
#define RADIX_SORT_H

#include <stdint.h>
#include <stdlib.h>

// uint32用Radix Sort
void radix_sort_int(uint32_t* arr, size_t len);

// uint64用Radix Sort（float64変換用）
void radix_sort_uint64(uint64_t* arr, size_t len);

#endif // RADIX_SORT_H
