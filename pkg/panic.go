package pkg

import (
	"fmt"
)

func WorkingWithPanic() {
	fmt.Println("Working with panic")
	panic("a problem") // This will cause the program to panic and not recover
}

func WorkingWithPanicAndRecover() {
	arr := []int{1, 2, 3}
	val, err := safeValue(arr, 10)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Value:", val)
}

func safeValue(val[] int, idx int) (n int, err error) {
	defer func() {
		if e:= recover(); e != nil {
			err = fmt.Errorf("%v", e)
			n = 0
			// return 0, err // This will not work as return is not allowed in defer
		}
	}()
	return val[idx], nil // if this line causes panic, then defer will catch it
}