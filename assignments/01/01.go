package first_assinment

import (
	"fmt"
	"math"
	"math/rand"
)

func add(a int, b int) int {
	return a + b
}

const Test = 50

var c, python, java bool

func Run_01() {
	var a, b, d = true, false, "Test"
	fmt.Println("Test", rand.Intn(5), math.Pi)
	println(add(5, 15))
	println(c, python, java)
	println(a, b, d)
	println(complex(10, 1))
}
