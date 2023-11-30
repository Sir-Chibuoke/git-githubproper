//We are here again to talk and learn. So help me God
package main
import "fmt"
func main(){
	arr1:=[6]int{10, 11, 12, 13, 14, 15}
	myslice:=arr1[2:4]
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("Lenght = %d\n", len(myslice))
	fmt.Printf("Capacity = %d\n", cap(myslice))
}