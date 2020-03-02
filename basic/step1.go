package main

import "fmt"

/* объявление переменной на уровне пакета */
var c, python, java = "test", true, false

func main() {
	fmt.Println("add result: ", add(42, 113))
	a, b := swap("привет", "олень")
	fmt.Println(a, b)
	fmt.Println(split(17))

	/* объявление переменной на уровне функции */
	var i int = 1
	fmt.Println(i, c, python, java)

	/* неявное указание типа переменной */
	k := 3
	fmt.Println(k)

	f := float64(i)
	u := uint(f)
	fmt.Println(i, f, u)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

/* функция без комментирования возврата */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
