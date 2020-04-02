package utils

import (
	"log"
	"strconv"
)

func StrToInt(val string) int {
	v1, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return int(v1)
}
