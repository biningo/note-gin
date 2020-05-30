package test

import (
	"log"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T){
	s:=time.Now().Format(time.RFC3339)
	log.Println(s)
	log.Println(time.Parse("2006-01-02 15:04:05",s))
}
