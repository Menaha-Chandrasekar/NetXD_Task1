package main

import "fmt"

func counting(main, sub string) int {
	count := 0
	Len_S:= len(sub)
	Len_M:=len(main)

	for i := 0; i <= Len_M-Len_S; i++ {
		if main[i:i+Len_S] == sub {
			count++
		}
	}
	return count
}

func main() {
	string := "go language on go process"
	sub := "go"
	count := counting(string, sub)

	fmt.Printf("%d\n", count)
}
