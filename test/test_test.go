package test

import (
	"container/list"
	"fmt"
	"note-gin/utils/RedisClient"
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

func TestRedis(t *testing.T) {
	RedisClient.RedisInit()
	length := RedisClient.RedisClient.LLen("folder_nav").Val()
	fmt.Println(length)
}
