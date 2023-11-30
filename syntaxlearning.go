// package main
// import "fmt"
// func main(){
// 	var i, j = 20,10
//     fmt.Printf("i has the value of %v and type of %T\n", i,i)
// 	fmt.Printf("j has the value of %v and type of %T\n", j, j)
// }

// package main
// import "fmt"
// func main(){
//     var i = 15.5
// 	var txt = "Hello World"
// 	fmt.Printf("%v\n", i)
// 	fmt.Printf("%#v\n", i)
// 	fmt.Printf("%v%%\n", i)
// 	fmt.Printf("%T\n", i)

// 	fmt. Printf("%v\n", txt)
// 	fmt.Printf("%#v\n", txt)
// 	fmt.Printf("%T\n", txt)
// }

// package main

// import "fmt"

// func main() {
// 	var i = 15
// 	fmt.Printf("%b\n", i)
// 	fmt.Printf("%d\n", i)
// 	fmt.Printf("%+d\n", i)
// 	fmt.Printf("%o\n", i)
// 	fmt.Printf("%O\n", i)
// 	fmt.Printf("%x\n", i)
// 	fmt.Printf("%X\n", i)
// 	fmt.Printf("%#x\n", i)
// 	fmt.Printf("%4d\n", i)
// 	fmt.Printf("%-4d\n", i)
// 	fmt.Printf("%04d\n", i)
// }

//We go again
// package main
// import "fmt"
// func main(){
// 	var txt = "Hello"
// 	fmt. Printf("%s\n", txt)
// 	fmt.Printf("%q\n", txt)
// 	fmt.Printf("%8s\n", txt)
// 	fmt.Printf("%-8s\n", txt)
// 	fmt.Printf("%x\n", txt)
// 	fmt.Printf("% x\n", txt)
//}
//We go again
// package main
// import "fmt"
// func main(){
// 	var a bool  = true
// 	var b int = 5
// 	var c float32 = 3.14
// 	var d string = "Hi!"

//     fmt.Println("Boolean: ", a)
// 	fmt.Println("Integer: ", b)
// 	fmt.Println("Float: ", c)
// 	fmt.Println("String: ", d)
// }

//
// package main
// import "fmt"
// func main(){
// 	var a bool = true
// 	var b = true
// 	var c bool
// 	d:=true
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

package main

import "fmt"

func main() {
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8, 9, 10, 11}
	arr2[3]= 200
	var Name = [...]string{"Volvo,", "Mercendez,", "Toyota,", "SUV"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(Name)
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
