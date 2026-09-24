package main

import (
	"fmt"
	"os"
)

func main() {
	//string creation
	var s1 string
	var s2 string = "Hello"

	s3 := "How are you"
	fmt.Println(s1, s2, s3)

	// Read file using bytes
	bytes_data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("data type %T and data %v", bytes_data, bytes_data)
	fmt.Println()
	fmt.Print(string(bytes_data))

}
