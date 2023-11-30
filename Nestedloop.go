// //Nested Loops
// package main
// import "fmt"
// func main(){
// 	adj:=[2]string{"big", "tasty"}
// 	fruits:=[3]string{"apple","orange", "banana"}
// 	for i:=0; i<len(adj); i++ {
//         for j:=0; j<len(fruits); j++ {
//             fmt.Printf("\n%s %s", adj[i], fruits[j])
//         }
//     }
// }
//Here we go again also
package main
import "fmt"
func main (){
fruits:=[3]string{"apple", "orange", "banana"}
for idx, val := range fruits{
	fmt. Printf("%v\t%v\n", idx, val)
}
}