package main

import (
	"time"
)

func Wait(num int) {
	time.Sleep(time.Duration(num) * time.Second)
}