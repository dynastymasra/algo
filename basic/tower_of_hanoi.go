package main

import "fmt"

func main() {
	TowerOfHanoi(5, "A", "B", "C")
}

func TowerOfHanoi(n int, source, aux, dest string) {
	if n < 1 {
		return
	}

	TowerOfHanoi(n-1, source, dest, aux)
	fmt.Println("Move disk", n, "from", source, "to", dest)
	TowerOfHanoi(n-1, aux, source, dest)
}
