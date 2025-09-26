package main
import "fmt"

//STRUCTS to rep str data
type car struct {
	Make string
	Model string
	Height int
	Width int	
	FrontWheel Wheel //nested struct
	BackWheel Wheel
}
type Wheel struct{
	Radius int
	Material string
}
//use dot operator . to access fields of a struct
func main(){
	myCar:=car{}
	myCar.FrontWheel.Radius = 3 //radius is accessed from wheel 
	fmt.Println(myCar)
}

//another eg of struct
// myCar := struct {
// 	Make string
// 	Model string
// } {
// 	Make : "tesla",
// 	Model : "model 3"
// }

// //embed (or inherit) a struct in another
// type car struct {
// 	make string
// 	model string
// }
// type truck struct {
// 	car //embed car
// 	bedSize int
// }

// //struct methods in Go
// type rect struct {
// 	width int
// 	height int
// }
// //area has a receiver of (r rect)
// func (r rect) area () int {
// 	return r.width * r.height
// }
// r := rect {
// 	width: 5,
// 	height: 10,
// }
// fmt.package(r.area())
// //prints 50
