
package main

import "fmt"

func oddOrEven(n *int) {

	if *n%2 == 0 {

		*n = 0

	} else {

		*n = 1

	}

}

func main() {

	x := 2
	y := 3

	oddOrEven(&x)

	fmt.Println(x)

	oddOrEven(&y)

	fmt.Println(y)

}
