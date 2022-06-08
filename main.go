package main

import (
	"fmt"
)

func main() {
	var name string = "material guuuurl"
	age := 23
	fmt.Printf("\n name is %s and age is %d", name, age)

	pastries := []string{"chocolate", "vanilla", "caramel", "strawberry", "pudding"}
	baked := pastries[0:2:4]
	baked = append(baked, "cookies")
	fmt.Printf("\n pastries is %v and baked is %v", pastries, baked)
}
