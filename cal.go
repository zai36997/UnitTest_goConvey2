package main

import (
	"errors"
	"fmt"
)

func Add(x ...int) int {
	z := 0
	for _, y := range x {
		z += y
	}
	return z
}

func Minus(t int, x ...int) int {
	r := t
	for _, y := range x {
		r -= y
	}
	return r
}

func Multiplies(x ...int) int {
	q := 1
	for _, y := range x {
		q *= y
	}
	return q
}

func DividedBy(e float64, x ...float64) (float64, error) {
	w := e
	for _, y := range x {
		if y == 0 {
			return y, errors.New("y != 0")
		}
		w /= y
	}
	return w, nil
}

func main() {
	add := Add(2, 2, 2)
	fmt.Println("2 + 2 + 2 = ", add)
	minus := Minus(2, 2, 2)
	fmt.Println("2 - 2 - 2 = ", minus)
	multiplies := Multiplies(2, 2, 2)
	fmt.Println("2 * 2 * 2 = ", multiplies)
	dividedBy, _ := DividedBy(10.0, 2.0, 2.0)
	fmt.Println("2 / 2 / 2 = ", dividedBy)

}
