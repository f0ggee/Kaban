package Handlers

import (
	"fmt"
	"time"
)

func Check() {

	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Second)
		fmt.Println(i)

	}
}
