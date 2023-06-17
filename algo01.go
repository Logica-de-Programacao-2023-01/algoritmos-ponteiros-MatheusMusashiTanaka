package main

import "fmt"

func somaNaturais(p *int, n int) {

	soma := 0

	for i := 1; i <= n; i++ {

		soma += i

	}

	*p = soma

}

func main() {

	x := 0

	somaNaturais(&x, 5)

	fmt.Println(x)

}
