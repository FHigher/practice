package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"wang", "li", "zhang"}

	fmt.Println(sort.StringsAreSorted(names))

	sortedNames := sort.StringSlice(names)
	sort.Sort(sortedNames)
	fmt.Println(sortedNames)
}
