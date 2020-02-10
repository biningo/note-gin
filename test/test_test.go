package test

import (
	"container/list"
	"fmt"
	"testing"
	"time"
)

func change(arr *[]int) {
	*arr = append(*arr, 8)

}

type Test struct {
	Flag bool
	Time time.Time
}

func TestSub(t *testing.T) {
	arr := []int{}
	change(&arr)
	fmt.Println(arr)

	arr = append(arr, 1)
	fmt.Println(arr)

	q := list.New()
	q.PushBack(1)
	fmt.Println()

	arr = []int{1, 2, 3}
	arr = append(arr, 8, 9)
	tt := Test{}
	fmt.Println(tt.Time.IsZero())
}
