package main

import "fmt"

func removeElement(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", s)

	i := 2 // индекс элемента, который нужно удалить
	s = removeElement(s, i)
	fmt.Println("Slice after removal:", s)
}
