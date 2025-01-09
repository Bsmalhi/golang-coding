package main

import "fmt"

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
}