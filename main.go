package main

import (
	"fmt"
	"strings"

	"github.com/bsmalhi/golang-coding/pkg"
	concurrency "github.com/bsmalhi/golang-coding/pkg/concurrency"
)


func main() {
		fmt.Println("Hello, World!")
		fmt.Println("Fun with GoLang")

		fmt.Println("Looping through array")
        arr := [5]string{"a", "b", "c"}
        for index, element := range arr {
                fmt.Println(index, "->", element)
        }

		var arr1 [10]int = [10]int{}
		for index, _ := range arr1 {
			if arr1[index] == 0 {
				arr1[index] = 1
			}
			fmt.Printf("Arr1 %v -> %v \n", index, arr1[index])
		}

		arr2 := [...]int{10, 100, 100} // array with size as per elements
		fmt.Println(len(arr2))
		fmt.Println(arr2)
		arr2[2] = 12
		fmt.Println(arr2)

		for i := 0; i < len(arr2); i++ {
			fmt.Printf("Arr2 %v -> %v \n", i, arr2[i])
		}

		var a = 0 // untyped variable declartion
		fmt.Println("Value is ", a)

		fmt.Println("Printing switch without condition")
		a = 10
		switch a {
			case 10:
				fmt.Println("Printing 10 now")
			default:
				fmt.Println("Printing 11 now")
		}
		fmt.Println("Printing switch with condition")
		switch {
			case a == 10:
				fmt.Println("Printing 10 now")
				fallthrough
			case a == 20:
				fmt.Println("Also Printing 20 now due to fallthrough")
			default:
				fmt.Println("Printing 11 now")
		}

		fmt.Println("Printing slices")
		arr4 := [5]int{10, 20, 90, 70, 60}
        slice := arr4[:3]
        fmt.Println(cap(slice))

        slice_2 := make([]int, 5, 20)
        new_slice := append(slice, slice_2...)
        fmt.Println(cap(new_slice))

		slice_3 := make([]int, 5, 10)
		fmt.Println(cap(slice_3))
		fmt.Println(len(slice_3))

		slice_4 := []int{10, 20, 30, 40, 50}
		fmt.Printf("type of slice_4 %T\n", slice_4)
		fmt.Println(cap(slice_4[:3]))	
		slice_5 := slice_4[:3]
		fmt.Println(cap(slice_5))	
		fmt.Println(len(slice_5))	

		fmt.Println("Printing maps")

		fmt.Println("Working with functions")

		fmt.Println("Working with anonymous functions")
		x := func(s string) string {
			return strings.ToUpper(s)
		}	
		fmt.Printf("%T \n", x)
		fmt.Println(x("Joe"))

		fmt.Println("Working with High Order Functions")

		fmt.Println("Working with pointers")
		var y int
        var ptr *int = &y

        *ptr = 0
        fmt.Println(y)

        *ptr += 5
        fmt.Println(y)

		fmt.Println("Working with Pass by Value")

		fmt.Println("Working with Pass by Reference")

		fmt.Println("Working with structs")

        square := &pkg.Square{X: 5, Y: 10, Length: 11}
        square.Move(4, 5)
        fmt.Printf("%+v\n", square)

		fmt.Println("Working with short hand declaration of structs")
		
		fmt.Println("Working with interfaces")

        car := pkg.BuildVehicle("car", 2021, 10000)
        pkg.DisplayVehicleInfo(car)

        motorbike := pkg.BuildVehicle("motorbike", 2021, 10000)
        pkg.DisplayVehicleInfo(motorbike)

		fmt.Println("Working with methods")

        fmt.Println("Working with errors")
        pkg.WorkingWithErrors()

        fmt.Println("Working with go panic")
        // pkg.WorkingWithPanic() // this stops the program from running due to panic
        
        fmt.Println("Working with go panic and recover")
        pkg.WorkingWithPanicAndRecover()

        fmt.Println("Working with go routines")
        concurrency.WorkingWithSerialExecutionUrlFetch()
        concurrency.WorkingWithGoRoutinesUrlFetch()

        fmt.Println("Working with channels")
        // concurrency.WorkingWithChannelsDeadlockExample() // This cause deadlock due to main channel waiting for value to be read
        concurrency.WorkingWithChannelGoRoutine()
        concurrency.WorkingWithChannelsGoRoutinesUrlFetch()
        
}