package main

import (
	"fmt"

	"github.com/chandubatta/my_app/painter"
)
func main(){
	var s1 string ="Chandu"
	fmt.Println("Before Painting")
	fmt.Println(s1)
	fmt.Println("After Painting")
	painter.Public_function(s1)
	
	
}