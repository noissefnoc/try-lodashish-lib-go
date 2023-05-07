package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
	fmt.Println(names)
}
