// //We go again
// package main
// import "fmt"
// func main(){
// 	prices:=[]int{10, 20, 30}
// 	prices[2]=5000
// 	fmt.Println(prices[0])
// 	fmt.Println(prices[2])
// }

// //We everly go
// package main
// import "fmt"
// func main(){
// 	myslice :=[]int{1, 2, 3, 4, 5, 6}
// 	fmt.Printf("myslice = %v\n", myslice)
// 	fmt.Printf("Lenght = %d\n", len(myslice))
// 	fmt.Printf("Capacity = %d\n", cap(myslice))

// 	myslice = append(myslice, 20, 21, 22)
// 	fmt.Printf("myslice = %v\n", myslice)
// 	fmt.Printf("Lenght = %d\n", len(myslice))
// 	fmt.Printf("Capacity = %d\n", cap(myslice))

// }
// //Here we go again
// package main
// import "fmt"
// func main(){
// 	myslice:=[]int{1,2,3}
// 	myslice1:=[]int{3, 4, 5, 6}
// 	myslice3:=append(myslice, myslice1...)

// 	fmt.Printf("myslice3 = %v\n", myslice3)
// 	fmt.Printf("Lenght = %d\n", len(myslice3))
// 	fmt.Printf("Capacity = %d\n", cap(myslice3))
// 	myslic=make()

// package main
// import "fmt"
// func main(){
// 	var a=20
// 	var b=18
// 	if a<b {
// 		fmt.Printf("\n20 is greather than 18")
// 	} 
// 	else {
// 		fmt. Printf("\nIt can't work my brother, please kindly try again")
// 	}
// }
//We go again
// package main
// import "fmt"
// func main(){
// 	time:=20
// 	if (time<18){
// 		fmt.Printf("\nGood day Boss")
// 	}else{
// 		fmt.Printf("\nGood Night, wishing you the very best")
// 	}
// }
package main
import "fmt"
func main (){
temperature:=14
if (temperature>15){
	fmt.Printf("\nIt is warm out there")
}else{
	fmt.Printf("It is cold out there")
}
}