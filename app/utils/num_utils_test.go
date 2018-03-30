package utils

import (
	"fmt"
	"testing"
)

func TestIsSerial(t *testing.T) {
	arr := []int{1, 2, 6, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(IsSerial(arr))
	fmt.Println(arr)
}
