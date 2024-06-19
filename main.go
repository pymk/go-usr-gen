package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	num := 1
	var err error

	if len(os.Args) == 2 {
		num, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("The argument must be an integer.")
			return
		}
	}

	for i := 0; i < num; i++ {
		fmt.Println(generator())
	}
}

func letters() []string {
	alphabet := []string{}

	for i := 0; i < 26; i++ {
		alphabet = append(alphabet, string(rune('A'+i)))
	}

	return alphabet

}

func generator() string {
	a := letters()

	c := make([]string, 5)
	d := make([]string, 5)

	for j := 0; j < 5; j++ {
		c[j] = strings.ToLower(a[rand.IntN(len(a))])
		d[j] = strconv.Itoa(rand.IntN(9))
	}

	r := []string{strings.Join(c, ""), strings.Join(d, "")}
	return strings.Join(r, "_")

}
