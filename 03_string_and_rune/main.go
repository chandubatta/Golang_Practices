package main

import (
	"fmt"
	"os"
	"strings"
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
	fmt.Println(string(bytes_data))

	//Rune creation
	var r1 rune = '@'
	var r2 rune
	fmt.Println(r1, r2)
	///Common Operations on String
	//Length
	s4 := "Hello"
	l1 := len(s4)
	fmt.Println(l1) //6
	s5 := "Héllo"
	s6 := "é"
	s7 := "e"

	l2 := len(s5)
	l3 := len(s6)
	l4 := len(s7)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)

	r8 := []rune(s5)
	l5 := len(r8)
	fmt.Println(l5)

	//Compare two strings
	s8 := "Hi"
	s9 := "Hello"
	t1 := s8 == s9
	t2 := s8 >= s9
	t3 := s8 <= s9
	t4 := s8 > s9
	t5 := s8 < s9
	fmt.Println(t1, t2, t3, t4, t5)

	//Using Compare function
	val := strings.Compare(s8, s9)
	fmt.Println(val)

	//lexicographically, meaning Go compares the characters/bytes from left to right.
	//The fact that "Hi" has 2 letters and "Hello" has 5 letters does not matter here.
	//It's better not to think simply:
	//"A" < "B" because A is alphabetically before B.

	val2 := strings.EqualFold(s8, s9)
	fmt.Println(val2)

	//examples
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println(strings.EqualFold("Golang", "golang"))
	fmt.Println(strings.EqualFold("Hi", "Hello"))
	/*
		--output--
		true
		true
		true
		false
	*/

	////Concatination
	//Using the + operator
	s10 := s8 + " " + s9
	fmt.Println(s10)
	s11 := s10 + " " + s8 + " " + s9
	fmt.Println(s11)
	// Using Builder
	B := strings.Builder{}
	B.Grow(1024) //Grow() function is used for fixed size fo the building string
	B.WriteString(s4)
	B.WriteString(" ")
	B.WriteString(s5)
	fmt.Println(B.String()) //result:=B.String()

	//Access Index
	s12 := "Hi Chandu, How are you"
	i1 := string(s12[5])
	fmt.Println(i1)
	ss := "Hello"
	vk1 := ss[1]
	fmt.Println(string(vk1))
	sk := "Héllo"
	vk := sk[1] //byte value coming
	fmt.Println(vk)

	//Substring
	k1 := "chandu is good boy"
	k2 := k1[3:9]
	fmt.Println(k2) //direct string value comimg
	k3 := k1[3:]
	fmt.Println(k3)

}
