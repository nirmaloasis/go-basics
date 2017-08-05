package main

import (
	"fmt"
)

func main() {
	var i int = 5
	fmt.Println(double(i))

}

// func double(x int) int {
// 	return x + x
// }

func double(x int) (result int) {
	defer func() {
        fmt.Printf("double (%d) = %d \n",x,result)
    }()
    return x+x
}
