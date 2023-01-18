package main

import "fmt"

func main() {
	colors := map[string]string{ //key,val type
		"red":   "#ff000",
		"green": "#4bf745",
	}
	fmt.Println(colors)

	colors2 := make(map[int]string) //also crt map.

	colors2[10] = "xyz" //add key ,val
	fmt.Println(colors2)

	delete(colors2, 10)
	fmt.Println(colors2)

	print_col(colors)
}

func print_col(col map[string]string) {
	for key, val := range col {
		fmt.Println(key, " : ", val)
	}
}
