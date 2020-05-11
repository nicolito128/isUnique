package main

import (
	"fmt"
	unique "isUnique/src/isUnique"
)

func main() {
	uniqueNums := []int{1, 2, 3, 4}
	nums := []int{2, 4, 5, 6, 6}

	fmt.Println(unique.Int(uniqueNums))
	fmt.Println(unique.Int(nums))

	uniqueNames := []string{"Tom", "Wallace", "Carl"}
	names := []string{"Jerry", "Lua", "Vicky", "Vicky"}

	fmt.Println(unique.Str(uniqueNames))
	fmt.Println(unique.Str(names))
}
