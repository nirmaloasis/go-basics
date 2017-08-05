package main

import (
	"fmt"
    "unsafe"
)

func main() {
	 i := 7
     s := "bretton woods"
     a := [] int {1,3,4}
	fmt.Println(unsafe.Sizeof(i))
    fmt.Println(unsafe.Sizeof(s))
    fmt.Println(unsafe.Sizeof(a))
    fmt.Println(a)

}

// func double(x int) int {
// 	return x + x
// }

// func double(x int) (result int) {
// 	defer func() {
//         fmt.Println("double (%d) = %d \n",x,result)
//     }()
//     return x+x
// }
