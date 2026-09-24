package main

import "fmt"

var Package_lev_pub_var string = "Hi Hello"

// we can use the out side of the package.
var package_lev_pri_var int = 180

// we can't use the from the out side package

func main() {
	// local variable
	var localtype string = "hai"

	//with out data type
	var notype = 55

	//shot hand declaration
	sh := "chandu"

	fmt.Println(Package_lev_pub_var, package_lev_pri_var, localtype, sh, notype)

	//integer type
	var Int_val int
	var float_val float32
	var complex_val complex64
	var string_val string
	var unsigned uint
	var rune_type rune
	var byte_type byte
	var uintptr_type uintptr
	var bool_value bool
	fmt.Println(Int_val, float_val, complex_val, string_val, unsigned, rune_type, byte_type, uintptr_type, bool_value)
	//0 0 (0+0i)  0 0 0 0

	//casting
	var old int8 = 25
	new := int(old)
	fmt.Println(new)

}
