package main

import (
	"fmt"
	"sort"
)

func Permutations(input string) []string {
	result := []string{}
	permute("", input, &result)
	sort.Strings(result)
	return result
}

func permute(prefix, suffix string, result *[]string) {
	if len(suffix) == 0 {
		*result = append(*result, prefix)
		return
	}
	for i := 0; i < len(suffix); i++ {
		permute(prefix+string(suffix[i]), suffix[:i]+suffix[i+1:], result)
		}
}

func main() {
	var str string
	fmt.Scanln(&str)
	fmt.Println(Permutations(str))
}