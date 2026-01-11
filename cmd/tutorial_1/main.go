package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	ownerInfo owner
}
type owner struct{
	name string
}
func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr2 := [...]int32{1, 2, 3}
	fmt.Println(intArr)
	fmt.Println(intArr2)

}
func printMe() {
	fmt.Println("Hello World")
}
