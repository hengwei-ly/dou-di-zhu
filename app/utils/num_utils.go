package utils

import "sort"

//在一组正整数里面看是不是连续  如果是  返回最大值  否则返回0
func IsSerial(arr []int) int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)

	sort.Ints(newArr)
	if newArr[len(newArr)-1]-newArr[0] != len(newArr)-1 {
		return 0
	}
	for k := 1; k < len(newArr); k++ {
		if newArr[k]-1 != newArr[k-1] {
			return 0
		}
	}
	return newArr[len(newArr)-1]
}
