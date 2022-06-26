package main

import "fmt"

func main() {
	var results []int
	for i := 2; i <= 10000; i += 2 {
		var n = i*3 + 1
		for n != 1 {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = n*3 + 1
			}
			if n == i {
				results = append(results, i)
				break
			}
		}
	}
	fmt.Printf("%d個\n", len(results))
	fmt.Println(results)
}
