package utils

import {
	"fmt"
	"syscall"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
