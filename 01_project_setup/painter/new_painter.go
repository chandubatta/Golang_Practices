package painter

import (
	//"fmt"

	"github.com/fatih/color"
)

func Public_function(s string) {
	private_function(s)
}
func private_function(s string) {
	//fmt.Println("HI I AM THE PAINTER")
	color.Red(s)
	color.White(s)
	color.Green(s)
}
