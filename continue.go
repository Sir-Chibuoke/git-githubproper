// This is for continue program
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("\n%d", i)
	}
}
